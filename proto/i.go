package proto

type IAuthInfo interface {
	// GetUserIds proxyId, shopId, userId 获取用户ID
	GetUserIds() (int64, int64, int64)
}
