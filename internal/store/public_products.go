package store

import (
	"casher-server/internal/errors"
	"casher-server/internal/lager"
	"context"
)

type PrdSnReq struct {
	Req
	Sn   string `json:"sn"`
	Ts   int64  `json:"ts"`
	Sign string `json:"sign"`
}

func (req *PrdSnReq) DecryptInfo() (string, int64) {
	return req.Sign, req.Ts
}

func (req *PrdSnReq) Validate() error {
	if req.Sn == "" {
		return errors.New("sn required")
	}
	return nil
}

type PrdSnReqReply struct {
	AppId string `json:"-"`
	Sn    string `json:"sn"`
	Ts    int64  `json:"ts"`
	Sign  string `json:"sign"`
}

// 实现 IEncrypt
func (reply *PrdSnReqReply) EncryptInfo() string {
	return reply.AppId
}
func (reply *PrdSnReqReply) SetSign(sign string, ts int64) error {
	reply.Sign = sign
	reply.Ts = ts
	return nil
}

/*
*
major:false
mark:"最小单位"
pack_a_price:0
pack_m_price:0
pack_name:"袋"
pk_amount:1   //style为2时，pk_amount为 1*1000
pk_weight:102
spec_name:"袋"
*/
type Units struct {
	Major      bool   `json:"major"`
	Mark       string `json:"mark"`
	PackMPrice int64  `json:"pack_m_price"`
	PackName   string `json:"pack_name"`
	PkAmount   int64  `json:"pk_amount"`
	PkWeight   int64  `json:"pk_weight"`
	SpecName   string `json:"spec_name"`
}

type Product struct {
	Id      int64  `json:"id"`
	Sn      string `json:"sn"`
	Name    string `json:"name"`
	Avatar  string `json:"avatar"`
	Cover   string `json:"cover"`
	Feature string `json:"feature"`
	Style   int64  `json:"style"`
	Unit    string `json:"unit"`
	Spec    int64  `json:"spec"` // 200
	// 零售价
	Price int64    `json:"price"`
	Supps []int64  `json:"supps"`
	Units []*Units `json:"units"`

	// 重量
	PkWeight int64 `json:"pk_weight"`
	PkAmount int64 `json:"pk_amount"`
	// 单位
	SpecWeight int64 `json:"spec_weight"`
	//保质期（天）
	KeepLife int64 `json:"keep_life"`
	// 保质期单位
	KeepLifeUnit string `json:"keep_life_unit"`
}

type ProductLiteModel struct {
	Id   int64  `json:"id"`
	Sn   string `json:"sn"`
	Name string `json:"name"`
}

type ProductModel struct {
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

func (s *Store) GetPublicProductBySn(ctx context.Context, req *PrdSnReq, reply *PrdSnReqReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("GetPublicProductBySn", "GetPublicProductBySn", "Supp")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 3 查询公司信息
	productModel, err := db.GetPublicProductBySn(link, req.Sn)
	if err != nil {
		log.ErrorExit("GetPublicProductBySn Query err", err)
		return lang.Error("msg_public_product_not_found", err.Error())
	}
	reply.Sn = productModel.Sn
	return nil
}
