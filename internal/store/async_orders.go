package store

import (
	"casher-server/internal/lager"
	"context"
)

// 同一个请求，同步商品信息
type AsyncRequest struct {
	Req
	LastTime int64 `json:"last_time"`
	CloudId  int64 `json:"cloud_id"`
	Limit    int64 `json:"limit"`
}

func (req *AsyncRequest) Validate() error {
	if req.OpenId == "" {
		return req.Tracker.Error("async_orders_validate", "请输入open_id")
	}
	return nil
}

// 返回同步商品信息
type AsyncProductReply struct {
	Req
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 返回商品扩展信息（品牌，分类，规格等）
type AsyncProductExtraReply struct {
	Req
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

/*
*

	`track_id` bigint(20) NOT NULL DEFAULT '0',
	`ticket` varchar(45) NOT NULL DEFAULT '',
	`proxy_id` bigint(20) NOT NULL DEFAULT '0',
	`com_id` bigint(20) NOT NULL DEFAULT '0',
	`shop_id` bigint(20) NOT NULL DEFAULT '0',
	`machine_id` varchar(45) NOT NULL DEFAULT '' COMMENT '设备编号',
	`date_time` int(11) NOT NULL DEFAULT '0' COMMENT '单据日期',
	`dr` bigint(20) NOT NULL DEFAULT '0' COMMENT '订单正算收多少',
	`cr` bigint(20) NOT NULL DEFAULT '0' COMMENT '收了多少钱主要是现金',
	`off` bigint(20) NOT NULL DEFAULT '0' COMMENT '打折',
	`off_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '折扣金额',
	`abatement` bigint(20) NOT NULL DEFAULT '0' COMMENT '减免',
	`debit` bigint(20) NOT NULL DEFAULT '0' COMMENT '抹零',
	`discount` bigint(20) NOT NULL DEFAULT '0' COMMENT '=off+abatement+debit',
	`change` bigint(20) NOT NULL DEFAULT '0' COMMENT '找零',
	`coupons` int(11) NOT NULL DEFAULT '0' COMMENT '优惠券抵扣',
	`points` int(11) NOT NULL DEFAULT '0' COMMENT '积分支付',
	`balance` bigint(20) NOT NULL DEFAULT '0' COMMENT '余额支付金额',
	`payed` bigint(20) NOT NULL DEFAULT '0' COMMENT '最终支付',
	`prd_num` int(11) NOT NULL DEFAULT '0' COMMENT '单品数量',
	`shop_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '这里每家的积分与余额不同',
	`user_name` varchar(45) NOT NULL DEFAULT '',
	`handler_id` bigint(20) NOT NULL DEFAULT '0',
	`handler_name` varchar(45) NOT NULL DEFAULT '',
	`mark` varchar(45) NOT NULL DEFAULT '',
	`prints` int(11) NOT NULL DEFAULT '0' COMMENT '打印次数',
	`pay_type` varchar(45) NOT NULL DEFAULT '' COMMENT '支付方式''zero''表示本单没有收钱',
	`pay_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1已支付0未支付',
	`rec_time` int(11) NOT NULL DEFAULT '0' COMMENT '处理记录时间',
	`status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1正常0无效',
	`intime` int(11) NOT NULL DEFAULT '0' COMMENT '入库时间',
*/
type OrderLite struct {
	Id int64 `json:"id"`
	// 订单ID
	TrackId int64 `json:"track_id"`
	// 订单号
	Ticket string `json:"ticket"`
	// 代理ID
	ProxyId int64 `json:"proxy_id"`
	// 客户ID
	ComId int64 `json:"com_id"`
	// 店铺ID
	ShopId int64 `json:"shop_id"`
	// 设备编号
	MachineId string `json:"machine_id"`
	// 订单日期
	DateTime int `json:"date_time"`
	// 订单金额（应收）
	Dr int64 `json:"dr"`
	// 收了多少钱（应付）
	Cr int64 `json:"cr"`
	// 折扣金额
	Off int64 `json:"off"`
	// 折扣金额
	OffPrice int64 `json:"off_price"`
	// 减免金额
	Abatement int64 `json:"abatement"`
	// 抹零金额
	Debit int64 `json:"debit"`
	// 折扣金额
	Discount int64 `json:"discount"`
	// 找零金额
	Change int64 `json:"change"`
	// 优惠券抵扣
	Coupons int `json:"coupons"`
	// 积分支付
	Points int `json:"points"`
	// 余额支付
	Balance int64 `json:"balance"`
	// 最终支付金额
	Payed int64 `json:"payed"`
	// 单品数量
	PrdNum int `json:"prd_num"`
	// 店铺用户ID
	ShopUserId int64 `json:"shop_user_id"`
	// 店铺用户名
	UserName string `json:"user_name"`
	Mark     string `json:"mark"`
	// 打印次数
	Prints int `json:"prints"`
	// 支付方式
	PayType string `json:"pay_type"`
	// 支付状态
	PayStatus int `json:"pay_status"`
	// 处理记录时间
	RecTime int `json:"rec_time"`
	// 状态
	Status int `json:"status"`
	// 入库时间
	Intime int `json:"intime"`
}

// 返回同步订单信息
type AsyncOrdersReply struct {
	Req
	// 同步订单信息
	Orders []*OrderLite `json:"orders"`
	// 可同步订单数
	TotalNum int64 `json:"total_num"`
}

func (s *Store) AsyncOrders(ctx context.Context, req *AsyncRequest, reply *AsyncOrdersReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("AsyncOrders", "AsyncOrders", "async")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 3 查询订单信息
	res, err := db.QueryOrders(link, req)
	if err != nil {
		log.ErrorExit("QueryOrders Query err", err)
		return lang.Error("msg_orders_not_found", err.Error())
	}

	reply.AppId = req.AppId
	reply.Orders = res.Orders
	reply.TotalNum = res.TotalNum
	return nil
}
