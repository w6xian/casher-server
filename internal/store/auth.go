package store

/*
 * mysql
CREATE TABLE `mi_cloud_companies_shops_auth` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `proxy_id` bigint(20) unsigned NOT NULL DEFAULT '0',
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
	ProxyId      int64  `json:"proxy_id"`
	AppId        string `json:"app_id"`
	AppSec       string `json:"app_sec"`
	ComId        int64  `json:"com_id"`
	ComName      string `json:"com_name"`
	AuthUserId   int64  `json:"auth_user_id"`
	AuthUserName string `json:"auth_user_name"`
	ExpireTime   int64  `json:"expire_time"`
	Status       int64  `json:"status"`
	Intime       int64  `json:"intime"`
}
