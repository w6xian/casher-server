CREATE TABLE [mi_com_ticket_ids] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [type] TINYINT NOT NULL DEFAULT '1',
  [next_id] BIGINT NOT NULL DEFAULT '0',
  [date] VARCHAR(8) NOT NULL DEFAULT '',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [tic_date] ON [mi_com_ticket_ids] ([shop_id] ASC, [date] ASC);
CREATE TABLE [mi_com_ticket_ids_logs] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [next_id] BIGINT NOT NULL DEFAULT '0',
  [date] VARCHAR(8) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [tic_log_date] ON [mi_com_ticket_ids_logs] ([shop_id] ASC, [handler_id] ASC, [date] ASC);
CREATE TABLE [mi_com_jwt] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [casher_id] INT NOT NULL DEFAULT '0',
  [md5] VARCHAR(32) NOT NULL DEFAULT '',
  [token] VARCHAR(4000) NOT NULL DEFAULT '',
  [expire_time] INTEGER NOT NULL DEFAULT '0',
  [lock_status] TINYINT NOT NULL DEFAULT '0',
  [lock_time] INTEGER NOT NULL DEFAULT '0',
  [status] TINYINT NOT NULL DEFAULT '1',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [md5] ON [mi_com_jwt] ([md5] ASC);
-- 文件管理
CREATE TABLE [mi_com_shops_files] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [original] VARCHAR(255) NOT NULL DEFAULT '',
  [path] VARCHAR(255) NOT NULL DEFAULT '',
  [module] VARCHAR(45) NOT NULL DEFAULT '',
  [file_name] VARCHAR(255) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [file_size] BIGINT NOT NULL DEFAULT '0',
  [type] VARCHAR(45) NOT NULL DEFAULT '',
  [md5] VARCHAR(45) NOT NULL DEFAULT '',
  [status] TINYINT NOT NULL DEFAULT '1',
  [uptime] INTEGER NOT NULL DEFAULT '0',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [file_md5] ON [mi_com_shops_files] ([md5] ASC);
