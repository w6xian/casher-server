package store

import (
	"casher-server/internal/lager"
	"casher-server/internal/timex"
	"casher-server/pkg/checker"
	"casher-server/pkg/mapv"
	"context"
	"fmt"
	"slices"
)

/*
*

	`id` bigint(20) NOT NULL AUTO_INCREMENT,=
	`proxy_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '公司ID',
	`shop_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '门店ID',
	`prd_id` bigint(20) NOT NULL COMMENT '产品id',
	`sn` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '产品编号',
	`avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '封面',
	`cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '背面',
	`title` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '产品名称',
	`feature` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '产品卖点',
	`price` bigint(20) NOT NULL DEFAULT '0' COMMENT '零售价',
	`line_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '划线价',
	`sale_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '销售数量',
	`times` int(11) NOT NULL DEFAULT '1' COMMENT '售卖单位',
	`weight` int(11) NOT NULL DEFAULT '1' COMMENT '单个重量',
	`pack_name` varchar(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '箱' COMMENT '销售单位',
	`shop_stock` bigint(20) NOT NULL DEFAULT '0' COMMENT '此仓库库存（单位按最小规格，程序自动换算）',
	`tags` varchar(45) NOT NULL DEFAULT '[]' COMMENT '标签，最多三个',
	`sort` int(11) NOT NULL DEFAULT '50' COMMENT '排序',
	`status` int(11) NOT NULL DEFAULT '0' COMMENT '上架状态:0-下架;1-上架',
	`intime` int(11) NOT NULL DEFAULT '0' COMMENT '入库时间',
*/
// type ProductLite struct {
// 	Id        int64  `json:"id"`
// 	ProxyId   int64  `json:"proxy_id"`
// 	ShopId    int64  `json:"shop_id"`
// 	PrdId     int64  `json:"prd_id"`
// 	Sn        string `json:"sn"`
// 	Avatar    string `json:"avatar"`
// 	Cover     string `json:"cover"`
// 	Title     string `json:"title"`
// 	Feature   string `json:"feature"`
// 	Price     int64  `json:"price"`
// 	LinePrice int64  `json:"line_price"`
// 	SaleNum   int64  `json:"sale_num"`
// 	Times     int32  `json:"times"`
// 	Weight    int32  `json:"weight"`
// 	PackName  string `json:"pack_name"`
// 	ShopStock int64  `json:"shop_stock"`
// 	Tags      string `json:"tags"`
// 	Sort      int32  `json:"sort"`
// 	Status    int32  `json:"status"`
// 	Intime    int32  `json:"intime"`
// }

type ProductLite struct {
	Id        int64  `json:"id"`
	ProxyId   int64  `json:"proxy_id"`
	ShopId    int64  `json:"shop_id"`
	PrdId     int64  `json:"prd_id"`
	Sn        string `json:"sn"`
	Avatar    string `json:"avatar"`
	Cover     string `json:"cover"`
	Title     string `json:"title"`
	Feature   string `json:"feature"`
	Price     int64  `json:"price"`
	LinePrice int64  `json:"line_price"`
	SaleNum   int64  `json:"sale_num"`
	Times     int32  `json:"times"`
	Weight    int32  `json:"weight"`
	PackName  string `json:"pack_name"`
	ShopStock int64  `json:"shop_stock"`
	Tags      string `json:"tags"`
	Sort      int32  `json:"sort"`
	// 4件套
	Status int32 `json:"status"`
	Uptime int32 `json:"uptime"`
	Intime int32 `json:"intime"`
	// 其他

	UnionId string `json:"union_id"`
	Name    string `json:"name"`
	Pinyin  string `json:"pinyin"`

	Style      int64 `json:"style"`
	StyleType  int64 `json:"style_type"`
	Num        int64 `json:"num,omitempty"`
	Total      int64 `json:"total,omitempty"`
	Source     int64 `json:"source"`
	Type       int64 `json:"type"`
	PkWeight   int64 `json:"pk_weight"`
	SpecNameId int64 `json:"spec_name_id"`
	SpecWeight int64 `json:"spec_weight"`
	// 200g*20袋 = 箱
	Spec     int64  `json:"spec"`
	Unit     string `json:"unit"`
	PkAmount int64  `json:"pk_amount"`
	SpecName string `json:"spec_name"`
	// 库存
	Stock int64 `json:"stock"`
	// 内部
	Mark string `json:"mark"`
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
	CategoryId   int64  `json:"category_id,omitempty"`
	CategoryName string `json:"category_name"`
	// shopPrd与prd不同
	MajorPackName string `json:"major_pack_name"`
	MajorSpecName string `json:"major_spec_name"`
	MajorPkAmount int64  `json:"major_pk_amount"`
}

