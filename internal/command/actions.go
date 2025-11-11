package command

const (
	// 回包操作
	ACTION_CALL  = -0xFF // 别名
	ACTION_REPLY = -0xFE //
	// 无效操作
	ACTION_INVALID = 0x00
	// 登录操作
	ACTION_LOGIN = 0xFE
	// 登出操作
	ACTION_LOGOUT = 0xFD
	//广播
	ACTION_BROADCAST = 0xFF
	//新建订单
	ACTION_NOTICE_ORDER = 0x11
)
