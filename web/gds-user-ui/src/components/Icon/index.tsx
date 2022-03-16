import { Icon, createIcon, IconProps } from '@chakra-ui/icon';
import { FC } from 'react';

const MenuIcon = ({ color }: any): any => (
  <svg width="24px" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
    <title>Menu</title>
    <path fill={color || 'white'} d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z" />
  </svg>
);

const CloseIcon = ({ color }: any) => (
  <svg width="24" viewBox="0 0 18 18" xmlns="http://www.w3.org/2000/svg">
    <title>Close</title>
    <path
      fill={color || 'white'}
      d="M9.00023 7.58599L13.9502 2.63599L15.3642 4.04999L10.4142 8.99999L15.3642 13.95L13.9502 15.364L9.00023 10.414L4.05023 15.364L2.63623 13.95L7.58623 8.99999L2.63623 4.04999L4.05023 2.63599L9.00023 7.58599Z"
    />
  </svg>
);

const PlugIcon = (props: IconProps) => {
  return (
    <Icon width="20" height="20" viewBox="0 0 106 96" fill="none">
      <path
        d="M39.75 0C40.6285 0 41.4711 0.316071 42.0923 0.87868C42.7135 1.44129 43.0625 2.20435 43.0625 3V18H62.9375V3C62.9375 2.20435 63.2865 1.44129 63.9077 0.87868C64.5289 0.316071 65.3715 0 66.25 0C67.1285 0 67.9711 0.316071 68.5923 0.87868C69.2135 1.44129 69.5625 2.20435 69.5625 3V18H76.1875C77.066 18 77.9086 18.3161 78.5298 18.8787C79.151 19.4413 79.5 20.2044 79.5 21V39C79.5 44.5695 77.057 49.911 72.7085 53.8492C68.36 57.7875 62.4622 60 56.3125 60C56.2993 62.604 56.2463 65.07 56.0475 67.32C55.7759 70.404 55.2127 73.338 53.9474 75.864C52.6645 78.5355 50.3915 80.7171 47.5211 82.032C44.5531 83.4 40.8696 84 36.4375 84C29.8258 84 25.7712 85.98 23.3597 88.308C21.1825 90.3904 19.9401 93.1328 19.875 96H13.25C13.25 92.304 14.787 87.798 18.5301 84.192C22.3527 80.52 28.2291 78 36.4375 78C40.2866 78 42.8107 77.472 44.5068 76.686C46.1034 75.948 47.1567 74.886 47.912 73.386C48.7136 71.79 49.1906 69.66 49.4357 66.846C49.6146 64.806 49.6676 62.538 49.6809 60C43.5323 59.9984 37.6362 57.7852 33.2891 53.8471C28.9421 49.909 26.5 44.5685 26.5 39V21C26.5 20.2044 26.849 19.4413 27.4702 18.8787C28.0914 18.3161 28.934 18 29.8125 18H36.4375V3C36.4375 2.20435 36.7865 1.44129 37.4077 0.87868C38.0289 0.316071 38.8715 0 39.75 0V0ZM33.125 24V39C33.125 42.9782 34.87 46.7936 37.976 49.6066C41.0821 52.4196 45.2949 54 49.6875 54H56.3125C60.7051 54 64.9179 52.4196 68.024 49.6066C71.13 46.7936 72.875 42.9782 72.875 39V24H33.125Z"
        fill="white"
      />
    </Icon>
  );
};

const CircleIcon = (props: IconProps) => {
  return (
    <Icon width="16" height="16" viewBox="0 0 16 16" fill="none">
      <path
        d="M8 0.25C3.71875 0.25 0.25 3.71875 0.25 8C0.25 12.2812 3.71875 15.75 8 15.75C12.2812 15.75 15.75 12.2812 15.75 8C15.75 3.71875 12.2812 0.25 8 0.25ZM8 14.25C4.54688 14.25 1.75 11.4531 1.75 8C1.75 4.54688 4.54688 1.75 8 1.75C11.4531 1.75 14.25 4.54688 14.25 8C14.25 11.4531 11.4531 14.25 8 14.25Z"
        fill="#C1C9D2"
      />
    </Icon>
  );
};

