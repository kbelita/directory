// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.1
// source: gds/models/v1/models.proto

package models

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CertificateRequestState int32

const (
	CertificateRequestState_INITIALIZED     CertificateRequestState = 0
	CertificateRequestState_READY_TO_SUBMIT CertificateRequestState = 1
	CertificateRequestState_PROCESSING      CertificateRequestState = 2
	CertificateRequestState_DOWNLOADING     CertificateRequestState = 3
	CertificateRequestState_DOWNLOADED      CertificateRequestState = 4
	CertificateRequestState_COMPLETED       CertificateRequestState = 5
	CertificateRequestState_CR_REJECTED     CertificateRequestState = 6
	CertificateRequestState_CR_ERRORED      CertificateRequestState = 7
)

// Enum value maps for CertificateRequestState.
var (
	CertificateRequestState_name = map[int32]string{
		0: "INITIALIZED",
		1: "READY_TO_SUBMIT",
		2: "PROCESSING",
		3: "DOWNLOADING",
		4: "DOWNLOADED",
		5: "COMPLETED",
		6: "CR_REJECTED",
		7: "CR_ERRORED",
	}
	CertificateRequestState_value = map[string]int32{
		"INITIALIZED":     0,
		"READY_TO_SUBMIT": 1,
		"PROCESSING":      2,
		"DOWNLOADING":     3,
		"DOWNLOADED":      4,
		"COMPLETED":       5,
		"CR_REJECTED":     6,
		"CR_ERRORED":      7,
	}
)

func (x CertificateRequestState) Enum() *CertificateRequestState {
	p := new(CertificateRequestState)
	*p = x
	return p
}

func (x CertificateRequestState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CertificateRequestState) Descriptor() protoreflect.EnumDescriptor {
	return file_gds_models_v1_models_proto_enumTypes[0].Descriptor()
}

func (CertificateRequestState) Type() protoreflect.EnumType {
	return &file_gds_models_v1_models_proto_enumTypes[0]
}

