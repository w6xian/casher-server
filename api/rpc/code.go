package rpc

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type SuppCodeRequest struct {
	Code       string `json:"code"`
	ShopCode   string `json:"shop_code"`
	CreateTime string `json:"createTime"`
	AuthToken  string `json:"authToken"` //仅tcp时使用，发送msg时带上
}

type SuppCodeReply struct {
	Code          string `json:"code"`
	Name          string `json:"name"`
	SortName      string `json:"sort_name"`
	SnCode        string `json:"sn_code"`
	Avatar        string `json:"avatar"`
	Pinyin        string `json:"pinyin"`
	Image         string `json:"image"`
	AreaCode      string `json:"area_code"`
	AreaPath      string `json:"area_path"`
	AreaStreet    string `json:"area_street"`
	Street        string `json:"street"`
	Address       string `json:"address"`
	LegalPersion  string `json:"legal_persion"`
	ContactName   string `json:"contact_name"`
	ContactPhone  string `json:"contact_phone"`
	ContactMobile string `json:"contact_mobile"`
	ContactEmail  string `json:"contact_email"`
	RegisterDate  string `json:"register_date"`
	ExpireDate    string `json:"expire_date"`
	CreateTime    string `json:"create_time"`
	AuthToken     string `json:"auth_token"`
}
type ProxyInfo struct {
	ProxyId      int64  `json:"com_id"`
	Avatar       string `json:"avatar"`
	Name         string `json:"name"`
	SortName     string `json:"sort_name"`
	SnCode       string `json:"sn_code"`
	LegalPersion string `json:"legal_persion"`
	ContactName  string `json:"contact_name"`
	ContactPhone string `json:"contact_phone"`
	ContactEmail string `json:"contact_email"`
	RegisterDate string `json:"register_date"`
	ExpireDate   string `json:"expire_date"`
	AreaCode     string `json:"area_code"`
	AreaPath     string `json:"area_path"`
	AreaStreet   string `json:"area_street"`
	Street       string `json:"street"`
	Address      string `json:"address"`
	CreateTime   string `json:"intime"`
}

/*
 * mysql
CREATE TABLE `mi_cloud_companies_shops_auth` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` varchar(45) NOT NULL COMMENT '应用ID号',
  `app_sec` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密匙',
  `com_id` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '0' COMMENT '被授权方',
  `com_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '授权公司名称',
  `auth_user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '同意人Id',
  `auth_user_name` varchar(45) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '同意人名称',
  `expire_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '过期时间',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '1有效0无效',
  `intime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '入库时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `idx_app_id` (`app_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='超市客户授权';
*/

type Auths struct {
	Id           int64  `json:"id"`
	AppId        string `json:"app_id"`
	AppSec       string `json:"app_sec"`
	ComId        string `json:"com_id"`
	ComName      string `json:"com_name"`
	AuthUserId   int64  `json:"auth_user_id"`
	AuthUserName string `json:"auth_user_name"`
	ExpireTime   int64  `json:"expire_time"`
	Status       int64  `json:"status"`
	Intime       int64  `json:"intime"`
}

func (c *Order) ReadProductSuppUseCode(ctx context.Context, req *SuppCodeRequest, reply *SuppCodeReply) error {
	fmt.Println("ReadProductSuppUseCode=", req.Code)
	ctx, close := c.Store.DbConnectWithClose(ctx)
	defer close()
	link := c.Store.GetLink(ctx)
	fmt.Println("ReadProductSuppUseCode link=", link)
	row, err := link.Table("crm_companies").Where("open_id = '%s'", req.Code).Query()
	if err != nil {
		c.Lager.Error("ReadProductSuppUseCode Query err", zap.Error(err))
		return err
	}
	proxyId, err := row.Get("proxy_id").Int64()
	if err != nil {
		c.Lager.Error("ReadProductSuppUseCode Query err", zap.Error(err))
		return err
	}
	com, err := link.Table("cloud_companies").Where("com_id = %d", proxyId).Query()
	if err != nil {
		c.Lager.Error("ReadProductSuppUseCode Query err", zap.Error(err))
		return err
	}

	proxy := &ProxyInfo{}
	err = com.Scan(proxy)
	if err != nil {
		c.Lager.Error("ReadProductSuppUseCode Scan err", zap.Error(err))
		return err
	}
	reply.Name = proxy.Name
	reply.SortName = proxy.SortName
	reply.SnCode = proxy.SnCode
	reply.LegalPersion = proxy.LegalPersion
	reply.ContactName = proxy.ContactName
	reply.ContactPhone = proxy.ContactPhone
	reply.ContactEmail = proxy.ContactEmail
	reply.RegisterDate = proxy.RegisterDate
	reply.ExpireDate = proxy.ExpireDate
	reply.AreaCode = proxy.AreaCode
	reply.AreaPath = proxy.AreaPath
	reply.AreaStreet = proxy.AreaStreet
	reply.Street = proxy.Street
	reply.Address = proxy.Address
	reply.CreateTime = proxy.CreateTime
	reply.Avatar = proxy.Avatar

	fmt.Println("ReadProductSuppUseCode reply=", reply)
	return nil
}