const NetworkIcon = (props: IconProps) => {
  return (
    <Icon width="15" height="20" viewBox="0 0 92 86" fill="none">
      <path
        d="M47.2703 5.53103L48.3322 5.58779C60.2626 6.22554 70.7769 11.725 77.6574 20.0163L78.359 20.8618L77.2606 20.8354C75.4748 20.7926 73.6975 20.7823 71.9278 20.8071C68.8553 20.8547 65.7857 21.0068 62.725 21.263L62.3928 21.2908L62.2394 20.9948C61.7195 19.9917 60.9055 19.1409 59.886 18.5435C58.8664 17.9459 57.6846 17.6273 56.4761 17.6269L47.2703 5.53103ZM47.2703 5.53103L47.904 6.38503M47.2703 5.53103L47.904 6.38503M47.904 6.38503C50.5338 9.92919 52.9305 13.6194 55.0809 17.4352L55.2467 17.7294M47.904 6.38503L55.2467 17.7294M55.2467 17.7294L55.5816 17.6854M55.2467 17.7294L55.5816 17.6854M55.5816 17.6854C55.8778 17.6465 56.1767 17.6269 56.4759 17.6269L55.5816 17.6854ZM45.9995 80.48C51.254 80.4853 56.4588 79.5229 61.3169 77.6471L61.5809 77.5452L61.6294 77.2663C63.1936 68.2667 63.5034 59.483 62.4907 50.9093L62.4482 50.5486L62.092 50.4776C60.6584 50.1917 59.3786 49.4579 58.4611 48.4051L58.2301 48.14L57.9024 48.2678C49.4271 51.5755 41.3603 55.7652 33.6708 60.7911L33.3684 60.9887L33.461 61.3379C35.1485 67.697 37.8335 73.9382 41.5477 80.0513L41.6781 80.2659L41.9281 80.2895C43.2665 80.4156 44.6252 80.48 45.9995 80.48ZM45.9995 80.48L46 79.98M45.9995 80.48C45.9997 80.48 45.9998 80.48 46 80.48V79.98M46 79.98C44.6412 79.98 43.298 79.9164 41.975 79.7917C38.284 73.7166 35.6187 67.5195 33.9443 61.2096C41.6061 56.2019 49.6423 52.0283 58.0842 48.7336C59.0772 49.873 60.4566 50.6613 61.9942 50.9679C63.0007 59.4896 62.6934 68.2247 61.1368 77.1807C56.3368 79.034 51.1933 79.9852 46 79.98ZM43.1074 5.84763L42.9412 5.63537L42.6726 5.65758C42.6105 5.66271 42.5475 5.66672 42.4788 5.67107L42.4766 5.67121C42.4102 5.67542 42.3383 5.67998 42.2659 5.68596L41.9916 5.70864L41.864 5.95261C38.746 11.9166 36.3045 17.8147 34.5583 23.6406L34.4449 24.019L34.7868 24.2169C35.4783 24.6171 36.0799 25.1355 36.5606 25.743L36.7677 26.0047L37.0888 25.9139C41.2586 24.7347 45.4897 23.7538 49.7659 22.9749L50.1056 22.913L50.1681 22.5735C50.378 21.433 50.942 20.3704 51.7958 19.5192L52.0651 19.2506L51.8782 18.9193C49.3832 14.4986 46.4683 10.1399 43.1074 5.84763ZM37.2381 7.25457L37.695 6.3208L36.6803 6.54675C26.9747 8.70796 18.4759 14.1847 12.8583 21.9112L12.4062 22.533L13.1579 22.6941C17.3919 23.6017 21.5836 24.6738 25.722 25.9076L26.0462 26.0043L26.2565 25.7393C26.7726 25.0891 27.4269 24.5418 28.1808 24.1327C28.9347 23.7236 29.7716 23.4617 30.6402 23.3643L30.9706 23.3272L31.0643 23.0082C32.5958 17.7888 34.6551 12.533 37.2381 7.25457ZM80.8519 24.4981L80.7141 24.2712L74.6221 24.5943L74.6251 24.0943C74.6251 24.0943 74.6251 24.0943 74.625 24.0943C70.7629 24.071 66.9424 24.2186 63.1571 24.5259L62.7983 24.5551L62.7121 24.9046C62.4045 26.1522 61.6709 27.2792 60.6137 28.113L60.3209 28.3439L60.4587 28.6904C61.7205 31.8616 62.7946 35.0952 63.6761 38.3765L63.7657 38.7102L64.1096 38.7443C64.95 38.8278 65.7634 39.0651 66.5033 39.4417C67.2431 39.8182 67.8941 40.3263 68.4202 40.9352L68.6277 41.1754L68.9334 41.0898C74.1891 39.6178 79.588 38.4499 85.1338 37.6068L85.6371 37.5302L85.5515 37.0284C84.7998 32.6209 83.2076 28.3768 80.8519 24.4981ZM10.8942 25.6199L10.5329 25.5455L10.3552 25.8688C8.33668 29.543 6.98523 33.5065 6.35703 37.5973L6.20499 38.5874L7.08766 38.1137C12.8586 35.0169 18.7266 32.3151 24.6935 30.0241L25.0274 29.8959L25.014 29.5385L25.0131 29.5148L25.0011 29.1553L24.6564 29.0524C20.1944 27.7208 15.6004 26.5882 10.8942 25.6199ZM50.8702 26.4432L50.6929 26.1435L50.3502 26.206C46.2409 26.9552 42.1755 27.8973 38.1674 29.0288L37.8183 29.1274L37.8037 29.49C37.801 29.5573 37.797 29.6247 37.7919 29.6919L37.7636 30.0594L38.1063 30.1951C45.2106 33.0088 52.0342 36.4044 58.4975 40.3422L58.8077 40.5312L59.0833 40.2946C59.3624 40.0551 59.6641 39.8389 59.985 39.6487L60.3119 39.455L60.2128 39.0882C59.3659 35.9535 58.3352 32.8648 57.1253 29.8362L56.9947 29.5095L56.6431 29.5221C56.5863 29.5241 56.5295 29.5254 56.4727 29.526C55.3242 29.5252 54.1988 29.2369 53.2134 28.6935C52.2274 28.1498 51.4195 27.3716 50.8702 26.4432ZM26.5897 33.1845L26.3594 32.9397L26.0458 33.0605C19.3133 35.654 12.7016 38.7923 6.19883 42.4445L5.94368 42.5878V42.8805C5.94368 42.8872 5.94338 42.8955 5.94238 42.917C5.94154 42.935 5.94 42.9663 5.94 43C5.94 52.9172 10.0405 61.9091 16.728 68.5996L17.0514 68.9232L17.4024 68.6297C21.1536 65.4932 25.0677 62.5309 29.1306 59.7536L29.4082 59.5639L29.3371 59.2352C27.6206 51.2982 27.3705 43.1819 28.4965 34.945L28.5438 34.5992L28.2355 34.4356C27.6147 34.1062 27.0587 33.6831 26.5897 33.1845ZM36.6405 33.1988L36.3346 33.078L36.1026 33.311C35.1223 34.2954 33.8099 34.9459 32.376 35.1503L32.002 35.2036L31.9511 35.578C30.9755 42.7612 31.114 49.826 32.4217 56.7598L32.5583 57.4841L33.1808 57.0894C40.6596 52.347 48.5003 48.3585 56.7124 45.1671L57.0324 45.0427L57.0312 44.6994V44.6624L57.0312 44.6618C57.0319 44.3741 57.055 44.0867 57.1004 43.802L57.1531 43.472L56.8683 43.2972C50.4661 39.3657 43.6957 35.9858 36.6405 33.1988ZM86.0163 41.3762L85.9912 40.8194L85.4404 40.9045C80.291 41.7005 75.2001 42.7985 70.198 44.1919L69.8331 44.2936L69.8322 44.6724C69.8301 45.5295 69.6287 46.3773 69.2405 47.1572L69.0566 47.5267L69.3846 47.7773C73.7474 51.11 77.8704 54.7074 81.7269 58.5463L82.2251 59.0422L82.5301 58.4088C84.8623 53.5657 86.066 48.3125 86.06 42.9986C86.06 42.4508 86.0403 41.9089 86.0163 41.3762ZM66.8337 50.1075L66.0307 49.4958V50.5052V50.5069V50.5366L66.0342 50.5661C66.9764 58.4734 66.835 66.5489 65.6586 74.7663L65.5161 75.7612L66.3955 75.2747C72.0539 72.1447 76.8065 67.7609 80.2325 62.5061L80.4516 62.1701L80.1711 61.8834C76.0383 57.6589 71.5806 53.7231 66.8337 50.1075ZM30.4779 63.7439L30.2835 63.0702L29.7099 63.4735C26.4388 65.7735 23.2367 68.2328 20.1017 70.8381L19.6282 71.2315L20.1104 71.6143C24.7428 75.2911 30.2102 77.9283 36.0947 79.327L37.1899 79.5873L36.6487 78.6002C33.9803 73.7326 31.9301 68.7752 30.4779 63.7439ZM0.5 43C0.5 19.5946 20.8767 0.5 46 0.5C71.1233 0.5 91.5 19.5946 91.5 43C91.5 66.4054 71.1233 85.5 46 85.5C20.8767 85.5 0.5 66.4054 0.5 43Z"
        fill="white"
        stroke="white"
      />
    </Icon>
  );
};
const GoogleIcon = (props: { h?: number; w?: number }) => {
  return (
    <svg
      width={props.w}
      height={props.h}
      viewBox="-3 0 262 262"
      xmlns="http://www.w3.org/2000/svg"
      preserveAspectRatio="xMidYMid">
      <path
        d="M255.878 133.451c0-10.734-.871-18.567-2.756-26.69H130.55v48.448h71.947c-1.45 12.04-9.283 30.172-26.69 42.356l-.244 1.622 38.755 30.023 2.685.268c24.659-22.774 38.875-56.282 38.875-96.027"
        fill="#4285F4"
      />
      <path
        d="M130.55 261.1c35.248 0 64.839-11.605 86.453-31.622l-41.196-31.913c-11.024 7.688-25.82 13.055-45.257 13.055-34.523 0-63.824-22.773-74.269-54.25l-1.531.13-40.298 31.187-.527 1.465C35.393 231.798 79.49 261.1 130.55 261.1"
        fill="#34A853"
      />
      <path
        d="M56.281 156.37c-2.756-8.123-4.351-16.827-4.351-25.82 0-8.994 1.595-17.697 4.206-25.82l-.073-1.73L15.26 71.312l-1.335.635C5.077 89.644 0 109.517 0 130.55s5.077 40.905 13.925 58.602l42.356-32.782"
        fill="#FBBC05"
      />
      <path
        d="M130.55 50.479c24.514 0 41.05 10.589 50.479 19.438l36.844-35.974C195.245 12.91 165.798 0 130.55 0 79.49 0 35.393 29.301 13.925 71.947l42.211 32.783c10.59-31.477 39.891-54.251 74.414-54.251"
        fill="#EB4335"
      />
    </svg>
  );
};
const OpenSourceIcon: FC = (props: IconProps) => {
  return (
    <Icon width="15" height="20" viewBox="0 0 98 85" fill="none">
      <g clipPath="url(#clip0_65_1708)">
        <path
          d="M-0.416504 44.3807C0.0489961 21.0986 18.8374 4.52365 41.8368 1.84084C69.2737 -1.344 92.3527 14.8432 96.7872 36.3163C100.998 56.6818 88.0928 75.7457 66.8146 83.3C64.9801 83.9482 63.9756 83.5896 63.2743 81.9719L52.6995 58.148C52.0962 56.7482 52.5066 55.8769 54.0868 55.258C58.8735 53.3747 61.8533 50.2218 62.6526 45.6955C62.9503 44.0358 62.8515 42.3407 62.3624 40.7155C61.8733 39.0903 61.0043 37.5695 59.8094 36.2477C58.6145 34.9259 57.1192 33.8312 55.4162 33.0314C53.7132 32.2317 51.8389 31.7441 49.9096 31.5988C46.3359 31.2987 42.7644 32.1895 39.9204 34.0903C37.0765 35.9911 35.1732 38.7595 34.5971 41.8333C33.5589 47.4938 36.7806 52.7638 42.8566 55.1571C44.7554 55.9035 45.1382 56.6366 44.4032 58.3127L33.7794 82.3013C33.2679 83.478 32.1532 83.8951 30.7169 83.409C19.5025 79.6036 10.3019 72.3423 4.87856 63.0169C-0.0490035 54.578 -0.0949414 47.9799 -0.416504 44.3674V44.3807ZM3.71481 44.0513C3.7975 45.1855 3.84037 46.4977 3.97206 47.8418C5.21237 60.1668 13.7629 72.4386 29.8563 79.2307C30.4964 79.4777 30.7352 79.3688 30.9741 78.8455C33.9111 72.1332 36.8756 65.4261 39.8584 58.7138C40.1126 58.1613 39.984 57.8983 39.3929 57.5796C33.1424 54.161 30.0309 49.2629 30.38 42.8321C30.5729 39.1372 32.2175 35.886 35.0871 33.174C41.0038 27.5719 50.3016 26.3527 57.9119 30.1591C60.9861 31.7036 63.4706 33.9961 65.0579 36.7528C66.6451 39.5096 67.2655 42.6095 66.8421 45.669C66.1224 50.8965 63.0324 54.9101 57.7986 57.6221C57.2718 57.8983 57.1126 58.1188 57.3545 58.6554C60.3527 65.381 63.3356 72.104 66.2847 78.8296C66.5267 79.3688 66.7656 79.4777 67.3872 79.2147C74.4218 76.3779 80.2742 72.3111 84.8067 66.8897C91.679 58.6155 94.5179 49.4143 93.1612 39.2887C90.4662 19.1437 69.5708 2.45975 42.6637 5.33912C21.7039 7.56771 4.06393 22.9102 3.71175 44.0486L3.71481 44.0513Z"
          fill="white"
        />
      </g>
      <defs>
        <clipPath id="clip0_65_1708">
          <rect width="98" height="85" fill="white" />
        </clipPath>
      </defs>
    </Icon>
  );
};

