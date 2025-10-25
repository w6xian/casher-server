package store

import (
	"casher-server/internal/i18n"
	"casher-server/internal/lager"
	"context"
	"fmt"
)

type CompanySnReq struct {
	Req
	Sn string `json:"sn"`
}

func (req *CompanySnReq) Validate() error {
	if req.Sn == "" {
		return req.Tracker.Error("get_company_by_sn_validate", "请输入公司编码")
	}
	return nil
}

type CompanyNameReq struct {
	Req
	Name string `json:"name"`
}

func (req *CompanyNameReq) Validate() error {
	if req.Name == "" {
		return req.Tracker.Error("get_company_by_name_validate", "请输入公司名称")
	}
	return nil
}

type CompanyReply struct {
	Req
	UnionId       string `json:"union_id"`
	Province      string `json:"province"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	Pinyin        string `json:"pinyin"`
	PinyinIdx     string `json:"pinyin_idx"`
	SnCode        string `json:"sn_code"`
	LegalPerson   string `json:"legal_person"`
	ContactName   string `json:"contact_name"`
	ContactPhone  string `json:"contact_phone"`
	ContactMobile string `json:"contact_mobile"`
	ContactEmail  string `json:"contact_email"`
	RegisterDate  string `json:"register_date"`
	ExpireDate    string `json:"expire_date"`
	AreaCode      string `json:"area_code"`
	AreaPath      string `json:"area_path"`
	AreaStreet    string `json:"area_street"`
	Address       string `json:"address"`
	Site          string `json:"site"`
	Scale         string `json:"scale"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
	GeoHash       string `json:"geo_hash"`
}

type CompanyLiteModel struct {
	Sn          string `json:"sn"`
	LegalPerson string `json:"legal_person"`
	Name        string `json:"name"`
	Province    string `json:"province"`
}

type CompanyModel struct {
	Id            int64  `json:"id"`
	UnionId       string `json:"union_id"`
	Province      string `json:"province"`
	Name          string `json:"name"`
	Avatar        string `json:"avatar"`
	Pinyin        string `json:"pinyin"`
	PinyinIdx     string `json:"pinyin_idx"`
	SnCode        string `json:"sn_code"`
	LegalPerson   string `json:"legal_person"`
	ContactName   string `json:"contact_name"`
	ContactPhone  string `json:"contact_phone"`
	ContactMobile string `json:"contact_mobile"`
	ContactEmail  string `json:"contact_email"`
	RegisterDate  string `json:"register_date"`
	ExpireDate    string `json:"expire_date"`
	AreaCode      string `json:"area_code"`
	AreaPath      string `json:"area_path"`
	AreaStreet    string `json:"area_street"`
	Address       string `json:"address"`
	Site          string `json:"site"`
	Scale         string `json:"scale"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
	GeoHash       string `json:"geo_hash"`
}

func (s *Store) GetCluldCompaniesBySn(ctx context.Context, req *CompanySnReq, reply *CompanyReply) error {
	lang := req.Tracker
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("SearchCluldCompaniesBySn", "SearchCluldCompaniesBySn", "Supp")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	fmt.Println("SearchCluldCompaniesBySn link=", link)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 3 查询公司信息
	res, err := db.GetPublicCompanyBySn(link, req.Sn)
	if err != nil {
		log.ErrorExit("SearchCluldCompaniesBySn Query err", err)
		return lang.Error("search_company_by_sn", "查询公司信息失败:{{.error}}", i18n.String("error", err.Error()))
	}
	reply.AppId = req.AppId
	reply.UnionId = res.UnionId
	reply.Province = res.Province
	reply.Name = res.Name
	reply.Avatar = res.Avatar
	reply.Pinyin = res.Pinyin
	reply.PinyinIdx = res.PinyinIdx
	reply.SnCode = res.SnCode
	reply.LegalPerson = res.LegalPerson
	reply.ContactName = res.ContactName
	reply.ContactPhone = res.ContactPhone
	reply.ContactMobile = res.ContactMobile
	reply.ContactEmail = res.ContactEmail
	reply.RegisterDate = res.RegisterDate
	reply.ExpireDate = res.ExpireDate
	reply.AreaCode = res.AreaCode
	reply.AreaPath = res.AreaPath
	reply.AreaStreet = res.AreaStreet
	reply.Address = res.Address
	reply.Site = res.Site
	reply.Scale = res.Scale
	reply.Longitude = res.Longitude
	reply.Latitude = res.Latitude
	reply.GeoHash = res.GeoHash
	return nil
}

func (s *Store) GetCluldCompaniesByName(ctx context.Context, req *CompanyNameReq, reply *CompanyReply) error {
	lang := req.Tracker
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("GetCluldCompaniesByName", "SearchCluldCompaniesBySn", "Supp")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	fmt.Println("GetCluldCompaniesByName link=", link)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 3 查询公司信息
	res, err := db.GetPublicCompanyByName(link, req.Name)
	if err != nil {
		log.ErrorExit("GetCluldCompaniesByName Query err", err)
		return lang.Error("search_company_by_name", "查询公司信息失败:{{.error}}", i18n.String("error", err.Error()))
	}
	reply.AppId = req.AppId
	reply.UnionId = res.UnionId
	reply.Province = res.Province
	reply.Name = res.Name
	reply.Avatar = res.Avatar
	reply.Pinyin = res.Pinyin
	reply.PinyinIdx = res.PinyinIdx
	reply.SnCode = res.SnCode
	reply.LegalPerson = res.LegalPerson
	reply.ContactName = res.ContactName
	reply.ContactPhone = res.ContactPhone
	reply.ContactMobile = res.ContactMobile
	reply.ContactEmail = res.ContactEmail
	reply.RegisterDate = res.RegisterDate
	reply.ExpireDate = res.ExpireDate
	reply.AreaCode = res.AreaCode
	reply.AreaPath = res.AreaPath
	reply.AreaStreet = res.AreaStreet
	reply.Address = res.Address
	reply.Site = res.Site
	reply.Scale = res.Scale
	reply.Longitude = res.Longitude
	reply.Latitude = res.Latitude
	reply.GeoHash = res.GeoHash
	return nil
}
