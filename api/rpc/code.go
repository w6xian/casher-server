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
	Pinyin        string `json:"pinyin"`
	Image         string `json:"image"`
	Area          string `json:"area"`
	AreaCode      string `json:"area_code"`
	AreaPath      string `json:"area_path"`
	AreaStreet    string `json:"area_street"`
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
	row.Scan(reply)
	fmt.Println("ReadProductSuppUseCode reply=", reply)
	return nil
}
