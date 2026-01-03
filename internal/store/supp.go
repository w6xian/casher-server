package store

import (
	"casher-server/internal/lager"
	"casher-server/internal/timex"
	"casher-server/internal/utils"
	"casher-server/internal/utils/id"
	"context"
	"fmt"
)

type SuppCodeRequest struct {
	Code string `json:"code"`
	// 本地的appId，用于以后推送订单
	AppId   string `json:"app_id"`
	AppCode string `json:"app_code"`
	// 服务器的openId
	OpenId string `json:"open_id"`
	// GEO信息
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Sign      string `json:"sign"`
}

type SuppCodeReply struct {
	AppId         string `json:"com_api"`
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
	AuthId        int64  `json:"auth_id"`
	AuthAppId     string `json:"auth_app_id"`
	AuthAppSec    string `json:"auth_app_sec"`
	Sign          string `json:"sign"`
}

func (s *Store) ReadProductSuppUseCode(ctx context.Context, req *SuppCodeRequest, reply *SuppCodeReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("ReadProductSuppUseCode", "ReadProductSuppUseCode", "Supp")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	fmt.Println("ReadProductSuppUseCode link=", link)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 3 查询公司信息
	comInfo, err := db.GetCRMCompanyByOpenId(link, req.Code)
	if err != nil {
		log.ErrorExit("ReadProductSuppUseCode Query err", err)
		return err
	}
	// 注册信息
	proxy, err := db.GetProxyInfoById(link, comInfo.ProxyId)
	if err != nil {
		log.ErrorExit("ReadProductSuppUseCode Query err", err)
		return err
	}

	// 授权信息返回
	reply.Name = proxy.Name
	reply.SortName = proxy.SortName
	reply.SnCode = proxy.SnCode
	reply.LegalPersion = proxy.LegalPersion
	reply.ContactName = proxy.ContactName
	reply.ContactPhone = utils.IsEmptyUseDefault(proxy.ContactPhone, proxy.ContactMobile)
	reply.ContactMobile = utils.IsEmptyUseDefault(proxy.ContactMobile, proxy.ContactPhone)
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

	// 4 查询授权信息

	auths, err := db.GetAuthByComId(link, comInfo.Id)
	if err == nil {
		reply.AuthId = auths.Id
		reply.AuthAppId = auths.AppId
		reply.AuthAppSec = auths.AppSec
		return nil
	}
	appId := id.ShortID()
	appSec := id.RandStr(64)
	authId, err := db.InsertAuth(link, &Auths{
		ProxyId:      comInfo.ProxyId,
		AppId:        appId,
		AppSec:       appSec,
		ComId:        comInfo.Id,
		ComName:      proxy.Name,
		AuthUserId:   0,
		AuthUserName: "",
		ExpireTime:   0,
		Status:       1,
		Intime:       timex.UnixTime(),
	})
	if err != nil {
		log.ErrorExit("ReadProductSuppUseCode Insert err", err)
		return err
	}
	// authId
	reply.AuthId = authId
	reply.AuthAppId = appId
	reply.AuthAppSec = appSec
	// reply写一个方法，传入结构按Tag字母升序排序，值转换为字符串后，按Tag:Value格式拼接，多个属性之间用分号分隔，算出md5再与key再md5得到sign
	reply.Sign = utils.CalcSign(reply, comInfo.ComPem)
	return nil
}
