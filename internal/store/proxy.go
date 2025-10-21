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

/**
`com_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '行号Id',
 `open_id` bigint(20) NOT NULL COMMENT '对外公司的ID号',
 `name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公司',
 `sort_name` varchar(20) NOT NULL DEFAULT '' COMMENT '公司短昵称',
 `sn_code` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '公司代码',
 `legal_persion` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '法人',
 `contact_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人',
 `contact_phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系方式',
 `contact_email` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '联系人邮箱',
 `register_date` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '注册时间',
 `expire_date` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '失效时间',
 `categories_id` int(11) NOT NULL DEFAULT '0' COMMENT '分类',
 `categories_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '分类名',
 `level_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0-100',
 `level_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '等级名称',
 `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '头像图片路径',
 `tutor_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '客服经理',
 `tutor_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '客服经理',
 `manager_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '销售经理',
 `manager_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '销售经理',
 `create_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建者',
 `create_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '创建者',
 `in_groups` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '组列表,隔开',
 `bill_status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否启帐：1已开启，0关闭',
 `init_account` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否初始化资金账户 1 已初始化 0 未初始化',
 `init_subject` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是够初始化会计科目 1 已初始化 0 未初始化',
 `init_time` int(11) NOT NULL DEFAULT '0' COMMENT '启账时间',
 `province` mediumint(9) NOT NULL DEFAULT '0' COMMENT '省编码',
 `city` mediumint(9) NOT NULL DEFAULT '0' COMMENT '市编码',
 `district` mediumint(9) NOT NULL DEFAULT '0' COMMENT '区编码',
 `area_code` varchar(64) NOT NULL DEFAULT '' COMMENT '2位或者12位跟统计用区域',
 `area_path` varchar(1024) NOT NULL DEFAULT '' COMMENT '带/的地址',
 `area_street` varchar(64) NOT NULL DEFAULT '' COMMENT '用户自己填写的',
 `street` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '详细地址',
 `address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '全地址',
 `site` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '网站',
 `mark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
 `supplier_id` int(11) NOT NULL DEFAULT '0' COMMENT '系统',
 `x_code` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '神秘代码，用于对接第三方',
 `use_protection` tinyint(1) NOT NULL DEFAULT '0' COMMENT '区域保护1、代理商只允许买已代理的商品0，可以卖全部',
 `use_audit` tinyint(1) NOT NULL DEFAULT '0' COMMENT '采购单审核1、需要审核组里的人员审核采购单',
 `use_stock_limit` tinyint(1) NOT NULL DEFAULT '1' COMMENT '开启库存保护，1、不永许超卖，0，无限制',
 `use_marketing` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否开启销售营销功能 与com_config里面use_marketing联动! 请同时改变两张表的值',
 `invoice_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开票单位名',
 `invoice_addr` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开票地址',
 `invoice_bank_name` varchar(55) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开票开户行',
 `invoice_bank_id` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开票账号',
 `invoice_bank_no` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开票行号',
 `invoice_tax_id` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开票税号',
 `invoice_phone` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '开票电话',
 `scale` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '规模',
 `theme` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '#B10DC9' COMMENT '皮肤',
 `role_id` int(11) NOT NULL DEFAULT '0',
 `role_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户权限',
 `code` bigint(20) NOT NULL DEFAULT '0' COMMENT '运输公司专用code 96543+Request::Id',
 `prev_sn` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '自产产品，需要国家物码中心代码即：prev_sn',
 `inventory_status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否在盘点 0 没在 1在',
 `locked` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否锁定系统',
 `server_start` int(11) NOT NULL DEFAULT '0' COMMENT '服务开始时间',
 `server_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后服务时间（试用需要后台自己处理）',
 `intime` int(11) NOT NULL DEFAULT '0' COMMENT '入库时间',
 `pay_times` int(11) NOT NULL DEFAULT '0' COMMENT '成功支付次数',
 `pay_amount` bigint(20) NOT NULL DEFAULT '0' COMMENT '支付总金额',
 `pay_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后一次支付时间',
 `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '1合作，0开发，-1退出',
 `shop_open_id_init` int(11) NOT NULL DEFAULT '1000000' COMMENT '门店初始编号',
 `channel_id` varchar(45) NOT NULL DEFAULT '' COMMENT '渠道',
 `channel_name` varchar(64) NOT NULL DEFAULT '',
*/

