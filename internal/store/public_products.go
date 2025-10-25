package store

import (
	"casher-server/internal/errors"
	"casher-server/internal/lager"
	"context"
)

type PrdSnReq struct {
	Req
	Sn string `json:"sn"`
}

func (req *PrdSnReq) Validate() error {
	if req.Sn == "" {
		return errors.New("sn required")
	}
	return nil
}

type PrdNameReq struct {
	Req
	Name string `json:"name"`
}

func (req *PrdNameReq) Validate() error {
	if req.Name == "" {
		return errors.New("name required")
	}
	return nil
}

type PublicProductReply struct {
	Req
	UnionId                  string  `json:"union_id"`
	Avatar                   string  `json:"avatar"`
	Cover                    string  `json:"cover"`
	Sn                       string  `json:"sn"`
	Name                     string  `json:"name"`
	Pinyin                   string  `json:"pinyin"`
	BrandName                string  `json:"brand_name"`
	Feature                  string  `json:"feature"`
	Price                    int64   `json:"price"`
	Spec                     int64   `json:"spec"`
	SpecName                 string  `json:"spec_name"`
	SpecWeight               int64   `json:"spec_weight"`
	PkAmount                 int64   `json:"pk_amount"`
	PkWeight                 int64   `json:"pk_weight"`
	PackName                 string  `json:"pack_name"`
	KeepLife                 int64   `json:"keep_life"`
	KeepLifeUnit             string  `json:"keep_life_unit"`
	StorageConditions        string  `json:"storage_conditions"`
	TransportationConditions string  `json:"transportation_conditions"`
	Unit                     string  `json:"unit"`
	Habitat                  string  `json:"habitat"`
	Style                    int64   `json:"style"`
	StyleType                int64   `json:"style_type"`
	Units                    []Units `json:"units"`
	Status                   int64   `json:"status"`
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

/**
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '产品id',
  `union_id` varchar(45) NOT NULL COMMENT '隐藏直实Id',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '封面',
  `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '背面',
  `sn` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '产品编号',
  `name` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '产品名称',
  `pinyin` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '首拼字母(SPZM)',
  `brand_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '商品品牌',
  `feature` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '产品卖点',
  `price` bigint(20) NOT NULL DEFAULT '0' COMMENT '零售价格、不低于a_price',
  `spec` int(11) NOT NULL DEFAULT '0' COMMENT '规格',
  `spec_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '瓶' COMMENT '最小规格名称，袋、瓶',
  `spec_weight` int(11) NOT NULL DEFAULT '0' COMMENT '最小规格重量（含包装）',
  `pk_amount` int(11) NOT NULL DEFAULT '1' COMMENT '装箱数量',
  `pk_weight` int(11) NOT NULL DEFAULT '1' COMMENT '装箱后总重量（包含包装）',
  `pack_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '箱' COMMENT '打包后的名称',
  `keep_life` mediumint(9) NOT NULL DEFAULT '0' COMMENT '开启保质期(需要开起批次号)',
  `keep_life_unit` varchar(64) NOT NULL DEFAULT '天' COMMENT '1*24表天，7*24表周，30*24表月，365*24表年，默认天',
  `storage_conditions` varchar(64) NOT NULL DEFAULT '' COMMENT '保存方法',
  `transportation_conditions` varchar(64) NOT NULL DEFAULT '' COMMENT '运输条件',
  `unit` varchar(2) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'g' COMMENT '计量单位名默认g(克)还有ml(毫升)等',
  `habitat` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '产地（主要用于源材料）',
  `style` tinyint(4) NOT NULL DEFAULT '0' COMMENT '包装方式1预包装，2散装',
  `style_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1 称重2 量体3 点数\r\n',
  `units` varchar(512) NOT NULL DEFAULT '' COMMENT '扩展规格',
  `status` int(11) NOT NULL DEFAULT '0' COMMENT '1有效',
  `intime` int(11) NOT NULL DEFAULT '0' COMMENT '入库时间',
*/

type ProductModel struct {
	UnionId                  string  `json:"union_id"`
	Avatar                   string  `json:"avatar"`
	Cover                    string  `json:"cover"`
	Sn                       string  `json:"sn"`
	Name                     string  `json:"name"`
	Pinyin                   string  `json:"pinyin"`
	BrandName                string  `json:"brand_name"`
	Feature                  string  `json:"feature"`
	Price                    int64   `json:"price"`
	Spec                     int64   `json:"spec"`
	SpecName                 string  `json:"spec_name"`
	SpecWeight               int64   `json:"spec_weight"`
	PkAmount                 int64   `json:"pk_amount"`
	PkWeight                 int64   `json:"pk_weight"`
	PackName                 string  `json:"pack_name"`
	KeepLife                 int64   `json:"keep_life"`
	KeepLifeUnit             string  `json:"keep_life_unit"`
	StorageConditions        string  `json:"storage_conditions"`
	TransportationConditions string  `json:"transportation_conditions"`
	Unit                     string  `json:"unit"`
	Habitat                  string  `json:"habitat"`
	Style                    int64   `json:"style"`
	StyleType                int64   `json:"style_type"`
	Units                    []Units `json:"units"`
	Status                   int64   `json:"status"`
}

func (s *Store) GetPublicProductBySn(ctx context.Context, req *PrdSnReq, reply *PublicProductReply) error {
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
	reply.AppId = req.AppId
	reply.Sn = productModel.Sn
	reply.Name = productModel.Name
	reply.UnionId = productModel.UnionId
	reply.Avatar = productModel.Avatar
	reply.Cover = productModel.Cover
	reply.Pinyin = productModel.Pinyin
	reply.BrandName = productModel.BrandName
	reply.Feature = productModel.Feature
	reply.Price = productModel.Price
	reply.Spec = productModel.Spec
	reply.SpecName = productModel.SpecName
	reply.SpecWeight = productModel.SpecWeight
	reply.PkAmount = productModel.PkAmount
	reply.PkWeight = productModel.PkWeight
	reply.PackName = productModel.PackName
	reply.KeepLife = productModel.KeepLife
	reply.KeepLifeUnit = productModel.KeepLifeUnit
	reply.StorageConditions = productModel.StorageConditions
	reply.TransportationConditions = productModel.TransportationConditions
	reply.Unit = productModel.Unit
	reply.Habitat = productModel.Habitat
	reply.Style = productModel.Style
	reply.StyleType = productModel.StyleType
	reply.Units = productModel.Units
	reply.Status = productModel.Status
	return nil
}
