package store

import (
	"casher-server/internal/lager"
	"casher-server/internal/utils/def"
	"context"
	"fmt"
	"time"
)

/**
CREATE TABLE `mi_com_shops_auths` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `proxy_id` bigint(20) NOT NULL DEFAULT '0',
  `shop_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '行号Id',
  `open_id` varchar(45) NOT NULL DEFAULT '' COMMENT '店铺',
  `app_id` varchar(45) NOT NULL DEFAULT '' COMMENT '支付对应的APPID',
  `dev_id` varchar(45) NOT NULL DEFAULT '' COMMENT '设备编号',
  `mch_id` varchar(45) NOT NULL DEFAULT '' COMMENT '商号',
  `api_key` varchar(45) NOT NULL DEFAULT '' COMMENT 'API Key',
  `api_secret` varchar(64) NOT NULL DEFAULT '' COMMENT '通信加密',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '这里每家的积分与余额不同',
  `user_name` varchar(45) NOT NULL DEFAULT '',
  `handler_id` bigint(20) NOT NULL DEFAULT '0',
  `handler_name` varchar(45) NOT NULL DEFAULT '',
  `status` tinyint(3) NOT NULL DEFAULT '0' COMMENT '状态： 0正常 1禁用',
  `intime` int(11) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_ma` (`mch_id`,`api_key`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='商店登录授';

*/

type LoginReply struct {
	Req
	UserId int64 `json:"user_id"`
	RoomId int64 `json:"room_id"`
}

// 同一个请求，同步商品信息
type LoginRequest struct {
	Req
	MchId  string `json:"mch_id"`
	ApiKey string `json:"api_key"`
}

func (req *LoginRequest) Validate() error {
	if req.MchId == "" {
		return req.Tracker.Error("async_orders_validate", "请输入mch_id")
	}
	if req.ApiKey == "" {
		return req.Tracker.Error("async_orders_validate", "请输入api_key")
	}
	return nil
}

type AuthInfo struct {
	Id          int64  `json:"id"`
	ProxyId     int64  `json:"proxy_id"`
	ShopId      int64  `json:"shop_id"`
	OpenId      string `json:"open_id"`
	AppId       string `json:"app_id"`
	DevId       string `json:"dev_id"`
	MchId       string `json:"mch_id"`
	ApiKey      string `json:"api_key"`
	ApiSecret   string `json:"api_secret"`
	UserId      int64  `json:"user_id"`
	UserName    string `json:"user_name"`
	HandlerId   int64  `json:"handler_id"`
	HandlerName string `json:"handler_name"`
	Status      int32  `json:"status"`
	Intime      int64  `json:"intime"`
}

// GetUserIds proxyId, shopId, userId 获取用户ID
func (a *AuthInfo) GetUserIds() (int64, int64, int64) {
	return a.ProxyId, a.ShopId, a.UserId
}

func (s *Store) GetAuthInfo(ctx context.Context, req *LoginRequest) (*AuthInfo, error) {
	// 1 从缓存中获取
	authInfo, err := s.cache.GetOrLoadCtx(ctx, req.MchId+req.ApiKey, func(ctx context.Context) (any, time.Duration, error) {
		link := s.GetLink(ctx)
		auth, err := s.driver.GetAuthInfo(link, req.MchId, req.ApiKey)
		if err != nil {
			return nil, 0, err
		}
		// 校验返回签名
		err = CheckSign(req, auth.AppId)
		if err != nil {
			return nil, 0, err
		}
		return auth, time.Duration(def.GetNumber(s.profile.Cache.Expire, 5)) * time.Second, nil
	})
	if err != nil {
		return nil, err
	}
	return authInfo.(*AuthInfo), nil
}

// GetAuthInfoUseMA 从数据库中获取登录信息()
func (s *Store) GetAuthInfoUseMA(ctx context.Context, appId, apiKey, mchId, sign, norm string, ts int64) (*AuthInfo, error) {
	// 1 从缓存中获取
	authInfo, err := s.cache.GetOrLoadCtx(ctx, mchId+apiKey, func(ctx context.Context) (any, time.Duration, error) {
		link := s.GetLink(ctx)
		auth, err := s.driver.GetAuthInfo(link, mchId, apiKey)
		if err != nil {
			return nil, 0, err
		}

		// 校验返回签名
		err = CheckHeaderSign(sign, norm, auth.AppId, auth.ApiKey, auth.ApiSecret, ts)
		if err != nil {
			return nil, 0, err
		}
		return auth, time.Duration(def.GetNumber(s.profile.Cache.Expire, 5)) * time.Second, nil
	})
	if err != nil {
		return nil, err
	}
	return authInfo.(*AuthInfo), nil
}

func (s *Store) Login(ctx context.Context, req *LoginRequest, reply *LoginReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	fmt.Println("Login req = ", req)
	log.SetOperation("AsyncOrders", "AsyncOrders", "async")
	// 2 获取数据库连接
	// link := s.GetLink(ctx)
	// // 2.1 数据驱动
	// driver := s.GetDriver()
	// // 2.2 语言
	// lang := req.Tracker
	// 3 校验参数
	if err := req.Validate(); err != nil {
		return err
	}
	// 4 校验通过，返回用户信息
	return nil
}
