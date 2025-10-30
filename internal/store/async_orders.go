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

	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`track_id` bigint(20) NOT NULL COMMENT '最始的ID，用于采购单转换',
	`proxy_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '对应的公司ID',
	`parent_track` bigint(20) NOT NULL DEFAULT '0' COMMENT '父订单trackId',
	`split_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已经拆单 0 没有 1拆过',
	`share_confirm_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '分享销售单确认人ID',
	`share_confirm_user_name` varchar(45) NOT NULL DEFAULT '' COMMENT '分享销售单确认人名称',
	`ticket` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '销售单编号',
	`po_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '对应的采购单ID',
	`date_time` int(11) NOT NULL DEFAULT '0',

	`shop_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '商店（商城的订单）',
	`user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '对应的销售经理ID，有可能是用户订单',
	`user_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '销售经理名称',
	`handler_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '经手人ID',
	`handler_name` varchar(45) NOT NULL DEFAULT '' COMMENT '经手人名称',

	`province` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '省',
	`province_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '省全省',
	`city` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '市',
	`city_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '市全称',
	`district` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '区',
	`area_code` varchar(64) NOT NULL DEFAULT '' COMMENT '2位或者12位跟统计用区域',
	`area_path` varchar(1024) NOT NULL DEFAULT '' COMMENT '带/的地址',
	`area_street` varchar(64) NOT NULL DEFAULT '' COMMENT '用户自己填写的',
	`district_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '区全称',
	`street` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '街道',
	`address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '省市区街道',
	`zipcode` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '邮编',
	`com_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '客户商品，后台指定，可以视为零售',
	`com_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '下单公司名称',
	`consignee` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '收货人',
	`mobile` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '手机号',

	`price` bigint(20) NOT NULL DEFAULT '0' COMMENT '单据金额',
	`discount` bigint(20) NOT NULL DEFAULT '0' COMMENT '优惠金额',
	`debt` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '非正常优惠（如抹零）',
	`payed` bigint(20) NOT NULL DEFAULT '0' COMMENT '实收金额',
	`tax_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '税金',
	`tax_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '含税价',

	`order_weight` bigint(20) NOT NULL DEFAULT '0' COMMENT '采购单重量',
	`pack_amount` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '总箱数',
	`box_amount` int(11) NOT NULL DEFAULT '0' COMMENT '各商品不成箱的单品数总和',
	`file` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '合同文件',
	`brief` varchar(128) NOT NULL DEFAULT '' COMMENT '摘要(用户填写)',

	`audit_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '审核状态0:未审核1:也提交,2主管审核，3进行中，4完成，-1终止',
	`audit_steps` tinyint(4) NOT NULL DEFAULT '0',
	`audit_pos` tinyint(4) NOT NULL DEFAULT '0',
	`audit_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '审核人id',
	`audit_user_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '审核人名称',
	`audit_time` int(11) NOT NULL DEFAULT '0' COMMENT '审核时间',

	`op_times` int(11) NOT NULL DEFAULT '0' COMMENT '安排生产/采购次数',
	`confirm_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '对方确认状态0:未确认1:已确认',
	`confirm_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '确认人id',
	`confirm_user_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '确认人名称',
	`confirm_time` int(11) NOT NULL DEFAULT '0' COMMENT '确认时间',

	`bill_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '账单类型 0:未结算|1:已结算',
	`bill_time` int(11) NOT NULL DEFAULT '0' COMMENT '开票时间',
	`prod_status` tinyint(4) NOT NULL DEFAULT '-1' COMMENT '生产装态 -1条件不足，0，可安排，1，进入计划，2生产任务，3生产确认，4生产中，5等待入库，6完成',
	`prod_time` int(11) NOT NULL DEFAULT '0' COMMENT '生产状态变化',
	`pay_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '支付状态：0未支付，1已支付',
	`pay_platform` varchar(45) NOT NULL DEFAULT '' COMMENT '支付平台',
	`pay_trans_id` varchar(64) NOT NULL DEFAULT '' COMMENT '支付交易ID',
	`pay_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '支付时间',

	`shipping_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '运输状态0,未发运，1，部分发运，2完成发运',
	`shipping_time` int(11) NOT NULL DEFAULT '0' COMMENT '发货时间',
	`shipping_enabled` tinyint(4) NOT NULL DEFAULT '0' COMMENT ' 能否发货，1表示可以发，0表示条件不允许 99 表示订单正在特批审核中',

	`transport_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否已上车，0，没有，1上车',
	`transport_time` int(11) NOT NULL DEFAULT '0' COMMENT '出库发运时间',
	`update_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后更新时间',

	`receive_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '收货人ID',
	`receive_user_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '收货人名称',
	`receive_time` int(11) NOT NULL DEFAULT '0' COMMENT '签收时间',
	`receive_status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '-1拒签，1部分签收，2完整签收',

	`readed` int(11) NOT NULL DEFAULT '0' COMMENT '是否阅读',
	`readtime` int(11) NOT NULL DEFAULT '0' COMMENT '阅读时间',
	`printed` int(11) NOT NULL DEFAULT '0' COMMENT '打印次数',
	`sort` int(11) NOT NULL DEFAULT '0' COMMENT '序号',
	`plan_lock` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 未加入发运计划 1 已加入发运计划',
	`plan_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 初建 1 部分发运 2 全部发运',
	`plan_time` int(11) NOT NULL DEFAULT '0' COMMENT '加入发运计划时间',
	`plan_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '所在发运计划ID',
	`plan_name` varchar(45) NOT NULL DEFAULT '' COMMENT '计划名称',
	`itinerary_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '配货路线',
	`itinerary_name` varchar(45) NOT NULL DEFAULT '',
	`produce_lock` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0 未加入生产计划 1 已加入生产计划',
	`source` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1表示销售下单，2经销商下单,3,店铺下单，4小程序下单',
	`store_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '收货仓库ID',
	`store_manager_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '收货仓库管理员员工ID',
	`reject_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '撤回',
	`reject_time` int(11) NOT NULL DEFAULT '0',
	`reject_user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '撤销人',
	`reject_user_name` varchar(45) NOT NULL DEFAULT '',
	`type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0新建,1部分完成，2完成',
	`po_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '对应的采购单状态',
	`so_status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '销售订单状态',
	`deposit` bigint(20) NOT NULL DEFAULT '0' COMMENT '订金相关',
	`deposit_status` tinyint(1) NOT NULL DEFAULT '0',
	`deposit_time` int(11) NOT NULL DEFAULT '0' COMMENT '收订金时间',
	`deposit_handler_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '订金经手人',
	`deposit_handler_name` varchar(45) NOT NULL DEFAULT '' COMMENT '订金经手人',
	`logistics_com` varchar(45) NOT NULL DEFAULT '' COMMENT '物流公司',
	`logistics_no` varchar(45) NOT NULL DEFAULT '' COMMENT '物流单号',
	`logistics_fee` int(11) NOT NULL DEFAULT '0' COMMENT '物流费用',
	`commented` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1已评价',
	`mark` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注信息',
	`status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '订单状态0无效，1有效',
	`poster_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '制单人',
	`poster_name` varchar(45) NOT NULL DEFAULT '' COMMENT '制单人',
	`vouched` tinyint(1) NOT NULL DEFAULT '0' COMMENT '有没有上传过凭证',
	`intime` int(10) unsigned NOT NULL DEFAULT '0',