-- 分类表
CREATE TABLE [mi_com_shops_products_categories] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [pid] BIGINT NOT NULL DEFAULT '0',
  [image] VARCHAR(255) NOT NULL DEFAULT '',
  [pinyin] VARCHAR(45) NOT NULL DEFAULT '',
  [name] VARCHAR(45) NOT NULL,
  [mark] VARCHAR(64) NOT NULL DEFAULT '',
  [sort] INTEGER NOT NULL DEFAULT '50',
  [status] TINYINT NOT NULL DEFAULT '1',
  [uptime] INTEGER NOT NULL DEFAULT '0',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_suppliers] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [name] VARCHAR(45) NOT NULL,
  [pinyin] VARCHAR(45) NOT NULL DEFAULT '',
  [avat] VARCHAR(255) NOT NULL DEFAULT '',
  [area_code] VARCHAR(64) NOT NULL DEFAULT '',
  -- COMMENT '2位或者12位跟统计用区域'
  [area_path] VARCHAR(1024) NOT NULL DEFAULT '',
  -- COMMENT '带/的地址'
  [area_street] VARCHAR(64) NOT NULL DEFAULT '',
  -- COMMENT '用户自己填写的'
  [address] VARCHAR(100) NOT NULL DEFAULT '',
  -- COMMENT '全地址'
  [legal_persion] VARCHAR(20) NOT NULL DEFAULT '',
  -- COMMENT '法人'
  [contact_name] VARCHAR(20) NOT NULL DEFAULT '',
  -- COMMENT '联系人'
  [contact_phone] VARCHAR(20) NOT NULL DEFAULT '',
  -- COMMENT '联系方式'
  [contact_mobile] VARCHAR(20) NOT NULL DEFAULT '',
  -- COMMENT '第二联系方式'
  [contact_email] VARCHAR(200) NOT NULL DEFAULT '',
  -- COMMENT '联系人邮箱'
  [register_date] VARCHAR(45) NOT NULL DEFAULT '注册时间',
  [expire_date] VARCHAR(45) NOT NULL DEFAULT '失效时间',
  [categories_id] BIGINT NOT NULL DEFAULT '0',
  -- COMMENT '分类'
  [categories_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- COMMENT '分类名'
  [level_id] INTEGER NOT NULL DEFAULT '0',
  -- COMMENT '0-100'
  [level_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- COMMENT '等级名称'
  [avatar] VARCHAR(255) NOT NULL DEFAULT '',
  -- COMMENT '头像图片路径'
  [manager_id] BIGINT NOT NULL DEFAULT '0',
  [manager_name] VARCHAR(45) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- COMMENT 'mch_id'
  [mch_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- COMMENT 'mch_app'
  [mch_app] VARCHAR(45) NOT NULL DEFAULT '',
  -- COMMENT 'mch_pem'
  [mch_pem] VARCHAR(512) NOT NULL DEFAULT '',
  [mark] VARCHAR(64) NOT NULL DEFAULT '',
  [sort] INTEGER NOT NULL DEFAULT '50',
  [status] TINYINT NOT NULL DEFAULT '1',
  [uptime] INTEGER NOT NULL DEFAULT '0',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_products_brands] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [image] VARCHAR(255) NOT NULL DEFAULT '',
  [pinyin] VARCHAR(45) NOT NULL DEFAULT '',
  [name] VARCHAR(45) NOT NULL,
  [mark] VARCHAR(64) NOT NULL DEFAULT '',
  [sort] INTEGER NOT NULL DEFAULT '50',
  [status] TINYINT NOT NULL DEFAULT '1',
  [uptime] INTEGER NOT NULL DEFAULT '0',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_products_specs] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 0 预包装,1称重,2量体
  [style] BIGINT NOT NULL DEFAULT '0',
  [image] VARCHAR(255) NOT NULL DEFAULT '',
  [pinyin] VARCHAR(45) NOT NULL DEFAULT '',
  [name] VARCHAR(45) NOT NULL,
  [mark] VARCHAR(64) NOT NULL DEFAULT '',
  [sort] INTEGER NOT NULL DEFAULT '50',
  [status] TINYINT NOT NULL DEFAULT '1',
  [uptime] INTEGER NOT NULL DEFAULT '0',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [proxy_id] BIGINT NOT NULL DEFAULT '0',
  [app_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 同步服务器的仓库
  [store_id] BIGINT NOT NULL DEFAULT '0',
  -- 同步服务器的CRM中的客户ID
  [com_id] BIGINT NOT NULL DEFAULT '0',
  [avatar] VARCHAR(255) NOT NULL DEFAULT '',
  [name] VARCHAR(45) NOT NULL,
  [area_code] VARCHAR(64) NOT NULL DEFAULT '',
  -- COMMENT '2位或者12位跟统计用区域'
  [area_path] VARCHAR(1024) NOT NULL DEFAULT '',
  -- COMMENT '带/的地址'
  [area_street] VARCHAR(64) NOT NULL DEFAULT '',
  -- COMMENT '用户自己填写的'
  [address] VARCHAR(100) NOT NULL DEFAULT '',
  [longitude] VARCHAR(45) NOT NULL DEFAULT '',
  [latitude] VARCHAR(45) NOT NULL DEFAULT '',
  [geo_hash] VARCHAR(64) NOT NULL DEFAULT '',
  [chief_id] VARCHAR(45) NOT NULL DEFAULT '',
  [chief_name] VARCHAR(45) NOT NULL DEFAULT '',
  [mobile] VARCHAR(20) NOT NULL DEFAULT '',
  [type] TINYINT NOT NULL DEFAULT '4',
  [mark] VARCHAR(500) NOT NULL DEFAULT '',
  [status] TINYINT NOT NULL DEFAULT '0',
  [uptime] INTEGER NOT NULL DEFAULT '0',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_cashers] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [name] VARCHAR(45) NOT NULL DEFAULT '',
  [pinyin] VARCHAR(45) NOT NULL DEFAULT '',
  [mobile] VARCHAR(20) NOT NULL DEFAULT '',
  [image] VARCHAR(255) NOT NULL DEFAULT '',
  [password] VARCHAR(100) NOT NULL DEFAULT '',
  [is_leader] TINYINT NOT NULL DEFAULT '0',
  [last_login] INTEGER NOT NULL DEFAULT '0',
  [fail_times] INTEGER NOT NULL DEFAULT '0',
  [mark] VARCHAR(64) NOT NULL DEFAULT '',
  [sort] INTEGER NOT NULL DEFAULT '50',
  [uptime] INTEGER NOT NULL DEFAULT '0',
  [status] TINYINT NOT NULL DEFAULT '1',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
-- 交班表
CREATE TABLE [mi_com_shops_cashers_shift] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 收银员
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 值班经理（复核人员）
  [manager_id] BIGINT NOT NULL DEFAULT '0',
  [manager_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 财务人员（负责做账）
  [accountant_id] BIGINT NOT NULL DEFAULT '0',
  [accountant_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- ^订单金额统计
  -- 应收（原价）
  [total] BIGINT NOT NULL DEFAULT '0',
  -- 实收()
  [dr] BIGINT NOT NULL DEFAULT '0',
  -- 应付这里需要与dr相同
  [cr] BIGINT NOT NULL DEFAULT '0',
  -- 结存现金收款
  [cash] BIGINT NOT NULL DEFAULT '0',
  [cash_num] INTEGER NOT NULL DEFAULT '0',
  -- 结存支付宝收款
  [alipay] BIGINT NOT NULL DEFAULT '0',
  [alipay_num] INTEGER NOT NULL DEFAULT '0',
  -- 结存微信收款
  [wechat] BIGINT NOT NULL DEFAULT '0',
  [wechat_num] INTEGER NOT NULL DEFAULT '0',
  -- 结存银行卡收款
  [bank] BIGINT NOT NULL DEFAULT '0',
  [bank_num] INTEGER NOT NULL DEFAULT '0',
  -- 结存其他收款
  [other] BIGINT NOT NULL DEFAULT '0',
  [other_num] INTEGER NOT NULL DEFAULT '0',
  -- dr  = cash+alipay+wechat+bank+other
  -- ^额外
  -- 充值
  [recharge] BIGINT NOT NULL DEFAULT '0',
  -- 
  -- 交班订单数
  [order_num] BIGINT NOT NULL DEFAULT '0',
  -- 期初现金
  [begin_cash] BIGINT NOT NULL DEFAULT '0',
  -- 提现金额（不能存太多金额，需要把现金上交到财务单独做账）
  [withdraw] BIGINT NOT NULL DEFAULT '0',
  -- 期末现金
  [end_cash] BIGINT NOT NULL DEFAULT '0',
  -- end_cash = begin_cash + cash - withdraw
  [machine_id] VARCHAR(45) NOT NULL DEFAULT '',
  [prints] INTEGER NOT NULL DEFAULT '0',
  -- 未支付订单数
  [unpaid_num] INTEGER NOT NULL DEFAULT '0',
  -- 未支付订单金额
  [unpaid_fee] INTEGER NOT NULL DEFAULT '0',
  -- 禁用订单数
  [disabled_num] INTEGER NOT NULL DEFAULT '0',
  -- 禁用订单金额
  [disabled_fee] INTEGER NOT NULL DEFAULT '0',
  -- 上班时间
  [begin_time] INTEGER NOT NULL DEFAULT '0',
  -- 交班时间
  [end_time] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '0',
  -- 复核时间
  [judge_time] INTEGER NOT NULL DEFAULT '0',
  -- 做账时间
  [account_time] INTEGER NOT NULL DEFAULT '0',
  -- 打印次数
  [print_times] INTEGER NOT NULL DEFAULT '0',
  -- 银行类型
  [bank_type] VARCHAR(45) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_shift_shop_id_intime] ON [mi_com_shops_cashers_shift] ([shop_id] ASC, [intime] ASC);
CREATE INDEX [idx_shift_handler_id_intime] ON [mi_com_shops_cashers_shift] ([handler_id] ASC, [intime] ASC);
-- 登录
CREATE TABLE [mi_com_shops_cashers_online] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [machine_id] VARCHAR(45) NOT NULL DEFAULT '',
  [login_time] INTEGER NOT NULL DEFAULT '0',
  [logout_time] INTEGER NOT NULL DEFAULT '0',
  -- 上一次登录是谁
  [prev_id] BIGINT NOT NULL DEFAULT '0',
  [prev_name] VARCHAR(45) NOT NULL DEFAULT '',
  --上期结存现金
  [prev_cash] BIGINT NOT NULL DEFAULT '0',
  [cash] BIGINT NOT NULL DEFAULT '0',
  -- 本期结存现金多少
  [next_cash] BIGINT NOT NULL DEFAULT '0',
  -- 备注
  [spell_mark] VARCHAR(255) NOT NULL DEFAULT '',
  -- 接班差额
  [spell_diff] BIGINT NOT NULL DEFAULT '0',
  -- 交班差额
  [shift_diff] BIGINT NOT NULL DEFAULT '0',
  -- 备注
  [shift_mark] VARCHAR(255) NOT NULL DEFAULT '',
  -- 当前状态 1表示登录中 0表示已退出，2表示已交班
  [type] TINYINT NOT NULL DEFAULT '1',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [csco_idx_shop_mch_intime] ON [mi_com_shops_cashers_online] ([shop_id] ASC, [machine_id] ASC);
-- 登录日志
CREATE TABLE [mi_com_shops_cashers_login_log] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [casher_id] BIGINT NOT NULL DEFAULT '0',
  [machine_id] VARCHAR(45) NOT NULL DEFAULT '',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_address_2023] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [name] VARCHAR(255) NOT NULL DEFAULT '',
  [code] VARCHAR(64) NOT NULL DEFAULT '',
  [parent_id] VARCHAR(64) DEFAULT '',
  [level] VARCHAR(16) NOT NULL DEFAULT '',
  [full_path] VARCHAR(1024) NOT NULL DEFAULT '',
  [loc_id] VARCHAR(64) NOT NULL DEFAULT '',
  [created_at] VARCHAR(64) NOT NULL DEFAULT '',
  [updated_at] VARCHAR(64) NOT NULL DEFAULT '',
  [deleted_at] VARCHAR(64) DEFAULT '',
  [created_by] VARCHAR(64) NOT NULL DEFAULT '',
  [updated_by] VARCHAR(64) NOT NULL DEFAULT '',
  [status] VARCHAR(64) NOT NULL DEFAULT '',
  [version] INTEGER(11) NOT NULL DEFAULT '1'
);
CREATE INDEX [idx_loc_id] ON [mi_com_shops_address_2023] ([loc_id] ASC);
CREATE INDEX [idx_loc_code] ON [mi_com_shops_address_2023] ([code] ASC);
CREATE TABLE [mi_com_shops_tutors] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [code] VARCHAR(100) NOT NULL DEFAULT '',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [user_id] BIGINT NOT NULL DEFAULT '0',
  [name] VARCHAR(45) NOT NULL DEFAULT '',
  [mobile] VARCHAR(20) NOT NULL DEFAULT '',
  [mark] VARCHAR(64) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_code] ON [mi_com_shops_tutors] ([code] ASC);
