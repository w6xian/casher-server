package store

type CompInfo struct {
	Id      int64  `json:"id"`
	ProxyId int64  `json:"proxy_id"`
	ComApi  string `json:"com_api"`
	ComPem  string `json:"com_pem"`
}
