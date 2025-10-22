package store

type CompInfo struct {
	Id      int64  `json:"id"`
	ProxyId int64  `json:"proxy_id"`
	ComApi  string `json:"com_api"`
	ComPem  string `json:"com_pem"`
}

type Admin struct {
	ProxyId   int64  `json:"proxy_id"`
	FailTimes int64  `json:"fail_times"`
	Mobile    string `json:"mobile"`
	Password  string `json:"password"`
}
