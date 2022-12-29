package controller

import (
	"context"
	"wiki/internal/model"
	"wiki/internal/model/entity"
	"wiki/internal/service"

	_ "github.com/gogf/gf/v2/frame/g"

	v1 "wiki/api/v1"
)

var (
	Element = cElement{}
)

type cElement struct{}

func (c *cElement) Create(ctx context.Context, req *v1.ElementCreateReq) (res *v1.ElementCreateRes, err error) {
	out, err := service.Element().Create(ctx, model.ElementCreateInput{
		PageId:      req.PageId,
		Path:        req.Path,
		Description: req.Description,
		Attribute:   req.Attribute,
		Father:      req.Father,
		Name:        req.Name},
	)
	if err != nil {
		return nil, err
	}
	res = &v1.ElementCreateRes{
		PageId: out,
	}
	return res, nil
}

func (c *cElement) Update(ctx context.Context, req *v1.ElementUpdateReq) (res *v1.ElementUpdateRes, err error) {
	err = service.Element().Update(ctx, &entity.Element{
		Description: req.Description,
		Id:          req.Id,
		Name:        req.Name,
		PageId:      req.PageId,
		Father:      req.Father,
		Path:        req.Path,
		Attribute:   req.Attribute,
	})
	r := &v1.ElementUpdateRes{Ok: true}
	if err != nil {
		r.Ok = false
	}
	return r, err
}

func (c *cElement) Remove(ctx context.Context, req *v1.ElementRemoveReq) (res *v1.ElementUpdateRes, err error) {
	err = service.Element().Delete(ctx, req.Id)
	r := &v1.ElementUpdateRes{Ok: true}
	if err != nil {
		r.Ok = false
	}
	return r, err
}

func (c *cElement) FindById(ctx context.Context, req *v1.ElementFindByIdReq) (res *v1.ElementGetOneRes, err error) {
	out, err := service.Element().GetById(ctx, req.Id)
	res = &v1.ElementGetOneRes{}
	if err != nil {
		return res, err
	}

	res.Element = out.Element
	res.Page = out.Page
	res.FatherElement = out.FatherElement

	return res, nil
}

func (c *cElement) FindList(ctx context.Context, req *v1.ElementGetListReq) (res *v1.ElementGetListRes, err error) {
	out, err := service.Element().GetList(ctx, model.ElementListInput{
		Page:        req.Page,
		Description: req.Description,
		ElementName: req.ElementName,
		PageId:      req.PageId,
		Father:      req.Father,
		Limit:       req.Limit,
		Sort:        req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.ElementGetListRes{}
	res.Row = out.List
	res.Limit = out.Limit
	res.Page = out.Page
	res.Total = out.Total

	return res, nil
}