CREATE INDEX [idx_mobile] ON [mi_com_shops_tutors] ([mobile] ASC);
-- 订单
CREATE TABLE [mi_com_shops_orders] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  [track_str] VARCHAR(45) NOT NULL DEFAULT '',
  [ticket] VARCHAR(45) NOT NULL DEFAULT '',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [machine_id] VARCHAR(45) NOT NULL DEFAULT '',
  [date_time] INTEGER NOT NULL DEFAULT '0',
  [dr] BIGINT NOT NULL DEFAULT '0',
  [cr] BIGINT NOT NULL DEFAULT '0',
  [off] BIGINT NOT NULL DEFAULT '0',
  [off_price] BIGINT NOT NULL DEFAULT '0',
  [abatement] BIGINT NOT NULL DEFAULT '0',
  [debit] BIGINT NOT NULL DEFAULT '0',
  [discount] BIGINT NOT NULL DEFAULT '0',
  [change] BIGINT NOT NULL DEFAULT '0',
  [coupons] INTEGER NOT NULL DEFAULT '0',
  [points] INTEGER NOT NULL DEFAULT '0',
  [balance] BIGINT NOT NULL DEFAULT '0',
  [payed] BIGINT NOT NULL DEFAULT '0',
  [prd_num] INTEGER NOT NULL DEFAULT '0',
  [shop_user_id] BIGINT NOT NULL DEFAULT '0',
  [user_name] VARCHAR(45) NOT NULL DEFAULT '',
  [user_mobile] VARCHAR(20) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [mark] VARCHAR(45) NOT NULL DEFAULT '',
  [prints] INTEGER NOT NULL DEFAULT '0',
  -- 支付方式（转账）当面支付
  -- 支付方式（收现）1：现金，2：支付宝，3：微信，4：银行卡，5：其他
  [currency] VARCHAR(45) NOT NULL DEFAULT '',
  [pay_type] VARCHAR(45) NOT NULL DEFAULT '',
  -- 支付状态（0：未支付，1：已支付）
  [pay_status] TINYINT NOT NULL DEFAULT '0',
  -- 支付时间
  [pay_time] INTEGER NOT NULL DEFAULT '0',
  -- 支付金额
  [pay_total] BIGINT NOT NULL DEFAULT '0',
  -- 打印次数
  [print_times] INTEGER NOT NULL DEFAULT '0',
  -- '退款金额',
  [refund_amount] BIGINT NOT NULL DEFAULT '0',
  -- '退款次数',
  [refund_times] INTEGER NOT NULL DEFAULT '0',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
-- 订单详情
CREATE TABLE [mi_com_shops_orders_items] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [store_id] BIGINT NOT NULL DEFAULT '0',
  [store_name] VARCHAR(45) NOT NULL DEFAULT '',
  [order_id] BIGINT NOT NULL,
  [prd_sn] VARCHAR(45) NOT NULL DEFAULT '',
  [prd_avatar] VARCHAR(255) NOT NULL DEFAULT '',
  [prd_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec] INTEGER NOT NULL DEFAULT '0',
  [weight] INTEGER NOT NULL DEFAULT '0',
  [style] TINYINT NOT NULL DEFAULT '1',
  [pack_name] VARCHAR(8) NOT NULL DEFAULT '箱',
  [style_type] TINYINT NOT NULL DEFAULT '0',
  [times] INTEGER NOT NULL DEFAULT '0',
  [debit] BIGINT NOT NULL DEFAULT '0',
  [off] BIGINT NOT NULL DEFAULT '0',
  [abatement] BIGINT NOT NULL DEFAULT '0',
  [coupons] BIGINT NOT NULL DEFAULT '0',
  [points] BIGINT NOT NULL DEFAULT '0',
  [balance] BIGINT NOT NULL DEFAULT '0',
  [price] BIGINT NOT NULL DEFAULT '0',
  [num] BIGINT NOT NULL DEFAULT '0',
  [total] BIGINT NOT NULL DEFAULT '0',
  [discount] BIGINT NOT NULL DEFAULT '0',
  [payed] BIGINT NOT NULL DEFAULT '0',
  [mark] VARCHAR(200) NOT NULL DEFAULT '',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
