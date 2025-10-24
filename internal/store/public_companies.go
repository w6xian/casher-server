package store

import (
	"casher-server/internal/i18n"
	"casher-server/internal/lager"
	"context"
	"fmt"
)

type CompanyReq struct {
	Req
	Sn   string `json:"sn"`
	Ts   int64  `json:"ts"`
	Sign string `json:"sign"`
}

type CompanyReqReply struct {
	AppId string `json:"-"`
	Sn    string `json:"sn"`
	Ts    int64  `json:"ts"`
	Sign  string `json:"sign"`
}

// 实现 IEncrypt
func (reply *CompanyReqReply) EncryptInfo() string {
	return reply.AppId
}
func (reply *CompanyReqReply) SetSign(sign string, ts int64) error {
	reply.Sign = sign
	reply.Ts = ts
	return nil
}

type CompanyLiteModel struct {
	Id   int64  `json:"id"`
	Sn   string `json:"sn"`
	Name string `json:"name"`
}

type CompanyModel struct {
	Id      int64  `json:"id"`
	OpenId  string `json:"open_id"`
	ShopId  int64  `json:"shop_id"`
	Sn      string `json:"sn"`
	Avatar  string `json:"avatar"`
	Cover   string `json:"cover"`
	Name    string `json:"name"`
	Pinyin  string `json:"pinyin"`
	Feature string `json:"feature"`

	Style     int64 `json:"style"`
	StyleType int64 `json:"style_type"`
	Num       int64 `json:"num,omitempty"`
	Total     int64 `json:"total,omitempty"`

	Source     int64 `json:"source"`
	Price      int64 `json:"price"`
	Type       int64 `json:"type"`
	LinePrice  int64 `json:"line_price"`
	SaleNum    int64 `json:"sale_num"`
	Times      int64 `json:"times"`
	PkWeight   int64 `json:"pk_weight"`
	SpecId     int64 `json:"spec_id"`
	SpecWeight int64 `json:"spec_weight"`
	// 200g*20袋 = 箱
	Spec     int64    `json:"spec"`
	Unit     string   `json:"unit"`
	PkAmount int64    `json:"pk_amount"`
	SpecName string   `json:"spec_name"`
	PackName string   `json:"pack_name"`
	Units    []*Units `json:"units" ignore:"i"`
	// 库存
	Stock int64  `json:"stock"`
	Tags  string `json:"tags"`
	Sort  int64  `json:"sort"`
	// 内部
	Mark   string `json:"mark"`
	Status int64  `json:"status"`
	// 经手人
	HandlerId   int64  `json:"handler_id"`
	HandlerName string `json:"handler_name"`
	// 品牌
	BrandId   int64  `json:"brand_id"`
	BrandName string `json:"brand_name"`
	// 单价
	UnitPrice int64 `json:"unit_price"`
	PackPrice int64 `json:"pack_price"`
	// 成本
	Cost int64 `json:"cost"`
	//保质期（天）
	KeepLife int64 `json:"keep_life"`
	// 保质期单位
	KeepLifeUnit string `json:"keep_life_unit"`
	// 分类
	CategoryId   int64   `json:"category_id,omitempty"`
	CategoryName string  `json:"category_name"`
	SupplierId   int64   `json:"supplier_id,omitempty"`
	SupplierName string  `json:"supplier_name"`
	SuppIds      []int64 `json:"supplier_ids,omitempty"`
	Intime       int64   `json:"intime,omitempty"`
}

func (s *Store) SearchCluldCompaniesBySn(ctx context.Context, req *CompanyReq, reply *CompanyReqReply) error {
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
	companyModel, err := db.GetPublicCompanyBySn(link, req.Sn)
	if err != nil {
		log.ErrorExit("SearchCluldCompaniesBySn Query err", err)
		return lang.Error("search_company_by_sn", "查询公司信息失败:{{.error}}", i18n.String("error", err.Error()))
	}
	reply.Sn = companyModel.Sn
	return nil
}