func (x CertificateRequestState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CertificateRequestState.Descriptor instead.
func (CertificateRequestState) EnumDescriptor() ([]byte, []int) {
	return file_gds_models_v1_models_proto_rawDescGZIP(), []int{0}
}

// Certificate requests are maintained separately from the VASP record since they should
// not be replicated. E.g. every directory process is responsible for certificate
// issuance and only public keys and certificate metadata should be exchanged between
// directories.
type CertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID only needs to be unique with the requesting directory service.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// VASP information for the request
	Vasp       string `protobuf:"bytes,2,opt,name=vasp,proto3" json:"vasp,omitempty"`
	CommonName string `protobuf:"bytes,3,opt,name=common_name,json=commonName,proto3" json:"common_name,omitempty"`
	// Request pipeline status
	Status CertificateRequestState `protobuf:"varint,4,opt,name=status,proto3,enum=gds.models.v1.CertificateRequestState" json:"status,omitempty"`
	// Sectigo create single certificate batch metadata
	AuthorityId  int64  `protobuf:"varint,5,opt,name=authority_id,json=authorityId,proto3" json:"authority_id,omitempty"`
	BatchId      int64  `protobuf:"varint,6,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	BatchName    string `protobuf:"bytes,7,opt,name=batch_name,json=batchName,proto3" json:"batch_name,omitempty"`
	BatchStatus  string `protobuf:"bytes,8,opt,name=batch_status,json=batchStatus,proto3" json:"batch_status,omitempty"`
	OrderNumber  int64  `protobuf:"varint,9,opt,name=order_number,json=orderNumber,proto3" json:"order_number,omitempty"`
	CreationDate string `protobuf:"bytes,10,opt,name=creation_date,json=creationDate,proto3" json:"creation_date,omitempty"`
	Profile      string `protobuf:"bytes,11,opt,name=profile,proto3" json:"profile,omitempty"`
	RejectReason string `protobuf:"bytes,12,opt,name=reject_reason,json=rejectReason,proto3" json:"reject_reason,omitempty"`
	// Currently there is no choice but to store the PKCS12 password on the server
	// because the certificate issuance command is separated in time by human review,
	// which means it cannot be ephemerally kept in memory during the certificate
	// generation request handoff. Moreover, the value cannot be hashed since it must
	// be sent in plaintext via the Sectigo API and used to decrypt the downloaded
	// packet in order to retrieve the public key material. To mitigate issues with the
	// storage of this data, it is encrypted so that only a directory service with the
	// correct secret key (e.g. the directory service process that created the request)
	// can read the original value. Because this value should not be replicated it is
	// stored in the certificate request which is maintained by independent replicas.
	// TODO: Rebecca to remove for Google Secrets
	Pkcs12Password  []byte `protobuf:"bytes,13,opt,name=pkcs12password,proto3" json:"pkcs12password,omitempty"`
	Pkcs12Signature []byte `protobuf:"bytes,14,opt,name=pkcs12signature,proto3" json:"pkcs12signature,omitempty"`
	// Logging information timestamps
	Created  string `protobuf:"bytes,15,opt,name=created,proto3" json:"created,omitempty"`
	Modified string `protobuf:"bytes,16,opt,name=modified,proto3" json:"modified,omitempty"`
}

func (x *CertificateRequest) Reset() {
	*x = CertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gds_models_v1_models_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertificateRequest) ProtoMessage() {}

func (x *CertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gds_models_v1_models_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertificateRequest.ProtoReflect.Descriptor instead.
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return file_gds_models_v1_models_proto_rawDescGZIP(), []int{0}
}

func (x *CertificateRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CertificateRequest) GetVasp() string {
	if x != nil {
		return x.Vasp
	}
	return ""
}

func (x *CertificateRequest) GetCommonName() string {
	if x != nil {
		return x.CommonName
	}
	return ""
}

func (x *CertificateRequest) GetStatus() CertificateRequestState {
	if x != nil {
		return x.Status
	}
	return CertificateRequestState_INITIALIZED
}

func (x *CertificateRequest) GetAuthorityId() int64 {
	if x != nil {
		return x.AuthorityId
	}
	return 0
}

func (x *CertificateRequest) GetBatchId() int64 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

func (x *CertificateRequest) GetBatchName() string {
	if x != nil {
		return x.BatchName
	}
	return ""
}

func (x *CertificateRequest) GetBatchStatus() string {
	if x != nil {
		return x.BatchStatus
	}
	return ""
}

func (x *CertificateRequest) GetOrderNumber() int64 {
	if x != nil {
		return x.OrderNumber
	}
	return 0
}

func (x *CertificateRequest) GetCreationDate() string {
	if x != nil {
		return x.CreationDate
	}
	return ""
}

func (x *CertificateRequest) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *CertificateRequest) GetRejectReason() string {
	if x != nil {
		return x.RejectReason
	}
	return ""
}

func (x *CertificateRequest) GetPkcs12Password() []byte {
	if x != nil {
		return x.Pkcs12Password
	}
	return nil
}

func (x *CertificateRequest) GetPkcs12Signature() []byte {
	if x != nil {
		return x.Pkcs12Signature
	}
	return nil
}

func (x *CertificateRequest) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *CertificateRequest) GetModified() string {
	if x != nil {
		return x.Modified
	}
	return ""
}

// GDSExtraData contains all GDS-specific extra data for a VASP record.
type GDSExtraData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Temporary: verification token for light weight authentication for verification
	// TODO: replace with admin API that uses authentication
	AdminVerificationToken string `protobuf:"bytes,1,opt,name=admin_verification_token,json=adminVerificationToken,proto3" json:"admin_verification_token,omitempty"`
}

func (x *GDSExtraData) Reset() {
	*x = GDSExtraData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gds_models_v1_models_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GDSExtraData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GDSExtraData) ProtoMessage() {}

func (x *GDSExtraData) ProtoReflect() protoreflect.Message {
	mi := &file_gds_models_v1_models_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GDSExtraData.ProtoReflect.Descriptor instead.
func (*GDSExtraData) Descriptor() ([]byte, []int) {
	return file_gds_models_v1_models_proto_rawDescGZIP(), []int{1}
}

func (x *GDSExtraData) GetAdminVerificationToken() string {
	if x != nil {
		return x.AdminVerificationToken
	}
	return ""
}

// GDSContactExtraData contains all GDS-specific extra data for a Contact record.
type GDSContactExtraData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Token for email verification
	Verified bool   `protobuf:"varint,1,opt,name=verified,proto3" json:"verified,omitempty"`
	Token    string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GDSContactExtraData) Reset() {
	*x = GDSContactExtraData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gds_models_v1_models_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GDSContactExtraData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GDSContactExtraData) ProtoMessage() {}

func (x *GDSContactExtraData) ProtoReflect() protoreflect.Message {
	mi := &file_gds_models_v1_models_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GDSContactExtraData.ProtoReflect.Descriptor instead.
func (*GDSContactExtraData) Descriptor() ([]byte, []int) {
	return file_gds_models_v1_models_proto_rawDescGZIP(), []int{2}
}

func (x *GDSContactExtraData) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *GDSContactExtraData) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_gds_models_v1_models_proto protoreflect.FileDescriptor

var file_gds_models_v1_models_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x67, 0x64, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x64,
	0x73, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xa8, 0x04, 0x0a, 0x12,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x61, 0x73, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x76, 0x61, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x67, 0x64, 0x73, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x26,
	0x0a, 0x0e, 0x70, 0x6b, 0x63, 0x73, 0x31, 0x32, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x70, 0x6b, 0x63, 0x73, 0x31, 0x32, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x6b, 0x63, 0x73, 0x31, 0x32,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0f, 0x70, 0x6b, 0x63, 0x73, 0x31, 0x32, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f,
	0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x22, 0x48, 0x0a, 0x0c, 0x47, 0x44, 0x53, 0x45, 0x78, 0x74,
	0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x18, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f,
	0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x47, 0x0a, 0x13, 0x47, 0x44, 0x53, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x45, 0x78,
	0x74, 0x72, 0x61, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2a, 0xa0, 0x01, 0x0a, 0x17, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c,
	0x49, 0x5a, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f,
	0x54, 0x4f, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50,
	0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x44,
	0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a,
	0x44, 0x4f, 0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0f, 0x0a, 0x0b, 0x43,
	0x52, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x0e, 0x0a, 0x0a,
	0x43, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x45, 0x44, 0x10, 0x07, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x69, 0x73, 0x61,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x64, 0x73, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gds_models_v1_models_proto_rawDescOnce sync.Once
	file_gds_models_v1_models_proto_rawDescData = file_gds_models_v1_models_proto_rawDesc
)

func file_gds_models_v1_models_proto_rawDescGZIP() []byte {
	file_gds_models_v1_models_proto_rawDescOnce.Do(func() {
		file_gds_models_v1_models_proto_rawDescData = protoimpl.X.CompressGZIP(file_gds_models_v1_models_proto_rawDescData)
	})
	return file_gds_models_v1_models_proto_rawDescData
}

var file_gds_models_v1_models_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gds_models_v1_models_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gds_models_v1_models_proto_goTypes = []interface{}{
	(CertificateRequestState)(0), // 0: gds.models.v1.CertificateRequestState
	(*CertificateRequest)(nil),   // 1: gds.models.v1.CertificateRequest
	(*GDSExtraData)(nil),         // 2: gds.models.v1.GDSExtraData
	(*GDSContactExtraData)(nil),  // 3: gds.models.v1.GDSContactExtraData
}
var file_gds_models_v1_models_proto_depIdxs = []int32{
	0, // 0: gds.models.v1.CertificateRequest.status:type_name -> gds.models.v1.CertificateRequestState
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gds_models_v1_models_proto_init() }
func file_gds_models_v1_models_proto_init() {
	if File_gds_models_v1_models_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gds_models_v1_models_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertificateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gds_models_v1_models_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GDSExtraData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gds_models_v1_models_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GDSContactExtraData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gds_models_v1_models_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gds_models_v1_models_proto_goTypes,
		DependencyIndexes: file_gds_models_v1_models_proto_depIdxs,
		EnumInfos:         file_gds_models_v1_models_proto_enumTypes,
		MessageInfos:      file_gds_models_v1_models_proto_msgTypes,
	}.Build()
	File_gds_models_v1_models_proto = out.File
	file_gds_models_v1_models_proto_rawDesc = nil
	file_gds_models_v1_models_proto_goTypes = nil
	file_gds_models_v1_models_proto_depIdxs = nil
}