-- 订单支付
CREATE TABLE [mi_com_shops_orders_pay] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [appid] VARCHAR(45) NOT NULL DEFAULT '0',
  [mchid] VARCHAR(45) NOT NULL DEFAULT '0',
  [user_id] BIGINT NOT NULL DEFAULT '0',
  [openid] VARCHAR(64) NOT NULL DEFAULT '',
  [title] VARCHAR(255) NOT NULL DEFAULT '',
  -- '交易类型 单号/type=1,就是订单号',
  [order_id] BIGINT NOT NULL DEFAULT '0',
  -- '金额(单位是分，1表示1分）',
  [amount] BIGINT NOT NULL DEFAULT '0',
  -- '交易平台（1为微信2为支付宝）',
  [platform] TINYINT NOT NULL DEFAULT '1',
  [remark] VARCHAR(255) NOT NULL DEFAULT '',
  [attach] VARCHAR(255) NOT NULL DEFAULT '',
  -- '1：申请退款中2：退款完成',
  [refund] TINYINT NOT NULL DEFAULT '0',
  -- '退款金额',
  [refund_amount] BIGINT NOT NULL DEFAULT '0',
  [refund_time] INTEGER NOT NULL DEFAULT '0',
  -- '退款操作人',
  [refund_handler_id] INTEGER NOT NULL DEFAULT '0',
  -- '退款操作人姓名',
  [refund_handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- '微信支付系统生成的订单号。',
  [transaction_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- '这里传给服务器的单号',
  [out_trade_no] VARCHAR(45) NOT NULL DEFAULT '',
  -- '支付完成时间，遵循rfc3339标准格式',
  [success_time] VARCHAR(64) NOT NULL DEFAULT '0',
  -- '交易状态，枚举值：SUCCESS：支付成功, REFUND：转入退款, NOTPAY：未支付, REVOKED：已撤销（付款码支付）, USERPAYING：用户支付中（付款码支付）, PAYERROR：支付失败(其他原因，如银行返回失败)',
  [trade_state] VARCHAR(45) NOT NULL DEFAULT '',
  --  '交易状态描述',
  [trade_state_desc] VARCHAR(256) NOT NULL DEFAULT '',
  [trade_type] VARCHAR(45) NOT NULL DEFAULT '',
  -- 支付授权码
  [pay_auth_code] VARCHAR(45) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- '状态（0：未完成交易1：完成关键交易）',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_out_trade_no] ON [mi_com_shops_orders_pay] ([out_trade_no] ASC, [shop_id] ASC);
-- '支付当面付日志',
-- 订单退款
CREATE TABLE [mi_com_shops_refund_pay] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [app_id] VARCHAR(45) NOT NULL DEFAULT '0',
  [mch_id] VARCHAR(45) NOT NULL DEFAULT '0',
  [openid] VARCHAR(64) NOT NULL DEFAULT '',
  [user_id] BIGINT NOT NULL DEFAULT '0',
  [title] VARCHAR(255) NOT NULL DEFAULT '',
  -- 订单号
  [local_order_id] BIGINT NOT NULL DEFAULT '0',
  -- 支付号
  [local_pay_id] BIGINT NOT NULL DEFAULT '0',
  -- refund表中的ID
  [local_refund_id] BIGINT NOT NULL DEFAULT '0',
  -- 这个是payTotal金额 （相当于订单支付总金额）退款不可能超过
  [order_payed] BIGINT NOT NULL DEFAULT '0',
  -- 退款号
  [refund_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 商户退款单号
  [out_refund_no] VARCHAR(45) NOT NULL DEFAULT '',
  -- 微信支付系统生成的订单号
  [transaction_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
  [out_trade_no] VARCHAR(45) NOT NULL DEFAULT '',
  -- 退款渠道
  [channel] VARCHAR(45) NOT NULL DEFAULT '',
  -- 退款入账账户
  [user_received_account] VARCHAR(45) NOT NULL DEFAULT '',
  -- 退款成功时间
  [success_time] VARCHAR(45) NOT NULL DEFAULT '',
  -- 退款创建时间
  [create_time] VARCHAR(45) NOT NULL DEFAULT '',
  -- 退款状态
  [status] VARCHAR(45) NOT NULL DEFAULT '',
  -- 资金账户
  [funds_account] VARCHAR(45) NOT NULL DEFAULT '',
  -- 订单总金额，单位为分
  [total] BIGINT NOT NULL DEFAULT '0',
  -- 退款标价金额，单位为分，可以做部分退款
  [refund] BIGINT NOT NULL DEFAULT '0',
  -- 用户支付金额，单位为分
  [payer_total] BIGINT NOT NULL DEFAULT '0',
  -- 用户退款金额，单位为分
  [payer_refund] BIGINT NOT NULL DEFAULT '0',
  -- 去掉非充值代金券退款金额后的退款金额，单位为分，退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
  [settlement_refund] BIGINT NOT NULL DEFAULT '0',
  -- 应结订单金额=订单金额-免充值代金券金额，应结订单金额<=订单金额，单位为分
  [settlement_total] BIGINT NOT NULL DEFAULT '0',
  -- 优惠退款金额,单位为分
  [discount_refund] BIGINT NOT NULL DEFAULT '0',
  -- CNY：人民币，境内商户号仅支持人民币
  [currency] VARCHAR(45) NOT NULL DEFAULT '',
  -- 手续费退款金额，单位为分
  [refund_fee] BIGINT NOT NULL DEFAULT '0',
  -- 处理人
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
-- 提现记录
CREATE TABLE [mi_com_shops_cash_withdraw] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [ticket] VARCHAR(45) NOT NULL DEFAULT '',
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 店铺
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 机器编码
  [machine_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 提现标题
  [title] VARCHAR(255) NOT NULL DEFAULT '',
  -- 变化
  [prev_cash] BIGINT NOT NULL DEFAULT '0',
  [cash] BIGINT NOT NULL DEFAULT '0',
  [next_cash] BIGINT NOT NULL DEFAULT '0',
  -- 处理人(cahserId)
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 备注
  [remark] VARCHAR(255) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_withdraw_track_id] ON [mi_com_shops_cash_withdraw] ([track_id] ASC, [shop_id] ASC);
CREATE TABLE [mi_com_shops_cash_topup] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [ticket] VARCHAR(45) NOT NULL DEFAULT '',
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 店铺
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 机器编码
  [machine_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 提现标题
  [title] VARCHAR(255) NOT NULL DEFAULT '',
  -- 变化
  [prev_cash] BIGINT NOT NULL DEFAULT '0',
  [cash] BIGINT NOT NULL DEFAULT '0',
  [next_cash] BIGINT NOT NULL DEFAULT '0',
  -- 处理人(cahserId)
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 备注
  [remark] VARCHAR(255) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_topup_track_id] ON [mi_com_shops_cash_topup] ([track_id] ASC, [shop_id] ASC);
-- 现金变化记录
CREATE TABLE [mi_com_shops_cash_log] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [ticket] VARCHAR(45) NOT NULL DEFAULT '',
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 店铺
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 机器编码
  [machine_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 1订单，-1提现,-2退货,-3报损
  [target_type] BIGINT NOT NULL DEFAULT '0',
  [target_id] BIGINT NOT NULL DEFAULT '0',
  -- 提现标题
  [title] VARCHAR(255) NOT NULL DEFAULT '',
  [prev_cash] BIGINT NOT NULL DEFAULT '0',
  [cash] BIGINT NOT NULL DEFAULT '0',
  [next_cash] BIGINT NOT NULL DEFAULT '0',
  -- 处理人(cahserId)
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 备注
  [remark] VARCHAR(255) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_cash_log_track_id] ON [mi_com_shops_cash_log] ([track_id] ASC, [shop_id] ASC);
-- 支付（不是当面付）
CREATE TABLE [mi_com_shops_payment] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  [nonce_str] VARCHAR(32) NOT NULL DEFAULT '',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [app_id] VARCHAR(45) NOT NULL DEFAULT '0',
  [user_id] BIGINT NOT NULL DEFAULT '0',
  [openid] VARCHAR(64) NOT NULL DEFAULT '',
  [title] VARCHAR(255) NOT NULL,
  [type] INTEGER NOT NULL DEFAULT '0',
  [type_id] BIGINT NOT NULL DEFAULT '0',
  [amount] BIGINT NOT NULL DEFAULT '0',
  [platform] TINYINT NOT NULL DEFAULT '1',
  [remark] VARCHAR(255) NOT NULL,
  [attach] VARCHAR(255) NOT NULL,
  [pay_time] INTEGER NOT NULL DEFAULT '0',
  [refund] TINYINT NOT NULL DEFAULT '0',
  [refund_time] TINYINT NOT NULL DEFAULT '0',
  [notify_transaction_id] VARCHAR(45) NOT NULL DEFAULT '',
  [notify_success_time] VARCHAR(45) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_payment_wechat_notify] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  [notify_id] VARCHAR(45) NOT NULL DEFAULT '',
  [create_time] VARCHAR(45) NOT NULL DEFAULT '',
  [event_type] VARCHAR(45) NOT NULL DEFAULT '',
  [resource_type] VARCHAR(45) NOT NULL DEFAULT '',
  [resource] VARCHAR(4500) NOT NULL DEFAULT '',
  [res_out_trade_no] VARCHAR(45) NOT NULL DEFAULT '',
  [res_transaction_id] VARCHAR(45) NOT NULL DEFAULT '',
  [res_trade_state] VARCHAR(45) NOT NULL DEFAULT '',
  [res_total] INTEGER NOT NULL DEFAULT '0',
  [res_payed] INTEGER NOT NULL DEFAULT '0',
  [success_time] VARCHAR(45) NOT NULL DEFAULT '',
  [summary] VARCHAR(66) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 状态
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
-- 商品表（SPU）
CREATE TABLE [mi_com_shops_products] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [open_id] VARCHAR(45) NOT NULL DEFAULT '',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  -- 13位
  [sn] VARCHAR(45) NOT NULL DEFAULT '',
  [code] VARCHAR(15) NOT NULL DEFAULT '',
  [avatar] VARCHAR(255) NOT NULL DEFAULT '',
  [cover] VARCHAR(255) NOT NULL DEFAULT '',
  [name] VARCHAR(64) NOT NULL DEFAULT '',
  [pinyin] VARCHAR(64) NOT NULL DEFAULT '',
  -- 200gx20袋
  -- 200
  [spec] INTEGER NOT NULL DEFAULT '0',
  -- g 
  [unit] VARCHAR NOT NULL DEFAULT 'g',
  -- 20
  [pk_amount] INTEGER NOT NULL DEFAULT '0',
  -- 袋
  [spec_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- '最小规格重量（含包装）',
  [spec_weight] INTEGER NOT NULL DEFAULT '0',
  -- end 200gx20袋
  -- 包装方式
  [pack_name] VARCHAR(8) NOT NULL DEFAULT '箱',
  -- '装箱后总重量（包含包装）',
  [pk_weight] BIGINT NOT NULL DEFAULT '1',
  -- '包装方式1预包装，2散装'
  [style] TINYINT NOT NULL DEFAULT '1',
  -- '1 称重2 量体3 点数
  [style_type] TINYINT NOT NULL DEFAULT '0',
  -- 卖点
  [feature] VARCHAR(200) NOT NULL DEFAULT '',
  -- 售卖价
  [price] BIGINT NOT NULL DEFAULT '0',
  -- 划线价
  [line_price] BIGINT NOT NULL DEFAULT '0',
  -- 显示销售量
  [sale_num] BIGINT NOT NULL DEFAULT '0',
  -- 销售方式(pk_amount,1或者提定，单次扣库存的数量)
  [times] INTEGER NOT NULL DEFAULT '1',
  -- 库存
  [stock] BIGINT NOT NULL DEFAULT '0',
  -- 打标签 用于销售显示
  [tags] VARCHAR(45) NOT NULL DEFAULT '[]',
  -- 销售税率
  [tax] INTEGER NOT NULL DEFAULT '0',
  -- 只有后台可见的备注
  [mark] VARCHAR(45) NOT NULL DEFAULT '',
  -- 显示位置
  [sort] INTEGER NOT NULL DEFAULT '50',
  -- 来源（0 自产 1 从供应商）
  [source] INTEGER NOT NULL DEFAULT '0',
  -- --  类型
  -- '1/商品2/源辅料/3物料/4包装物',
  [type] INTEGER NOT NULL DEFAULT '0',
  [units] VARCHAR(250) NOT NULL DEFAULT '[]',
  -- 处理人
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 第一供应商
  [supplier_id] BIGINT NOT NULL DEFAULT '0',
  [supplier_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 第一分类
  [category_id] BIGINT NOT NULL DEFAULT '0',
  [category_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 第一品牌
  [brand_id] BIGINT NOT NULL DEFAULT '0',
  [brand_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 成本
  [unit_price] BIGINT NOT NULL DEFAULT '0',
  [pack_price] BIGINT NOT NULL DEFAULT '0',
  -- 成本
  [cost] BIGINT NOT NULL DEFAULT '0',
  --保质期（天）
  [keep_life] INTEGER NOT NULL DEFAULT '0',
  -- 保质期单位
  [keep_life_unit] VARCHAR(8) NOT NULL DEFAULT '',
  -- 状态
  [status] INTEGER NOT NULL DEFAULT '0',
  -- 最更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 创建时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
-- 商品库存表
CREATE TABLE [mi_com_shops_products_stock] (
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 商品ID SPU
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  -- 商品ID sku
  [sku_id] BIGINT NOT NULL DEFAULT '0',
  -- 数量
  [stock] INTEGER NOT NULL DEFAULT '0',
  -- 批号
  [batch] VARCHAR(255) NOT NULL DEFAULT '',
  --保质期（天）
  [keep_life] INTEGER NOT NULL DEFAULT '0',
  -- 保质期单位
  [keep_life_unit] VARCHAR(8) NOT NULL DEFAULT '',
  -- 最更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 创建时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [sps_s_idx_prd_id] ON [mi_com_shops_products_stock] ([prd_id] ASC);
CREATE INDEX [sps_s_idx_sku_id] ON [mi_com_shops_products_stock] ([sku_id] ASC);
-- 商品库存表
CREATE TABLE [mi_com_shops_products_store] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 商品ID SPU
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  -- 商品ID sku
  [sku_id] BIGINT NOT NULL DEFAULT '0',
  -- 数量
  [stock] INTEGER NOT NULL DEFAULT '0',
  -- 批号
  [batch] VARCHAR(255) NOT NULL DEFAULT '',
  --保质期（天）
  [keep_life] INTEGER NOT NULL DEFAULT '0',
  -- 保质期单位
  [keep_life_unit] VARCHAR(8) NOT NULL DEFAULT '',
  -- 最更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 创建时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [sps_s_idx_prd_id] ON [mi_com_shops_products_store] ([prd_id] ASC);
CREATE INDEX [sps_s_idx_sku_id] ON [mi_com_shops_products_store] ([sku_id] ASC);
-- store变化日志库
CREATE TABLE [mi_com_shops_products_store_log] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [sku_id] BIGINT NOT NULL DEFAULT '0',
  [event] VARCHAR(45) NOT NULL DEFAULT '',
  -- 操作类型 >0 为增加<0表示减少符号应和num相同。
  -- -1 销售 1 进货（采购）
  -- -2 报损 2 报溢
  [event_type] INTEGER NOT NULL DEFAULT '0',
  -- 变化前
  [prev_stock] INTEGER NOT NULL DEFAULT '0',
  -- 变化数量
  [num] INTEGER NOT NULL DEFAULT '0',
  -- 变化后
  [next_stock] INTEGER NOT NULL DEFAULT '0',
  -- 变化前
  [prev_cost] INTEGER NOT NULL DEFAULT '0',
  -- 变化数量
  [cost] INTEGER NOT NULL DEFAULT '0',
  -- 变化后
  [next_cost] INTEGER NOT NULL DEFAULT '0',
  -- 批号
  [batch] VARCHAR(255) NOT NULL DEFAULT '',
  -- 处理人
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [mark] VARCHAR(45) NOT NULL DEFAULT '',
  -- 最更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 1有效
  [status] INTEGER NOT NULL DEFAULT '0',
  -- 创建时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [spsl_s_idx_prd_id] ON [mi_com_shops_products_store_log] ([prd_id] ASC);
CREATE INDEX [spsl_s_idx_sku_id] ON [mi_com_shops_products_store_log] ([sku_id] ASC);
CREATE INDEX [spsl_s_idx_track_id] ON [mi_com_shops_products_store_log] ([track_id] ASC);
-- 分类属性分组
CREATE TABLE [mi_com_shops_products_attr_group] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [category_id] BIGINT NOT NULL DEFAULT '0',
  [name] VARCHAR(255) NOT NULL DEFAULT '',
  [pinyin] VARCHAR(255) NOT NULL DEFAULT '',
  -- 最更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 创建时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [psag_idx_c_id] ON [mi_com_shops_products_attr_group] ([category_id] ASC);
-- 分类属性参数
-- 根据group分组
CREATE TABLE [mi_com_shops_products_attr_param] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [category_id] BIGINT NOT NULL DEFAULT '0',
  [group_id] BIGINT NOT NULL DEFAULT '0',
  [name] VARCHAR(255) NOT NULL DEFAULT '',
  [pinyin] VARCHAR(255) NOT NULL DEFAULT '',
  -- 通用属性 1是 0否(sku独有属性)
  [generic] TINYINT NOT NULL DEFAULT '0',
  -- 值需要解析成的类型 1float，2string,3array,4json
  [type] TINYINT NOT NULL DEFAULT '0',
  -- 最更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 创建时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [psav_idx_c_id] ON [mi_com_shops_products_attr_param] ([category_id] ASC);
CREATE INDEX [psav_idx_g_id] ON [mi_com_shops_products_attr_param] ([group_id] ASC);
-- 分类值
CREATE TABLE [mi_com_shops_products_attr_values] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [name] VARCHAR(255) NOT NULL DEFAULT '',
  [pinyin] VARCHAR(255) NOT NULL DEFAULT '',
  -- 最更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 创建时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [psav_idx_prd_id] ON [mi_com_shops_products_attr_values] ([prd_id] ASC);
-- 商品表详情
CREATE TABLE [mi_com_shops_products_detail] (
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 商品ID SPU
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  -- 名称
  [description] TEXT NOT NULL DEFAULT '',
  -- 通用属性
  [gen_attr] VARCHAR(1000) NOT NULL DEFAULT '',
  -- 单独属性
  [sku_attr] VARCHAR(1000) NOT NULL DEFAULT ''
);
CREATE INDEX [spk_d_idx_prd_id] ON [mi_com_shops_products_detail] ([prd_id] ASC);
-- 商品表（SKU）
CREATE TABLE [mi_com_shops_products_sku] (
  [sku_id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [sku_shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 商品ID SPU
  [sku_prd_id] BIGINT NOT NULL DEFAULT '0',
  -- 名称
  [sku_name] VARCHAR(255) NOT NULL DEFAULT '',
  -- 拼音
  [sku_pinyin] VARCHAR(255) NOT NULL DEFAULT '',
  -- 图片
  [sku_avatar] VARCHAR(255) NOT NULL DEFAULT '',
  [sku_cover] VARCHAR(255) NOT NULL DEFAULT '',
  -- 价格
  [sku_price] BIGINT NOT NULL DEFAULT '0',
  -- 成本
  [sku_cost] BIGINT NOT NULL DEFAULT '0',
  -- 根据SPU中单独属性sku_attr 索引 [0_0_0]
  [sku_attr_index] VARCHAR(255) NOT NULL DEFAULT '',
  -- 根据单独属性单独值 json格式
  [sku_attr_desc] VARCHAR(1000) NOT NULL DEFAULT '',
  -- 有效性
  [sku_status] TINYINT NOT NULL DEFAULT '1',
  -- 最更新时间
  [sku_uptime] INTEGER NOT NULL DEFAULT '0',
  -- 创建时间
  [sku_intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [spk_s_idx_prd_id] ON [mi_com_shops_products_sku] ([sku_prd_id] ASC);
CREATE INDEX [spk_s_idx_attr_index] ON [mi_com_shops_products_sku] ([sku_attr_index] ASC);
--商品与供应商的关系
CREATE TABLE [mi_com_shops_products_suppliers_ref] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [supp_id] BIGINT NOT NULL DEFAULT '0',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [ps_idx_prd_id] ON [mi_com_shops_products_suppliers_ref] ([prd_id] ASC);
CREATE INDEX [ps_idx_supp_id] ON [mi_com_shops_products_suppliers_ref] ([supp_id] ASC);
CREATE INDEX [ps_idx_shop_id] ON [mi_com_shops_products_suppliers_ref] ([shop_id] ASC);
CREATE TABLE [mi_com_shops_products_comments] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  [rate_convey] TINYINT NOT NULL DEFAULT '1',
  [rate_good] TINYINT NOT NULL DEFAULT '1',
  [rate_service] TINYINT NOT NULL DEFAULT '1',
  [anonymous] TINYINT NOT NULL DEFAULT '0',
  [order_id] BIGINT NOT NULL DEFAULT '0',
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [spec_info] VARCHAR(45) NOT NULL DEFAULT '',
  [memo] VARCHAR(256) NOT NULL DEFAULT '',
  [user_id] BIGINT NOT NULL DEFAULT '0',
  [user_name] VARCHAR(45) NOT NULL DEFAULT '',
  [user_avatar] VARCHAR(200) NOT NULL DEFAULT '',
  [judge_id] INTEGER NOT NULL DEFAULT '0',
  [judge_status] TINYINT NOT NULL DEFAULT '0',
  [judge_time] INTEGER NOT NULL DEFAULT '0',
  [type] TINYINT NOT NULL DEFAULT '1',
  [has_images] TINYINT NOT NULL DEFAULT '0',
  [auto_comment] TINYINT NOT NULL DEFAULT '0',
  [seller_reply] VARCHAR(200) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_products_comments_medias] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [order_id] BIGINT NOT NULL DEFAULT '0',
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [user_id] BIGINT NOT NULL DEFAULT '0',
  [comment_id] BIGINT NOT NULL DEFAULT '0',
  [url] VARCHAR(200) NOT NULL DEFAULT ''
);
-- 挂单
CREATE TABLE [mi_com_shops_temp] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  [ticket] VARCHAR(45) NOT NULL DEFAULT '',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [machine_id] VARCHAR(45) NOT NULL DEFAULT '',
  [date_time] INTEGER NOT NULL DEFAULT '0',
  [dr] BIGINT NOT NULL DEFAULT '0',
  [cr] BIGINT NOT NULL DEFAULT '0',
  [off] BIGINT NOT NULL DEFAULT '0',
  [off_price] BIGINT NOT NULL DEFAULT '0',
  [abatement] BIGINT NOT NULL DEFAULT '0',
  [debit] BIGINT NOT NULL DEFAULT '0',
  [discount] BIGINT NOT NULL DEFAULT '0',
  [change] BIGINT NOT NULL DEFAULT '0',
  [coupons] INTEGER NOT NULL DEFAULT '0',
  [points] INTEGER NOT NULL DEFAULT '0',
  [balance] BIGINT NOT NULL DEFAULT '0',
  [payed] BIGINT NOT NULL DEFAULT '0',
  [prd_num] INTEGER NOT NULL DEFAULT '0',
  [shop_user_id] BIGINT NOT NULL DEFAULT '0',
  [user_name] VARCHAR(45) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [mark] VARCHAR(45) NOT NULL DEFAULT '',
  [prints] INTEGER NOT NULL DEFAULT '0',
  [pay_type] VARCHAR(45) NOT NULL DEFAULT '',
  [pay_status] TINYINT NOT NULL DEFAULT '0',
  [rec_time] INTEGER NOT NULL DEFAULT '0',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_sh] ON [mi_com_shops_temp] ([shop_id] ASC, [handler_id] ASC);
CREATE INDEX [idx_ticket] ON [mi_com_shops_temp] ([ticket] ASC);
-- 挂单商品列表
CREATE TABLE [mi_com_shops_temp_items] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [temp_id] BIGINT NOT NULL,
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [prd_sn] VARCHAR(45) NOT NULL DEFAULT '',
  [prd_avatar] VARCHAR(255) NOT NULL DEFAULT '',
  [prd_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec] INTEGER NOT NULL DEFAULT '0',
  [weight] INTEGER NOT NULL DEFAULT '0',
  [style] TINYINT NOT NULL DEFAULT '1',
  [pack_name] VARCHAR(8) NOT NULL DEFAULT '箱',
  [style_type] TINYINT NOT NULL DEFAULT '0',
  [times] INTEGER NOT NULL DEFAULT '0',
  [debit] BIGINT NOT NULL DEFAULT '0',
  [off] BIGINT NOT NULL DEFAULT '0',
  [abatement] BIGINT NOT NULL DEFAULT '0',
  [coupons] BIGINT NOT NULL DEFAULT '0',
  [points] BIGINT NOT NULL DEFAULT '0',
  [balance] BIGINT NOT NULL DEFAULT '0',
  [price] BIGINT NOT NULL DEFAULT '0',
  [num] BIGINT NOT NULL DEFAULT '0',
  [total] BIGINT NOT NULL DEFAULT '0',
  [discount] BIGINT NOT NULL DEFAULT '0',
  [payed] BIGINT NOT NULL DEFAULT '0',
  [mark] VARCHAR(200) NOT NULL DEFAULT '',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_tmp_id] ON [mi_com_shops_temp_items] ([temp_id] ASC);
CREATE INDEX [idx_tmp_si] ON [mi_com_shops_temp_items] ([shop_id] ASC);
CREATE INDEX [idx_tmp_handler_id] ON [mi_com_shops_temp_items] ([shop_id] ASC, [handler_id] ASC) -- 退货
CREATE TABLE [mi_com_shops_refund] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  [track_str] VARCHAR(45) NOT NULL DEFAULT '',
  [ticket] VARCHAR(45) NOT NULL DEFAULT '',
  -- 原订单ID
  [order_id] BIGINT NOT NULL DEFAULT '0',
  -- 原订单号
  [order_ticket] VARCHAR(45) NOT NULL DEFAULT '',
  [machine_id] VARCHAR(45) NOT NULL DEFAULT '',
  [date_time] INTEGER NOT NULL DEFAULT '0',
  [dr] BIGINT NOT NULL DEFAULT '0',
  [cr] BIGINT NOT NULL DEFAULT '0',
  [off] BIGINT NOT NULL DEFAULT '0',
  [off_price] BIGINT NOT NULL DEFAULT '0',
  [abatement] BIGINT NOT NULL DEFAULT '0',
  [debit] BIGINT NOT NULL DEFAULT '0',
  [discount] BIGINT NOT NULL DEFAULT '0',
  [change] BIGINT NOT NULL DEFAULT '0',
  [coupons] INTEGER NOT NULL DEFAULT '0',
  [points] INTEGER NOT NULL DEFAULT '0',
  [balance] BIGINT NOT NULL DEFAULT '0',
  [payed] BIGINT NOT NULL DEFAULT '0',
  [prd_num] INTEGER NOT NULL DEFAULT '0',
  [shop_user_id] BIGINT NOT NULL DEFAULT '0',
  [user_name] VARCHAR(45) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [currency] VARCHAR(45) NOT NULL DEFAULT '',
  [pay_type] VARCHAR(45) NOT NULL DEFAULT '',
  -- 支付状态（0：未支付，1：已支付）
  [pay_status] TINYINT NOT NULL DEFAULT '0',
  -- 支付时间
  [pay_time] INTEGER NOT NULL DEFAULT '0',
  -- 支付金额
  [pay_total] BIGINT NOT NULL DEFAULT '0',
  [mark] VARCHAR(45) NOT NULL DEFAULT '',
  [print_times] INTEGER NOT NULL DEFAULT '0',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_refund_sh] ON [mi_com_shops_refund] ([shop_id] ASC, [handler_id] ASC);
CREATE INDEX [idx_refund_ticket] ON [mi_com_shops_refund] ([ticket] ASC);
-- 挂单商品列表
CREATE TABLE [mi_com_shops_refund_items] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [refund_id] BIGINT NOT NULL,
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [prd_sn] VARCHAR(45) NOT NULL DEFAULT '',
  [prd_avatar] VARCHAR(255) NOT NULL DEFAULT '',
  [prd_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec] INTEGER NOT NULL DEFAULT '0',
  [weight] INTEGER NOT NULL DEFAULT '0',
  [style] TINYINT NOT NULL DEFAULT '1',
  [pack_name] VARCHAR(8) NOT NULL DEFAULT '箱',
  [style_type] TINYINT NOT NULL DEFAULT '0',
  [times] INTEGER NOT NULL DEFAULT '0',
  [debit] BIGINT NOT NULL DEFAULT '0',
  [off] BIGINT NOT NULL DEFAULT '0',
  [abatement] BIGINT NOT NULL DEFAULT '0',
  [coupons] BIGINT NOT NULL DEFAULT '0',
  [points] BIGINT NOT NULL DEFAULT '0',
  [balance] BIGINT NOT NULL DEFAULT '0',
  [price] BIGINT NOT NULL DEFAULT '0',
  [num] BIGINT NOT NULL DEFAULT '0',
  [total] BIGINT NOT NULL DEFAULT '0',
  [discount] BIGINT NOT NULL DEFAULT '0',
  [payed] BIGINT NOT NULL DEFAULT '0',
  [mark] VARCHAR(200) NOT NULL DEFAULT '',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_refund_id] ON [mi_com_shops_refund_items] ([refund_id] ASC);
CREATE INDEX [idx_refund_si] ON [mi_com_shops_refund_items] ([shop_id] ASC);
CREATE INDEX [idx_refund_handler_id] ON [mi_com_shops_refund_items] ([shop_id] ASC, [handler_id] ASC) -- 显示表单设置
CREATE TABLE [mi_com_setting_table] (
  [shop_id] BIGINT NOT NULL,
  [user_id] BIGINT NOT NULL,
  [table_id] VARCHAR(45) NOT NULL,
  [value] VARCHAR(1024) NOT NULL DEFAULT '',
  PRIMARY KEY ([shop_id], [user_id], [table_id])
);
-- 2025-07-10
CREATE TABLE [mi_com_shops_users_tags] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [name] VARCHAR(32) NOT NULL,
  [mark] VARCHAR(200) NOT NULL DEFAULT '',
  [sort] INTEGER NOT NULL DEFAULT '50',
  [status] INTEGER NOT NULL DEFAULT '1',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_users_levels] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL,
  [name] VARCHAR(45) NOT NULL,
  [mark] VARCHAR(64) NULL DEFAULT '',
  [sort] INTEGER NULL DEFAULT '50',
  [status] TINYINT NULL DEFAULT '1',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_users_tags_rel] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_user_id] BIGINT NOT NULL,
  [tag_id] VARCHAR(45) NOT NULL
);
CREATE INDEX [idx_sid] ON [mi_com_shops_users_tags_rel] ([shop_user_id] ASC);
CREATE TABLE [mi_com_shops_users] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [user_id] BIGINT NOT NULL DEFAULT '0',
  [point] BIGINT NOT NULL DEFAULT '0',
  [level] BIGINT NOT NULL DEFAULT '0',
  [tags] VARCHAR(64) NOT NULL DEFAULT '[]',
  [avatar] VARCHAR(200) NOT NULL DEFAULT '',
  [name] VARCHAR(45) NOT NULL DEFAULT '',
  [pinyin] VARCHAR(45) NOT NULL DEFAULT '',
  [nick_name] VARCHAR(45) NOT NULL DEFAULT '',
  [phone_number] VARCHAR(20) NOT NULL DEFAULT '',
  [country_code] VARCHAR(8) NOT NULL DEFAULT '',
  [gender] TINYINT NOT NULL DEFAULT '0',
  [birthday] INTEGER NOT NULL DEFAULT '0',
  [mark] VARCHAR(64) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [sort] INTEGER NOT NULL DEFAULT '50',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
-- 余额变化
CREATE TABLE [mi_com_shops_users_balance_log] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] BIGINT NOT NULL DEFAULT '0',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [user_id] BIGINT NOT NULL DEFAULT '0',
  [type] INTEGER NOT NULL DEFAULT '1',
  [tar_id] BIGINT NOT NULL DEFAULT '1',
  [amount] BIGINT NOT NULL DEFAULT '0',
  [amount_i] BIGINT NOT NULL DEFAULT '0',
  [amount_o] BIGINT NOT NULL DEFAULT '0',
  [value] INTEGER NOT NULL DEFAULT '0',
  [rec_time] INTEGER NOT NULL DEFAULT '0',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
