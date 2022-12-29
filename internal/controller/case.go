package controller

import (
	"context"
	"fmt"
	v1 "wiki/api/v1"
	"wiki/internal/model"
	"wiki/internal/model/entity"
	"wiki/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

var (
	Cases = cCases{}
)

type cCases struct{}

func (c *cCases) Create(ctx context.Context, req *v1.CasesCreateReq) (res *v1.CasesCreateRes, err error) {
	out, err := service.Cases().Create(ctx, model.CasesCreateInput{
		Description: req.Description,
		HisId:       fmt.Sprint(gtime.Now().Timestamp()),
		ExecCount:   0,
		PassCount:   0,
		FailedCount: 0,
		SkipCount:   0,
	},
	)
	if err != nil {
		return nil, err
	}
	res = &v1.CasesCreateRes{
		CaseId: out,
	}
	return res, nil
}

func (c *cCases) Update(ctx context.Context, req *v1.CasesUpdateReq) (res *v1.CasesUpdateRes, err error) {
	err = service.Cases().Update(ctx, &entity.Cases{
		Description: req.Description,
		Id:          req.Id,
		OpenId:      req.OpenId,
	})
	r := &v1.CasesUpdateRes{Ok: true}
	if err != nil {
		r.Ok = false
	}
	return r, err
}

func (c *cCases) FindById(ctx context.Context, req *v1.CasesFindByIdReq) (res *v1.CasesGetOneRes, err error) {
	out, err := service.Cases().GetById(ctx, req.Id)
	res = &v1.CasesGetOneRes{}
	if err != nil {
		return res, err
	}

	res.Cases = out.Cases
	res.CasesOutput.CaseStepOutput = out.CaseStepOutput

	return res, nil
}

func (c *cCases) FindList(ctx context.Context, req *v1.CasesGetListReq) (res *v1.CasesGetListRes, err error) {
	out, err := service.Cases().GetList(ctx, model.CasesListInput{
		Page:        req.Page,
		Description: req.Description,
		Limit:       req.Limit,
		Sort:        req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.CasesGetListRes{}
	res.Row = out.List
	res.Limit = out.Limit
	res.Page = out.Page
	res.Total = out.Total
	return res, nil
}
