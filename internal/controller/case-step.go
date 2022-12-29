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
	CaseStep = cCaseStep{}
)

type cCaseStep struct{}

func (c *cCaseStep) Create(ctx context.Context, req *v1.CaseStepCreateReq) (res *v1.CaseStepCreateRes, err error) {
	out, err := service.CaseStep().Create(ctx, model.CaseStepCreateInput{
		StepId:       req.StepId,
		Input:        req.Input,
		AssertExpect: req.AssertExpect,
		Description:  req.Description,
		Score:        req.Score,
		CaseId:       req.CaseId,
	},
	)
	if err != nil {
		return nil, err
	}

	res = &v1.CaseStepCreateRes{
		PageId: out,
	}
	return res, nil
}

func (c *cCaseStep) Update(ctx context.Context, req *v1.CaseStepUpdateReq) (res *v1.CaseStepUpdateRes, err error) {
	err = service.CaseStep().Update(ctx, &entity.CaseStep{
		StepId:       req.StepId,
		Input:        req.Input,
		AssertExpect: req.AssertExpect,
		Description:  req.Description,
		Id:           req.Id,
		Score:        req.Score,
		CaseId:       req.CaseId,
	})
	r := &v1.CaseStepUpdateRes{Ok: true}
	if err != nil {
		r.Ok = false
	}
	return r, err
}

func (c *cCaseStep) Remove(ctx context.Context, req *v1.CaseStepRemoveReq) (res *v1.CaseStepUpdateRes, err error) {
	err = service.CaseStep().Delete(ctx, req.Id)
	r := &v1.CaseStepUpdateRes{Ok: true}
	if err != nil {
		r.Ok = false
	}
	return r, err
}

func (c *cCaseStep) FindById(ctx context.Context, req *v1.CaseStepFindByIdReq) (res *v1.CaseStepGetOneRes, err error) {
	out, err := service.CaseStep().GetById(ctx, req.Id)
	res = &v1.CaseStepGetOneRes{}
	if err != nil {
		return res, err
	}

	res.CaseStep = out.CaseStep
	res.Step = out.Step

	return res, nil
}

func (c *cCaseStep) FindList(ctx context.Context, req *v1.CaseStepGetListReq) (res *v1.CaseStepGetListRes, err error) {
	out, err := service.CaseStep().GetList(ctx, model.CaseStepListInput{
		Page:         req.Page,
		StepId:       req.StepId,
		Input:        req.Input,
		AssertExpect: req.AssertExpect,
		CaseId:       req.CaseId,
		Limit:        req.Limit,
		Sort:         req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.CaseStepGetListRes{}
	res.Row = out.List
	res.Limit = out.Limit
	res.Page = out.Page
	res.Total = out.Total
	return res, nil
}