-- 积分日志
CREATE TABLE [mi_com_shops_users_points_log] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] BIGINT NOT NULL DEFAULT '0',
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  [user_id] BIGINT NOT NULL DEFAULT '0',
  [type] INTEGER NOT NULL DEFAULT '1',
  [tar_id] BIGINT NOT NULL DEFAULT '1',
  [amount] BIGINT NOT NULL DEFAULT '0',
  [amount_i] BIGINT NOT NULL DEFAULT '0',
  [amount_o] BIGINT NOT NULL DEFAULT '0',
  [value] INTEGER NOT NULL DEFAULT '0',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_configs] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] INT NULL,
  [key] VARCHAR(45) NULL,
  [value] VARCHAR(250) NULL
);
-- token k-v
CREATE TABLE [mi_com_shops_token] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [track_id] VARCHAR(45) NOT NULL DEFAULT '',
  -- 店铺
  [shop_id] BIGINT NOT NULL DEFAULT '0',
  -- 哪个人
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  -- 次数
  [times] BIGINT NOT NULL DEFAULT '0',
  -- 
  [data] VARCHAR(4000) NOT NULL DEFAULT '',
  -- 过时时间
  [extime] INTEGER NOT NULL DEFAULT '0',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_token_track_id] ON [mi_com_shops_token] ([track_id] ASC);