const SecureIcon: FC = (props: IconProps) => {
  return (
    <Icon width="40" height="20" viewBox="0 0 120 110" fill="none">
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M37.5 38.9583C37.5 33.4882 39.8705 28.2422 44.0901 24.3743C48.3097 20.5063 54.0326 18.3333 60 18.3333C65.9674 18.3333 71.6903 20.5063 75.9099 24.3743C80.1295 28.2422 82.5 33.4882 82.5 38.9583V41.25H87.5V38.9583C87.5 25.0365 75.1875 13.75 60 13.75C44.8125 13.75 32.5 25.0365 32.5 38.9583V41.25H37.5V38.9583ZM77.5 41.25V38.9583C77.5 36.8517 77.0474 34.7657 76.1679 32.8195C75.2884 30.8732 73.9994 29.1048 72.3744 27.6152C70.7493 26.1256 68.8202 24.9439 66.697 24.1378C64.5738 23.3316 62.2981 22.9167 60 22.9167C57.7019 22.9167 55.4262 23.3316 53.303 24.1378C51.1798 24.9439 49.2507 26.1256 47.6256 27.6152C46.0006 29.1048 44.7116 30.8732 43.8321 32.8195C42.9526 34.7657 42.5 36.8517 42.5 38.9583V41.25H47.5V38.9583C47.5 35.9194 48.817 33.0049 51.1612 30.8561C53.5054 28.7072 56.6848 27.5 60 27.5C63.3152 27.5 66.4946 28.7072 68.8388 30.8561C71.183 33.0049 72.5 35.9194 72.5 38.9583V41.25H77.5Z"
        fill="white"
      />
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M42.5 45.8333H92.5V91.6667H27.5V45.8333H37.5V41.25H27.5C26.1739 41.25 24.9021 41.7329 23.9645 42.5924C23.0268 43.452 22.5 44.6178 22.5 45.8333V91.6667C22.5 92.8822 23.0268 94.048 23.9645 94.9076C24.9021 95.7671 26.1739 96.25 27.5 96.25H92.5C93.8261 96.25 95.0979 95.7671 96.0355 94.9076C96.9732 94.048 97.5 92.8822 97.5 91.6667V45.8333C97.5 44.6178 96.9732 43.452 96.0355 42.5924C95.0979 41.7329 93.8261 41.25 92.5 41.25H42.5V45.8333Z"
        fill="white"
      />
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M82.5 59.5833H37.5V55H82.5V59.5833Z"
        fill="white"
      />
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M82.5 71.0416H37.5V66.4583H82.5V71.0416Z"
        fill="white"
      />
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M82.5 82.5001H37.5V77.9167H82.5V82.5001Z"
        fill="white"
      />
    </Icon>
  );
};

const getIcon = (slug: any): any => {
  switch (slug) {
    case 'secure':
      return <SecureIcon />;
    case 'network':
      return <NetworkIcon />;
    case 'opensource':
      return <OpenSourceIcon />;
    case 'plug':
      return <PlugIcon />;
    default:
      return <SecureIcon />;
  }
};
export {
  MenuIcon,
  CloseIcon,
  SecureIcon,
  NetworkIcon,
  OpenSourceIcon,
  PlugIcon,
  getIcon,
  GoogleIcon,
  CircleIcon
};
