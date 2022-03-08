package bff

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/trisacrypto/directory/pkg"
	"github.com/trisacrypto/directory/pkg/bff/api/v1"
	"github.com/trisacrypto/directory/pkg/bff/config"
	"github.com/trisacrypto/directory/pkg/utils/logger"
	gds "github.com/trisacrypto/trisa/pkg/trisa/gds/api/v1beta1"
)

func init() {
	// Initialize zerolog with GCP logging requirements
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.TimestampFieldName = logger.GCPFieldKeyTime
	zerolog.MessageFieldName = logger.GCPFieldKeyMsg

	// Add the severity hook for GCP logging
	var gcpHook logger.SeverityHook
	log.Logger = zerolog.New(os.Stdout).Hook(gcpHook).With().Timestamp().Logger()
}

// New creates a new BFF server from the specified configuration.
func New(conf config.Config) (s *Server, err error) {
	// Load the default configuration from the environment if config is empty
	if conf.IsZero() {
		if conf, err = config.New(); err != nil {
			return nil, err
		}
	}

	// Set the global level
	zerolog.SetGlobalLevel(conf.GetLogLevel())

	// Set human readable logging if specified
	if conf.ConsoleLog {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	// Create the server and prepare to serve
	s = &Server{
		conf:  conf,
		echan: make(chan error, 1),
	}

	// Connect to the TestNet and MainNet directory services
	if s.testnet, err = ConnectGDS(conf.TestNet); err != nil {
		return nil, fmt.Errorf("could not connect to the TestNet: %s", err)
	}

	if s.mainnet, err = ConnectGDS(conf.MainNet); err != nil {
		return nil, fmt.Errorf("could not connect to the MainNet: %s", err)
	}

	// Create the router
	gin.SetMode(conf.Mode)
	s.router = gin.New()
	if err = s.setupRoutes(); err != nil {
		return nil, err
	}

	// Create the http server
	s.srv = &http.Server{
		Addr:         s.conf.BindAddr,
		Handler:      s.router,
		ErrorLog:     nil,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	return s, nil
}

type Server struct {
	sync.RWMutex
	conf    config.Config
	srv     *http.Server
	router  *gin.Engine
	testnet gds.TRISADirectoryClient
	mainnet gds.TRISADirectoryClient
	healthy bool
	echan   chan error
}

// Serve API requests on the specified address.
func (s *Server) Serve() (err error) {
	// Catch OS signals for graceful shutdowns
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit
		s.echan <- s.Shutdown()
	}()

	// Set the health of the service to true unless we're in maintenance mode.
	// The server should still start so that it can return unavailable to requests.
	s.SetHealth(!s.conf.Maintenance)
	if s.conf.Maintenance {
		log.Warn().Msg("starting server in maintenance mode")
	}

	// Listen for HTTP requests on the specified address and port
	go func() {
		log.Info().
			Str("listen", s.conf.BindAddr).
			Str("version", pkg.Version()).
			Msg("gds bff server started")

		if err = s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.echan <- err
		}
	}()

	// Listen for any errors that might have occurred and wait for all go routines to stop
	if err = <-s.echan; err != nil {
		return err
	}
	return nil
}

func (s *Server) Shutdown() (err error) {
	log.Info().Msg("gracefully shutting down")

	s.SetHealth(false)
	s.srv.SetKeepAlivesEnabled(false)

	// Require shutdown in 30 seconds without blocking
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err = s.srv.Shutdown(ctx); err != nil {
		return err
	}

	log.Debug().Msg("successfully shutdown server")
	return nil
}

func (s *Server) SetHealth(health bool) {
	s.Lock()
	s.healthy = health
	s.Unlock()
	log.Debug().Bool("healthy", health).Msg("server health set")
}

func (s *Server) setupRoutes() (err error) {
	// Application Middleware
	s.router.Use(ginzerolog.Logger("gin"))
	s.router.Use(gin.Recovery())
	s.router.Use(s.Available())

	// NotFound and NotAllowed routes
	s.router.NoRoute(api.NotFound)
	s.router.NoMethod(api.NotAllowed)
	return nil
}

//===========================================================================
// Accessors - used primarily for testing
//===========================================================================

// GetConf returns a copy of the current configuration
func (s *Server) GetConf() config.Config {
	return s.conf
}

// GetRouter returns the Gin API router for testing purposes.
func (s *Server) GetRouter() http.Handler {
	return s.router
}

// GetTestNet returns the TestNet directory client for testing purposes.
func (s *Server) GetTestNet() gds.TRISADirectoryClient {
	return s.testnet
}

// GetMainNet returns the MainNet directory client for testing purposes.
func (s *Server) GetMainNet() gds.TRISADirectoryClient {
	return s.mainnet
}