// IdRequestProductReply 商品更新信息请求参数
type IdRequestProductReply struct {
	Req
	Product *ProductLite `json:"product"`
}

// 返回同步商品信息
type AsyncProductsReply struct {
	Req
	// 同步商品信息
	Products []*ProductLite `json:"products"`
	// 可同步商品数
	TotalNum int64 `json:"total_num"`
	LastTime int64 `json:"last_time"`
}

func (s *Store) AsyncProducts(ctx context.Context, req *AsyncRequest, reply *AsyncProductsReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("AsyncProducts", "AsyncProducts", "async")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 3 查询订单信息
	res, err := db.QueryProducts(link, req)
	if err != nil {
		log.ErrorExit("QueryProducts Query err", err)
		return lang.Error("msg_products_not_found", err.Error())
	}

	reply.AppId = req.AppId
	reply.Products = res.Products
	reply.TotalNum = res.TotalNum
	reply.LastTime = timex.UnixTime()
	return nil
}

/*
CREATE TABLE `mi_com_products_categories` (
  `ctg_id` bigint(20) NOT NULL AUTO_INCREMENT,
  `proxy_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '对应的公司ID',
  `parent_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '了',
  `image` varchar(255) NOT NULL DEFAULT '' COMMENT '分类图片',
  `name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `mark` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '描述',
  `sort` int(11) NOT NULL DEFAULT '50' COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1有效0无效',
  `private` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1表示私有',
  `mall` smallint(6) NOT NULL DEFAULT '0' COMMENT '是否商城 0不是 1是',
  `materiel` tinyint(1) NOT NULL COMMENT '是否物料',
  `intime` int(11) NOT NULL DEFAULT '0' COMMENT '入库时间',
  `home` tinyint(1) NOT NULL DEFAULT '0' COMMENT '首页显示1与否0',
  PRIMARY KEY (`ctg_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3113 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='产品分类';

*/

type CategoryLite struct {
	Id       int64  `json:"ctg_id"`
	ProxyId  int64  `json:"proxy_id"`
	ParentId int64  `json:"parent_id"`
	Image    string `json:"image"`
	Name     string `json:"name"`
	Sort     int32  `json:"sort"`
	Mark     string `json:"mark"`
	Status   int32  `json:"status"`
	Intime   int32  `json:"intime"`
}

/*
*
CREATE TABLE `mi_crm_brands` (

	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`proxy_id` int(11) NOT NULL DEFAULT '0' COMMENT '对应的公司ID',
	`sn` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '编号',
	`name` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名字',
	`proprietor` varchar(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '甩有者',
	`mark` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
	`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态： 0正常 1禁用',
	`intime` int(11) NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_brands` (`proxy_id`) USING BTREE

) ENGINE=InnoDB AUTO_INCREMENT=2917 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='品牌管理';
*/
type BrandLite struct {
	Id      int64 `json:"id"`
	ProxyId int64 `json:"proxy_id"`

	Sn         string `json:"sn"`
	Name       string `json:"name"`
	Proprietor string `json:"proprietor"`

	Sort   int32  `json:"sort"`
	Mark   string `json:"mark"`
	Status int32  `json:"status"`
	Intime int32  `json:"intime"`
}

// 返回同步商品信息
type AsyncProductsExtraReply struct {
	Req
	// 同步商品信息
	Categories []*CategoryLite `json:"categories"`
	// 品牌
	Brands   []*BrandLite `json:"brands"`
	LastTime int64        `json:"last_time"`
}

func (s *Store) AsyncProductsExtra(ctx context.Context, req *AsyncRequest, reply *AsyncProductsExtraReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("AsyncProductsExtra", "AsyncProductsExtra", "async")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 3 查询订单信息
	res, err := db.QueryProductsExtra(link, req)
	if err != nil {
		log.ErrorExit("QueryProductsExtra Query err", err)
		return lang.Error("msg_products_not_found", err.Error())
	}

	reply.AppId = req.AppId
	reply.Categories = res.Categories
	reply.Brands = res.Brands
	reply.LastTime = timex.UnixTime()
	return nil
}

