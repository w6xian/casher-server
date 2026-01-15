package action

const (
	// 回包操作
	CALL  = -0xFF // 别名
	REPLY = -0xFE //
	// 无效操作
	INVALID = 0x00
	// 登录操作
	LOGIN = 0xFE
	// 登出操作
	LOGOUT = 0xFD
	//广播
	BROADCAST = 0xFF
	//新建订单
	NOTICE_ORDER = 0x11
)
