package store

type ProxyInfo struct {
	ProxyId       int64  `json:"com_id"`
	Avatar        string `json:"avatar"`
	Name          string `json:"name"`
	SortName      string `json:"sort_name"`
	SnCode        string `json:"sn_code"`
	LegalPersion  string `json:"legal_persion"`
	ContactName   string `json:"contact_name"`
	ContactPhone  string `json:"contact_phone"`
	ContactMobile string `json:"contact_mobile"`
	ContactEmail  string `json:"contact_email"`
	RegisterDate  string `json:"register_date"`
	ExpireDate    string `json:"expire_date"`
	AreaCode      string `json:"area_code"`
	AreaPath      string `json:"area_path"`
	AreaStreet    string `json:"area_street"`
	Street        string `json:"street"`
	Address       string `json:"address"`
	CreateTime    string `json:"intime"`
}