*/
type OrderLite struct {
	Id int64 `json:"id"`
	// 订单来源
	TrackId     int64  `json:"track_id"`
	TrackStr    string `json:"track_str"`
	ProxyId     int64  `json:"proxy_id"`
	ParentTrack int64  `json:"parent_track"`
	SplitStatus int64  `json:"split_status"`
	//share
	ShareConfirmUserId   int64  `json:"share_confirm_user_id"`
	ShareConfirmUserName string `json:"share_confirm_user_name"`
	Ticket               string `json:"ticket"`
	PoId                 int64  `json:"po_id"`
	DateTime             int64  `json:"date_time"`
	// 订单来源
	ShopId      int64  `json:"shop_id"`
	UserId      int64  `json:"user_id"`
	UserName    string `json:"user_name"`
	HandlerId   int64  `json:"handler_id"`
	HandlerName string `json:"handler_name"`

	// 订单渠道
	Source    int    `json:"source"`
	Channel   string `json:"channel"`
	MachineId string `json:"machine_id"`

	ComId   int64  `json:"com_id"`
	ComName string `json:"com_name"`

	// 订单商品数量
	PrdNum       int    `json:"prd_num"`
	ShopUserId   int64  `json:"shop_user_id"`
	ShopUserName string `json:"shop_user_name"`

	// 订单地址
	Zipcode    string `json:"zipcode"`
	AreaCode   string `json:"area_code"`
	AreaPath   string `json:"area_path"`
	AreaStreet string `json:"area_street"`
	Address    string `json:"address"`
	Consignee  string `json:"consignee"`
	Mobile     string `json:"mobile"`

	// 金额相关
	Price     int64 `json:"price"`
	Discount  int64 `json:"discount"`
	Debt      int64 `json:"debt"`
	Payed     int64 `json:"payed"`
	TaxFee    int64 `json:"tax_fee"`
	TaxAmount int64 `json:"tax_amount"`
	// 订单重量
	OrderWeight int64 `json:"order_weight"`
	// 数量统计
	PackAmount int64 `json:"pack_amount"`
	BoxAmount  int64 `json:"box_amount"`

	File  string `json:"file"`
	Brief string `json:"brief"`

	//审核 '审核状态0:未审核1:也提交,2主管审核，3进行中，4完成，-1终止',
	AuditType     int    `json:"audit_type"`
	AuditSteps    int    `json:"audit_steps"`
	AuditPos      int    `json:"audit_pos"`
	AuditStatus   int    `json:"audit_status"`
	AuditUserId   int64  `json:"audit_user_id"`
	AuditUserName string `json:"audit_user_name"`
	AuditTime     int64  `json:"audit_time"`

	ShippingEnabled bool   `json:"shipping_enabled"`
	ShippingStatus  int    `json:"shipping_status"`
	ShippingTime    int64  `json:"shipping_time"`
	ShippingMark    string `json:"shipping_mark"`

	//物流公司
	LogisticsCom string `json:"logistics_com"`
	LogisticsNo  string `json:"logistics_no"`
	LogisticsFee int64  `json:"logistics_fee"`
	/*运输相关*/
	// '是否已上车，0，没有，1上车',
	TransportStatus int   `json:"transport_status"`
	TransportTime   int64 `json:"transport_time"`

	// 收货人
	ReceiveUserId   int64  `json:"receive_user_id"`
	ReceiveUserName string `json:"receive_user_name"`
	ReceiveTime     int64  `json:"receive_time"`
	ReceiveStatus   string `json:"receive_status"`

	// 配货路线
	ItineraryId   int64  `json:"itinerary_id"`
	ItineraryName string `json:"itinerary_name"`

	// 仓库
	StoreId        int64  `json:"store_id"`
	StoreName      string `json:"store_name"`
	StoreManagerId int64  `json:"store_manager_id"`

	// 1撤回
	RejectStatus   int    `json:"reject_status"`
	RejectTime     int64  `json:"reject_time"`
	RejectUserId   int64  `json:"reject_user_id"`
	RejectUserName string `json:"reject_user_name"`

	// type
	Type int `json:"type"`

	// 订金相关
	Deposit            int64  `json:"deposit"`
	DepositStatus      int    `json:"deposit_status"`
	DepositTime        int64  `json:"deposit_time"`
	DepositHandlerId   int64  `json:"deposit_handler_id"`
	DepositHandlerName string `json:"deposit_handler_name"`

	// 支付相关
	BillStatus int   `json:"bill_status"`
	BillTime   int64 `json:"bill_time"`
	/*支付状态*/
	PayStatus int `json:"pay_status"`
	/*支付平台*/
	PayPlatform int `json:"pay_platform"`
	/*支付交易id*/
	PayTransId string `json:"pay_trans_id"`
	/*支付时间*/
	PayTime int64 `json:"pay_time"`

	// 评论状态
	Commented int `json:"commented"`
	// 备注信息
	Mark int `json:"mark"`

	// 状态
	Status int `json:"status"`
	// 打印
	Printed int `json:"printed"`
	/*更新时间*/
	UpdateTime int64 `json:"update_time"`
	// 确认时间
	Intime int64 `json:"intime"`
	// 订单商品信息
	Items []*OrderLiteItem `json:"items"`
}

