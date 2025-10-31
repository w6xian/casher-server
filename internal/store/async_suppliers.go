package store

import (
	"casher-server/internal/lager"
	"context"
	"time"
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
type SupplierLite struct {
	AppId    string `json:"com_api"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	SortName string `json:"sort_name"`
	SnCode   string `json:"sn_code"`
	Avatar   string `json:"avatar"`
	Pinyin   string `json:"pinyin"`
	Image    string `json:"image"`

	AreaCode   string `json:"area_code"`
	AreaPath   string `json:"area_path"`
	AreaStreet string `json:"area_street"`
	Street     string `json:"street"`
	Address    string `json:"address"`

	LegalPersion  string `json:"legal_persion"`
	ContactName   string `json:"contact_name"`
	ContactPhone  string `json:"contact_phone"`
	ContactMobile string `json:"contact_mobile"`
	ContactEmail  string `json:"contact_email"`
	RegisterDate  string `json:"register_date"`
	ExpireDate    string `json:"expire_date"`
	CreateTime    string `json:"create_time"`
	AuthId        int64  `json:"auth_id"`
	AuthAppId     string `json:"auth_app_id"`
	AuthAppSec    string `json:"auth_app_sec"`

	Status int64 `json:"status"`
	Intime int64 `json:"intime"`
}

// 返回同步供应商信息
type AsyncSuppliersReply struct {
	Req
	// 同步供应商信息
	Suppliers []*SupplierLite `json:"suppliers"`
	// 可同步供应商数
	TotalNum int64 `json:"total_num"`
	LastTime int64 `json:"last_time"`
}

func (s *Store) AsyncSuppliers(ctx context.Context, req *AsyncRequest, reply *AsyncSuppliersReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("AsyncSuppliers", "AsyncSuppliers", "async")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 3 查询订单信息
	res, err := db.QuerySuppliers(link, req)
	if err != nil {
		log.ErrorExit("QuerySuppliers Query err", err)
		return lang.Error("msg_suppliers_not_found", err.Error())
	}

	reply.AppId = req.AppId
	reply.Suppliers = res.Suppliers
	reply.TotalNum = res.TotalNum
	reply.LastTime = time.Now().Unix()
	return nil
}