-- mall_sp
CREATE TABLE [mi_com_shops_mall_sp] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] INTEGER NOT NULL DEFAULT '0',
  -- 票据
  [ticket] VARCHAR(45) NOT NULL,
  [track_id] BIGINT NOT NULL,
  [date_time] INTEGER NOT NULL DEFAULT '0',
  [supplier_id] BIGINT NOT NULL DEFAULT '0',
  [supplier_name] VARCHAR(45) NOT NULL DEFAULT '',
  [prd_nums] VARCHAR(64) NOT NULL DEFAULT '',
  [store_id] BIGINT NOT NULL DEFAULT '0',
  [store_name] VARCHAR(45) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [prd_mark] VARCHAR(255) NOT NULL DEFAULT '',
  [mark] VARCHAR(255) NOT NULL DEFAULT '',
  [price] BIGINT NOT NULL DEFAULT '0',
  [discount] BIGINT NOT NULL DEFAULT '0',
  [debt] BIGINT NOT NULL DEFAULT '0',
  [payed] BIGINT NOT NULL DEFAULT '0',
  [tax_fee] BIGINT NOT NULL DEFAULT '0',
  [tax_amount] BIGINT NOT NULL DEFAULT '0',
  [order_weight] BIGINT NOT NULL DEFAULT '0',
  [box_amount] BIGINT NOT NULL DEFAULT '0',
  [pack_amount] BIGINT NOT NULL DEFAULT '0',
  [rec_time] INTEGER NOT NULL DEFAULT '0',
  [rec_handler_id] BIGINT NOT NULL DEFAULT '0',
  [rec_handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [type] TINYINT NOT NULL DEFAULT '0',
  -- 打印次数
  [print_times] INTEGER NOT NULL DEFAULT '0',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_mall_sp_items] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [sp_id] BIGINT NOT NULL,
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [prd_sn] VARCHAR(45) NOT NULL DEFAULT '',
  [batch] VARCHAR(45) NOT NULL DEFAULT '',
  [avatar] VARCHAR(255) NOT NULL DEFAULT '',
  [prd_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec] INTEGER NOT NULL DEFAULT '0',
  [spec_weight] INTEGER NOT NULL DEFAULT '0',
  [pk_amount] INTEGER NOT NULL DEFAULT '0',
  [pk_weight] INTEGER NOT NULL DEFAULT '0',
  [unit_price] BIGINT NOT NULL DEFAULT '0',
  [pack_price] BIGINT NOT NULL DEFAULT '0',
  [times] INTEGER NOT NULL DEFAULT '0',
  [num] BIGINT NOT NULL DEFAULT '0',
  [price_total] BIGINT NOT NULL DEFAULT '0',
  [price_discount] BIGINT NOT NULL DEFAULT '0',
  [price_payed] BIGINT NOT NULL DEFAULT '0',
  [off] INTEGER NOT NULL DEFAULT '10000',
  [tax_unit_price] BIGINT NOT NULL DEFAULT '0',
  [tax_pack_price] BIGINT NOT NULL DEFAULT '0',
  [tax] INTEGER NOT NULL DEFAULT '0',
  [tax_fee] BIGINT NOT NULL DEFAULT '0',
  [tax_amount] BIGINT NOT NULL DEFAULT '0',
  [mark] VARCHAR(200) NOT NULL DEFAULT '',
  [unit] VARCHAR(2) NOT NULL DEFAULT 'g',
  [style] TINYINT NOT NULL DEFAULT '1',
  [pack_name] VARCHAR(8) NOT NULL DEFAULT '箱',
  [style_type] TINYINT NOT NULL DEFAULT '0',
  [sup_prd_id] BIGINT NOT NULL DEFAULT '0',
  [sup_prd_sn] VARCHAR(45) NOT NULL DEFAULT '',
  --过期时间
  [expire_time] INTEGER NOT NULL DEFAULT '0',
  --保质期（天）
  [keep_life] INTEGER NOT NULL DEFAULT '0',
  -- 保质期单位
  [keep_life_unit] VARCHAR(8) NOT NULL DEFAULT '',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_com_shops_mall_sp_items_sp_id] ON [mi_com_shops_mall_sp_items] ([sp_id] ASC);