// TODO 把COMMENT 以备注的形式迁移到结构体对应的字段上
type CloudCompany struct {
	ComId           int64  `json:"com_id"`            // 公司id
	UnionId         string `json:"union_id"`          // 集团公司
	OpenId          int64  `json:"open_id"`           // 对外公司的ID号
	Name            string `json:"name"`              // 公司名称
	SortName        string `json:"sort_name"`         // 公司短昵称
	SnCode          string `json:"sn_code"`           // 公司代码
	LegalPersion    string `json:"legal_persion"`     // 法人
	ContactName     string `json:"contact_name"`      // 联系人
	ContactPhone    string `json:"contact_phone"`     // 联系人电话
	ContactMobile   string `json:"contact_mobile"`    // 联系人手机号
	ContactEmail    string `json:"contact_email"`     // 联系人邮箱
	RegisterDate    string `json:"register_date"`     // 注册时间
	ExpireDate      string `json:"expire_date"`       // 失效时间
	CategoryId      int64  `json:"categories_id"`     // 分类
	CategoryName    string `json:"categories_name"`   // 分类名
	LevelId         int64  `json:"level_id"`          // 0-100
	LevelName       string `json:"level_name"`        // 等级名称
	Avatar          string `json:"avatar"`            // 头像图片路径
	TutorId         int64  `json:"tutor_id"`          // 客服经理
	TutorName       string `json:"tutor_name"`        // 客服经理
	ManagerId       int64  `json:"manager_id"`        // 销售经理
	ManagerName     string `json:"manager_name"`      // 销售经理名称
	CreateId        int64  `json:"create_id"`         // 创建者
	CreateName      string `json:"create_name"`       // 创建者
	InGroups        string `json:"in_groups"`         // 组列表,隔开
	BillStatus      int64  `json:"bill_status"`       // 是否启帐：1已开启，0关闭
	InitAccount     int64  `json:"init_account"`      // 是否初始化资金账户 1 已初始化 0 未初始化
	InitSubject     int64  `json:"init_subject"`      // 是够初始化会计科目 1 已初始化 0 未初始化
	InitTime        int64  `json:"init_time"`         // 启账时间
	AreaCode        string `json:"area_code"`         // 2位或者12位跟统计用区域
	AreaPath        string `json:"area_path"`         // 带/的地址
	AreaStreet      string `json:"area_street"`       // 用户自己填写的街道
	Street          string `json:"street"`            // 街道
	Address         string `json:"address"`           // 详细地址
	Site            string `json:"site"`              // 网站
	Mark            string `json:"mark"`              // 备注
	SupplierId      int64  `json:"supplier_id"`       // 系统
	MchCode         string `json:"mch_code"`          // 商户ID编码shortUUID
	XCode           string `json:"x_code"`            // 神秘代码，用于对接第三方
	UseProtection   bool   `json:"use_protection"`    // 区域保护1、代理商只允许买已代理的商品0，可以卖全部
	UseAudit        bool   `json:"use_audit"`         // 采购单审核1、需要审核组里的人员审核采购单
	UseStockLimit   bool   `json:"use_stock_limit"`   // 开启库存保护，1、不永许超卖，0，无限制
	UseMarketing    bool   `json:"use_marketing"`     // 是否开启销售营销功能 与com_config里面use_marketing联动! 请同时改变两张表的值
	InvoiceName     string `json:"invoice_name"`      // 发票名称
	InvoiceAddr     string `json:"invoice_addr"`      // 发票地址
	InvoiceBankName string `json:"invoice_bank_name"` // 发票银行名称
	InvoiceBankId   string `json:"invoice_bank_id"`   // 发票银行账号
	InvoiceBankNo   string `json:"invoice_bank_no"`   // 发票银行账号
	InvoiceTaxId    string `json:"invoice_tax_id"`    // 发票税号
	InvoicePhone    string `json:"invoice_phone"`     // 发票联系电话
	Scale           string `json:"scale"`             // 公司规模
	Theme           string `json:"theme"`             // 公司主题
	RoleId          int64  `json:"role_id"`           // 权限
	RoleName        string `json:"role_name"`         // 用户权限
	Code            int64  `json:"code"`              // 运输公司专用code 96543+Request::Id
	PrevSn          string `json:"prev_sn"`           // 自产产品，需要国家物码中心代码即：prev_sn
	Locked          bool   `json:"locked"`            // 是否锁定系统
	ServerStart     int64  `json:"server_start"`      // 系统启动时间
	ServerTime      int64  `json:"server_time"`       // 最后服务时间（试用需要后台自己处理）
	Intime          int64  `json:"intime"`            // 入库时间
	PayTimes        int64  `json:"pay_times"`         // 成功支付次数
	PayAmount       int64  `json:"pay_amount"`        // 支付总金额
	PayTime         int64  `json:"pay_time"`          // 最后一次支付时间
	Status          int64  `json:"status"`            // 1合作，0开发，-1退出
	ShopOpenIdInit  int64  `json:"shop_open_id_init"` // 门店初始编号
	ChannelId       string `json:"channel_id"`        // 渠道
	ChannelName     string `json:"channel_name"`
}
