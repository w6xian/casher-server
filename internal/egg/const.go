package egg

const (
	EGG_TYPE_PING            uint8 = 0x00
	EGG_TYPE_PONG            uint8 = 0xFF
	EGG_TYPE_SESSION         uint8 = 0xFE
	EGG_TYPE_CLOSE           uint8 = 0xFD
	EGG_TYPE_OAUTH           uint8 = 0xFC
	EGG_TYPE_REGISTER        uint8 = 0xFB
	EGG_TYPE_ROGER           uint8 = 0xFA
	EGG_TYPE_ERROR           uint8 = 0xF9
	EGG_TYPE_STRING          uint8 = 0xEF
	EGG_TYPE_EVENT_NAME      uint8 = 0xEF
	EGG_TYPE_INT32           uint8 = 0xEE
	EGG_TYPE_INT64           uint8 = 0xED
	EGG_TYPE_UINT32          uint8 = 0xEC
	EGG_TYPE_UINT64          uint8 = 0xEB
	EGG_TYPE_JSON            uint8 = 0xEA
	EGG_TYPE_EVENT_ARGUMENTS uint8 = 0xEA
	EGG_TYPE_BIN             uint8 = 0xE9
	EGG_TYPE_INT             uint8 = 0xE8
	EGG_TYPE_UINT            uint8 = 0xE7

	EGG_TYPE_COMMAND  uint8 = 0xD7
	EGG_TYPE_RESPONSE uint8 = 0xD6

	// 打印服务
	EGG_TYPE_SETTING      uint8 = 0x10
	EGG_TYPE_SETTING_BACK uint8 = 0x1F
	EGG_TYPE_PRINTER      uint8 = 0x11
	EGG_TYPE_ORDER        uint8 = 0x12
	EGG_TYPE_TODO         uint8 = 0x13
	EGG_TYPE_MESSAGE      uint8 = 0x14

	EGG_TYPE_SHAKE          uint8 = 0x0F
	EGG_TYPE_HEADERS        uint8 = 0x0E
	EGG_TYPE_EVENT_MAP      uint8 = 0x0D
	EGG_TYPE_EVENT_LIST     uint8 = 0x0C
	EGG_TYPE_EVENT_BACK_MAP uint8 = 0x0B
	EGG_TYPE_CRYPTO         uint8 = 0x01

	// same
	EGG_TYPE_VALUES uint8  = 0x02
	TLAV_HEAD_LEN   int    = 9 // tlav+1t+3l+1a+nv
	TLAV            string = "TLAV"
)