-- 退货单
-- mall_rp
CREATE TABLE [mi_com_shops_mall_rp] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [shop_id] INTEGER NOT NULL DEFAULT '0',
  -- 票据
  [ticket] VARCHAR(45) NOT NULL,
  [track_id] BIGINT NOT NULL,
  [date_time] INTEGER NOT NULL DEFAULT '0',
  [supplier_id] BIGINT NOT NULL DEFAULT '0',
  [supplier_name] VARCHAR(45) NOT NULL DEFAULT '',
  [prd_nums] VARCHAR(64) NOT NULL DEFAULT '',
  [store_id] BIGINT NOT NULL DEFAULT '0',
  [store_name] VARCHAR(45) NOT NULL DEFAULT '',
  [handler_id] BIGINT NOT NULL DEFAULT '0',
  [handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [prd_mark] VARCHAR(255) NOT NULL DEFAULT '',
  [mark] VARCHAR(255) NOT NULL DEFAULT '',
  [price] BIGINT NOT NULL DEFAULT '0',
  [discount] BIGINT NOT NULL DEFAULT '0',
  [debt] BIGINT NOT NULL DEFAULT '0',
  [payed] BIGINT NOT NULL DEFAULT '0',
  [tax_fee] BIGINT NOT NULL DEFAULT '0',
  [tax_amount] BIGINT NOT NULL DEFAULT '0',
  [order_weight] BIGINT NOT NULL DEFAULT '0',
  [box_amount] BIGINT NOT NULL DEFAULT '0',
  [pack_amount] BIGINT NOT NULL DEFAULT '0',
  [rec_time] INTEGER NOT NULL DEFAULT '0',
  [rec_handler_id] BIGINT NOT NULL DEFAULT '0',
  [rec_handler_name] VARCHAR(45) NOT NULL DEFAULT '',
  [type] TINYINT NOT NULL DEFAULT '0',
  -- 打印次数
  [print_times] INTEGER NOT NULL DEFAULT '0',
  -- 最后更新时间
  [uptime] INTEGER NOT NULL DEFAULT '0',
  -- 称重状态 1 正常 2 异常
  [status] TINYINT NOT NULL DEFAULT '1',
  -- 确认时间
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE TABLE [mi_com_shops_mall_rp_items] (
  [id] INTEGER PRIMARY KEY AUTOINCREMENT,
  [rp_id] BIGINT NOT NULL,
  [prd_id] BIGINT NOT NULL DEFAULT '0',
  [prd_sn] VARCHAR(45) NOT NULL DEFAULT '',
  [batch] VARCHAR(45) NOT NULL DEFAULT '',
  [avatar] VARCHAR(255) NOT NULL DEFAULT '',
  [prd_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec_name] VARCHAR(45) NOT NULL DEFAULT '',
  [spec] INTEGER NOT NULL DEFAULT '0',
  [spec_weight] INTEGER NOT NULL DEFAULT '0',
  [pk_amount] INTEGER NOT NULL DEFAULT '0',
  [pk_weight] INTEGER NOT NULL DEFAULT '0',
  [unit_price] BIGINT NOT NULL DEFAULT '0',
  [pack_price] BIGINT NOT NULL DEFAULT '0',
  [times] INTEGER NOT NULL DEFAULT '0',
  [num] BIGINT NOT NULL DEFAULT '0',
  [price_total] BIGINT NOT NULL DEFAULT '0',
  [price_discount] BIGINT NOT NULL DEFAULT '0',
  [price_payed] BIGINT NOT NULL DEFAULT '0',
  [off] INTEGER NOT NULL DEFAULT '10000',
  [tax_unit_price] BIGINT NOT NULL DEFAULT '0',
  [tax_pack_price] BIGINT NOT NULL DEFAULT '0',
  [tax] INTEGER NOT NULL DEFAULT '0',
  [tax_fee] BIGINT NOT NULL DEFAULT '0',
  [tax_amount] BIGINT NOT NULL DEFAULT '0',
  [mark] VARCHAR(200) NOT NULL DEFAULT '',
  [unit] VARCHAR(2) NOT NULL DEFAULT 'g',
  [style] TINYINT NOT NULL DEFAULT '1',
  [pack_name] VARCHAR(8) NOT NULL DEFAULT '箱',
  [style_type] TINYINT NOT NULL DEFAULT '0',
  [sup_prd_id] BIGINT NOT NULL DEFAULT '0',
  [sup_prd_sn] VARCHAR(45) NOT NULL DEFAULT '',
  --过期时间
  [expire_time] INTEGER NOT NULL DEFAULT '0',
  --保质期（天）
  [keep_life] INTEGER NOT NULL DEFAULT '0',
  -- 保质期单位
  [keep_life_unit] VARCHAR(8) NOT NULL DEFAULT '',
  [intime] INTEGER NOT NULL DEFAULT '0'
);
CREATE INDEX [idx_com_shops_mall_rp_items_rp_id] ON [mi_com_shops_mall_rp_items] ([rp_id] ASC);