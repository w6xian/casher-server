package store

import (
	"casher-server/internal/lager"
	"context"
)

func (s *Store) GetPublicProductBySn(ctx context.Context, req *PrdSnReq, reply *PublicProductReply) error {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("GetPublicProductBySn", "GetPublicProductBySn", "Supp")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	lang := req.Tracker
	// 3 查询公司信息
	productModel, err := db.GetPublicProductBySn(link, req.Sn)
	if err != nil {
		log.ErrorExit("GetPublicProductBySn Query err", err)
		return lang.Error("msg_public_product_not_found", err.Error())
	}
	reply.AppId = req.AppId
	reply.Sn = productModel.Sn
	reply.Name = productModel.Name
	reply.UnionId = productModel.UnionId
	reply.Avatar = productModel.Avatar
	reply.Cover = productModel.Cover
	reply.Pinyin = productModel.Pinyin
	reply.BrandName = productModel.BrandName
	reply.Feature = productModel.Feature
	reply.Price = productModel.Price
	reply.Spec = productModel.Spec
	reply.SpecName = productModel.SpecName
	reply.SpecWeight = productModel.SpecWeight
	reply.PkAmount = productModel.PkAmount
	reply.PkWeight = productModel.PkWeight
	reply.PackName = productModel.PackName
	reply.KeepLife = productModel.KeepLife
	reply.KeepLifeUnit = productModel.KeepLifeUnit
	reply.StorageConditions = productModel.StorageConditions
	reply.TransportationConditions = productModel.TransportationConditions
	reply.Unit = productModel.Unit
	reply.Habitat = productModel.Habitat
	reply.Style = productModel.Style
	reply.StyleType = productModel.StyleType
	reply.Units = productModel.Units
	reply.Status = productModel.Status
	return nil
}
func (s *Store) GetPublicProductBySnV2(ctx context.Context, tracker *lager.Tracker, sn string) (*ProductModel, error) {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("GetPublicProductBySn", "GetPublicProductBySn", "Supp")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	// 2.2 语言
	// 3 查询公司信息
	productModel, err := db.GetPublicProductBySn(link, sn)
	if err != nil {
		log.ErrorExit("GetPublicProductBySn Query err", err)
		return nil, tracker.Error("msg_public_product_not_found", err.Error())
	}
	return productModel, nil
}
func (s *Store) ReplacePublicProduct(ctx context.Context, tracker *lager.Tracker, prd *ProductModel) (int64, error) {
	// 1 日志
	log := lager.FromContext(ctx)
	defer log.Sync()
	log.SetOperation("ReplacePublicProduct", "ReplacePublicProduct", "Supp")
	// 2 获取数据库连接
	link := s.GetLink(ctx)
	// 2.1 数据驱动
	db := s.GetDriver()
	if prd.PackName == "" {
		if prd.PkAmount == 1 {
			prd.PackName = prd.SpecName
		}
	}
	productModel, err := db.GetPublicProductBySn(link, prd.Sn)
	if err == nil {
		// 有记录的情况下q，更新版本
		if productModel.Equal(prd) {
			return 1, nil
		}
		id, err := db.InsertPublicProductVersion(link, productModel)
		if err != nil {
			log.ErrorExit("ReplacePublicProduct Insert err", err)
			return 0, tracker.Error("msg_public_product_not_found", err.Error())
		}
		return id, nil
	}
	// 2.2 语言
	// 3 查询公司信息
	id, err := db.InsertPublicProduct(link, prd)
	if err != nil {
		log.ErrorExit("ReplacePublicProduct Insert err", err)
		return 0, tracker.Error("msg_public_product_not_found", err.Error())
	}
	return id, nil
}