// 同步单一商品更新信息
type AsyncProductUpdateReply struct {
	Req
	// 同步商品信息
	Product  *ProductLite `json:"product"`
	LastTime int64        `json:"last_time"`
}

// AsyncProductLite 同步单一商品更新信息
func (s *Store) AsyncProductLite(ctx context.Context, req *IdRequest, reply *IdRequestProductReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("AsyncProductLite", "AsyncProductLite", "async")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 3 查询订单信息
	res, err := db.QueryProductUpdate(link, req)
	if err != nil {
		log.ErrorExit("QueryProductUpdate Query err", err)
		return lang.Error("msg_products_not_found", err.Error())
	}

	reply.AppId = req.AppId
	reply.Product = res
	return nil
}

// AsyncUpdateProduct 主动更新商品信息（如库存，价格，状态等）
func (s *Store) AsyncUpdateProduct(ctx context.Context, req *UpdateRequest, reply *UpdateReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("AsyncUpdateProduct", "AsyncUpdateProduct", "async")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 商品更新信息限定
	if req.Values == nil {
		return lang.Error("msg_products_not_found", "values is nil")
	}
	cols := []string{"stock", "price", "cost", "unit_price", "pack_price", "keep_life", "keep_life_unit", "status"}
	for k := range req.Values {
		if !slices.Contains(cols, k) {
			return lang.Error("msg_products_not_found", fmt.Sprintf("key %s not in %v", k, cols))
		}
	}

	checks := checker.New(checker.Int64("stock"),
		checker.Int64("price"),
		checker.Int64("cost"),
		checker.Int64("unit_price"),
		checker.Int64("pack_price"),
		checker.Int64("style"),
		checker.Int64("keep_life"),
		checker.String("keep_life_unit"),
		checker.Int32("status"),
	)
	kv, err := checks.CheckMap(req.Values)
	if err != nil {
		return err
	}
	// 4 查询商品信息
	product, err := db.GetProductByUnionId(link, req.Tracker.ProxyId, req.Tracker.ShopId, req.UnionId)
	if err != nil {
		log.ErrorExit("GetProductByUnionId Query err", err)
		return lang.Error("msg_products_not_found", err.Error())
	}
	if product == nil {
		return lang.Error("msg_products_not_found", fmt.Sprintf("union_id %s not found", req.UnionId))
	}

	//只更新有差异的字段
	if product.Uptime > req.Uptime {
		return lang.Error("msg_products_uptime_invalid", fmt.Sprintf("uptime %d is invalid", req.Uptime))
	}
	upValues := map[string]any{}
	v := mapv.NewMapv(kv)
	if product.Stock != v.Int64("stock") {
		upValues["stock"] = v.Int64("stock")
	}
	// "stock", "price", "cost", "unit_price", "pack_price", "keep_life", "keep_life_unit", "status"
	if product.Price != v.Int64("price") {
		upValues["price"] = v.Int64("price")
	}
	if product.Cost != v.Int64("cost") {
		upValues["cost"] = v.Int64("cost")
	}
	if product.UnitPrice != v.Int64("unit_price") {
		upValues["unit_price"] = v.Int64("unit_price")
	}
	if product.PackPrice != v.Int64("pack_price") {
		upValues["pack_price"] = v.Int64("pack_price")
	}
	if product.KeepLife != v.Int64("keep_life") {
		upValues["keep_life"] = v.Int64("keep_life")
	}
	if product.KeepLifeUnit != v.String("keep_life_unit") {
		upValues["keep_life_unit"] = v.String("keep_life_unit")
	}
	if product.Status != v.Int32("status") {
		upValues["status"] = v.Int32("status")
	}

	// 3 查询订单信息
	status, err := db.AsyncUpdateProduct(link, req, upValues)
	if err != nil {
		log.ErrorExit("AsyncUpdateProduct Query err", err)
		return lang.Error("msg_products_not_found", err.Error())
	}

	reply.AppId = req.AppId
	reply.Status = status
	return nil
}
