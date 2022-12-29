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
	Step = cStep{}
)

type cStep struct{}

func (c *cStep) Create(ctx context.Context, req *v1.StepCreateReq) (res *v1.StepCreateRes, err error) {
	out, err := service.Step().Create(ctx, model.StepCreateInput{
		Description:     req.Description,
		TargetPage:      req.TargetPage,
		TargetElementId: req.TargetElementId,
		ActionId:        req.ActionId,
	},
	)
	if err != nil {
		return nil, err
	}
	res = &v1.StepCreateRes{
		Id: out,
	}
	return res, nil
}

func (c *cStep) Update(ctx context.Context, req *v1.StepUpdateReq) (res *v1.StepUpdateRes, err error) {
	err = service.Step().Update(ctx, &entity.Step{
		Description:     req.Description,
		Id:              req.Id,
		TargetPage:      req.TargetPage,
		TargetElementId: req.TargetElementId,
		ActionId:        req.ActionId,
	})
	r := &v1.StepUpdateRes{Ok: true}
	if err != nil {
		r.Ok = false
	}
	return r, err
}

func (c *cStep) FindById(ctx context.Context, req *v1.StepFindByIdReq) (res *v1.StepGetOneRes, err error) {
	out, err := service.Step().GetById(ctx, req.Id)
	res = &v1.StepGetOneRes{}
	if err != nil {
		return res, err
	}

	res.Element = out.Element.Element
	res.Page = out.Page
	res.Step = out.Step

	return res, nil
}

func (c *cStep) FindList(ctx context.Context, req *v1.StepGetListReq) (res *v1.StepGetListRes, err error) {
	out, err := service.Step().GetList(ctx, model.StepListInput{
		Page:            req.Page,
		Description:     req.Description,
		ActionId:        req.ActionId,
		TargetPage:      req.TargetPage,
		TargetElementId: req.TargetElementId,
		Ids:             req.Ids,
		Limit:           req.Limit,
		Sort:            req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.StepGetListRes{}
	res.Row = out.List
	res.Limit = out.Limit
	res.Page = out.Page
	res.Total = out.Total
	return res, nil
}

func (c *cStep) Remove(ctx context.Context, req *v1.StepRemoveReq) (res *v1.StepUpdateRes, err error) {
	err = service.Step().Delete(ctx, req.Id)
	r := &v1.StepUpdateRes{Ok: true}
	if err != nil {
		r.Ok = false
	}
	return r, err
}