/*
*

	    `id` bigint(20) NOT NULL AUTO_INCREMENT,
		`proxy_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '公司id',
		`shop_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '门店',
		`prd_id` bigint(20) NOT NULL DEFAULT '0',
		`store_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '仓库，0表示商品默认',

		`store_name` varchar(45) NOT NULL DEFAULT '',
		`so_id` bigint(20) NOT NULL COMMENT '销售单编号',
		`prd_sn` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '产品sn',
		`prd_avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
		`prd_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '产品名称',

		`spec_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '格规名',
		`spec` int(11) NOT NULL DEFAULT '0' COMMENT '规格',
		`spec_weight` int(11) NOT NULL DEFAULT '0',
		`unit` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'g' COMMENT '重量或体积单位',
		`style` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1预包装\n2散装',

		`pack_name` varchar(8) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '箱' COMMENT '打包方式',
		`style_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1称重\n2量体\n3点数',
		`pk_amount` int(11) NOT NULL DEFAULT '0' COMMENT '打包数量',
		`pk_weight` int(11) NOT NULL DEFAULT '0',
		`unit_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '价格/unit=pk_amount就是箱价',

		`pack_price` bigint(20) NOT NULL DEFAULT '0' COMMENT '箱价格',
		`times` int(11) NOT NULL DEFAULT '0' COMMENT '选择的单位（单个，还是箱）',
		`num_pay` bigint(20) NOT NULL DEFAULT '0' COMMENT '记数价数',
		`num_gift` bigint(20) NOT NULL DEFAULT '0' COMMENT '非记价数',
		`num_total` int(11) NOT NULL DEFAULT '0' COMMENT '总数',

		`price_total` bigint(20) NOT NULL DEFAULT '0' COMMENT '合计总价',
		`price_discount` bigint(20) NOT NULL DEFAULT '0' COMMENT '合计优惠',
		`price_payed` bigint(20) NOT NULL DEFAULT '0' COMMENT '合计支付',
		`off` int(11) NOT NULL DEFAULT '10000',
		`tax_unit_price` bigint(20) NOT NULL DEFAULT '0',

		`tax_pack_price` bigint(20) NOT NULL DEFAULT '0',
		`tax` int(11) NOT NULL DEFAULT '0' COMMENT '税率（1300=13%）',
		`tax_fee` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付税金',
		`tax_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '含税总价',
		`mark` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',

		`intime` int(11) NOT NULL DEFAULT '0' COMMENT '入库时间',
		`produce_plan_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '加入到哪个生产计划',
		`plan_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '这个被提到了哪个计划里',
		`plan_exp_num_pay` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '计划发货计价数',
		`plan_exp_num_total` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '计划发货计数数',
		`trans_num_pay` int(11) NOT NULL DEFAULT '0' COMMENT '计划发货计价数',
		`trans_num_total` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '计划发货计数数',
		`print_mark` varchar(200) NOT NULL DEFAULT '' COMMENT '打印备注',
		`print_batch_time` varchar(45) NOT NULL DEFAULT '' COMMENT '打印生产日期',
		`print_produce_time` varchar(45) NOT NULL DEFAULT '' COMMENT '打印保质日期',
*/
type OrderLiteItem struct {
	Id      int64 `json:"id"`
	ProxyId int64 `json:"proxy_id"`
	ShopId  int64 `json:"shop_id"`
	PrdID   int64 `json:"prd_id"`
	StoreId int64 `json:"store_id"`

	StoreName string `json:"store_name"`
	SoId      int64  `json:"so_id"`
	PrdSn     string `json:"prd_sn"`
	PrdAvatar string `json:"prd_avatar"`
	PrdName   string `json:"prd_name"`

	SpecName   string `json:"spec_name"`
	Spec       int64  `json:"spec"`
	SpecWeight int64  `json:"spec_weight"`
	Unit       string `json:"unit"`
	Style      int64  `json:"style"`

	PackName  string `json:"pack_name"`
	StyleType int64  `json:"style_type"`
	PkAmount  int64  `json:"pk_amount"`
	PkWeight  int64  `json:"pk_weight"`
	UnitPrice int64  `json:"unit_price"`

	PackPrice int64 `json:"pack_price"`
	Times     int64 `json:"times"`
	NumPay    int64 `json:"num_pay"`
	NumGift   int64 `json:"num_gift"`
	NumTotal  int64 `json:"num_total"`

	PriceTotal    int64 `json:"price_total"`
	PriceDiscount int64 `json:"price_discount"`
	PricePayed    int64 `json:"price_payed"`
	Off           int64 `json:"off"`
	TaxUnitPrice  int64 `json:"tax_unit_price"`

	TaxPackPrice int64  `json:"tax_pack_price"`
	Tax          int64  `json:"tax"`
	TaxFee       int64  `json:"tax_fee"`
	TaxAmount    int64  `json:"tax_amount"`
	Mark         string `json:"mark"`

	Intime int64 `json:"intime"`
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
	// 4 订单数据处理
	for _, order := range res.Orders {
		id := order.Id
		items, err := db.SelectOrderItems(link, id)
		if err != nil {
			log.ErrorExit("SelectOrderItems Query err", err)
			items = []*OrderLiteItem{}
		}
		order.Items = items
	}

	reply.AppId = req.AppId
	reply.Orders = res.Orders
	reply.TotalNum = res.TotalNum
	return nil
}
