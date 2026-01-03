package store

import (
	"casher-server/internal/lager"
	"casher-server/internal/timex"
	"context"
)

/*
*
CREATE TABLE `mi_com_shops_users` (

	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`track_id` bigint(20) NOT NULL DEFAULT '0',
	`proxy_id` bigint(20) NOT NULL DEFAULT '0',
	`shop_id` bigint(20) NOT NULL DEFAULT '0',
	`user_id` bigint(20) NOT NULL DEFAULT '0',
	`point` bigint(20) NOT NULL DEFAULT '0' COMMENT '积分',
	`level` bigint(20) NOT NULL DEFAULT '0' COMMENT '会员等级',
	`tags` varchar(64) NOT NULL DEFAULT '[]' COMMENT '标签',
	`avatar` varchar(200) NOT NULL DEFAULT '' COMMENT '头像',
	`name` varchar(45) NOT NULL DEFAULT '' COMMENT '姓名，后台完善',
	`pinyin` varchar(45) NOT NULL DEFAULT '',
	`nick_name` varchar(45) NOT NULL DEFAULT '' COMMENT '昵称',
	`phone_number` varchar(20) NOT NULL DEFAULT '' COMMENT '电话',
	`country_code` varchar(8) NOT NULL DEFAULT '' COMMENT '国家代码',
	`gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '姓别1女，2男',
	`birthday` int(11) NOT NULL DEFAULT '0' COMMENT '生日',
	`mark` varchar(64) NOT NULL DEFAULT '' COMMENT '备注（后台可见）',
	`status` tinyint(4) NOT NULL DEFAULT '1',
	`intime` int(11) NOT NULL DEFAULT '0',
	PRIMARY KEY (`id`) USING BTREE,
	KEY `idx_user` (`user_id`)

) ENGINE=InnoDB AUTO_INCREMENT=88 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='shop-user表';
*/
type UserLite struct {
	Id      int64 `json:"id"`
	TrackId int64 `json:"track_id"`
	ProxyId int64 `json:"proxy_id"`
	ShopId  int64 `json:"shop_id"`
	UserId  int64 `json:"user_id"`

	Point int64  `json:"point"`
	Level int64  `json:"level"`
	Tags  string `json:"tags"`

	Avatar      string `json:"avatar"`
	Name        string `json:"name"`
	Pinyin      string `json:"pinyin"`
	NickName    string `json:"nick_name"`
	PhoneNumber string `json:"phone_number"`
	CountryCode string `json:"country_code"`
	Gender      int64  `json:"gender"`
	Birthday    int64  `json:"birthday"`

	Mark   string `json:"mark"`
	Status int64  `json:"status"`
	Intime int64  `json:"intime"`
}

// 返回同步用户信息
type AsyncUsersReply struct {
	Req
	// 同步用户信息
	Users []*UserLite `json:"users"`
	// 可同步用户数
	TotalNum int64 `json:"total_num"`
	LastTime int64 `json:"last_time"`
}

func (s *Store) AsyncUsers(ctx context.Context, req *AsyncRequest, reply *AsyncUsersReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("AsyncUsers", "AsyncUsers", "async")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 3 查询用户信息
	res, err := db.QueryUsers(link, req)
	if err != nil {
		log.ErrorExit("QueryUsers Query err", err)
		return lang.Error("msg_users_not_found", err.Error())
	}

	reply.AppId = req.AppId
	reply.Users = res.Users
	reply.TotalNum = res.TotalNum
	reply.LastTime = timex.UnixTime()
	return nil
}

/*
*
CREATE TABLE `mi_com_shops_users_levels` (

	`id` int(11) NOT NULL AUTO_INCREMENT,
	`proxy_id` int(11) NOT NULL DEFAULT '0',
	`name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
	`mark` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '' COMMENT '描述',
	`sort` int(11) DEFAULT '50' COMMENT '排序',
	`status` tinyint(1) DEFAULT '1' COMMENT '1有效0无效',
	`handler_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '操作人',
	`handler_name` varchar(45) NOT NULL DEFAULT '' COMMENT '职员名称',
	`update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
	`intime` int(11) NOT NULL COMMENT '入库时间',
	PRIMARY KEY (`id`),
	KEY `Idx_proxy` (`proxy_id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商城会员等级';
*/
type LevelLite struct {
	Id      int64 `json:"id"`
	ProxyId int64 `json:"proxy_id"`

	Name   string `json:"name"`
	Mark   string `json:"mark"`
	Sort   int64  `json:"sort"`
	Status int64  `json:"status"`

	HandlerId   int64  `json:"handler_id"`
	HandlerName string `json:"handler_name"`

	UpdateTime int64 `json:"update_time"`
	Intime     int64 `json:"intime"`
}

/*
*
CREATE TABLE `mi_com_shops_users_tags` (

	`id` bigint(20) NOT NULL AUTO_INCREMENT,
	`proxy_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '商户ID',
	`name` varchar(32) NOT NULL COMMENT '名称',
	`mark` varchar(200) NOT NULL DEFAULT '' COMMENT '备注信息',
	`sort` int(11) NOT NULL DEFAULT '50' COMMENT '序号，大靠前',
	`status` int(11) NOT NULL DEFAULT '1' COMMENT '状态 0-失效 1-有效',
	`handler_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '操作人',
	`handler_name` varchar(45) NOT NULL DEFAULT '' COMMENT '职员名称',
	`update_time` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
	`intime` int(11) NOT NULL COMMENT '入库时间',
	PRIMARY KEY (`id`),
	KEY `Idx_proxy` (`proxy_id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='商城会员tag,区别user_tag用于商城用户自定义，如地址中的家，还是公司等';
*/
type TagLite struct {
	Id      int64 `json:"id"`
	ProxyId int64 `json:"proxy_id"`

	Name   string `json:"name"`
	Mark   string `json:"mark"`
	Sort   int64  `json:"sort"`
	Status int64  `json:"status"`

	HandlerId   int64  `json:"handler_id"`
	HandlerName string `json:"handler_name"`

	UpdateTime int64 `json:"update_time"`
	Intime     int64 `json:"intime"`
}

type AsyncUsersExtraReply struct {
	Req
	// 同步用户等级信息
	Levels []*LevelLite `json:"levels"`
	// 同步用户tag信息
	Tags     []*TagLite `json:"tags"`
	LastTime int64      `json:"last_time"`
}

func (s *Store) AsyncUsersExtra(ctx context.Context, req *AsyncRequest, reply *AsyncUsersExtraReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("AsyncUsersExtra", "AsyncUsersExtra", "async")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 3 查询用户信息
	res, err := db.QueryUsersExtra(link, req)
	if err != nil {
		log.ErrorExit("QueryUsersExtra Query err", err)
		return lang.Error("msg_users_not_found", err.Error())
	}

	reply.AppId = req.AppId
	reply.Levels = res.Levels
	reply.Tags = res.Tags
	reply.LastTime = timex.UnixTime()
	return nil
}
