package controller

import (
	"context"
	v1 "wiki/api/v1"
	"wiki/internal/model"
	"wiki/internal/model/entity"
	"wiki/internal/service"

	"github.com/gogf/gf/v2/os/gtime"
)

var (
	CaseReport = cCaseReport{}
)

type cCaseReport struct{}

func (c *cCaseReport) Create(ctx context.Context, req *v1.CaseReportCreateReq) (res *v1.CaseReportCreateRes, err error) {
	err = service.Report().Create(ctx, &entity.Casereport{
		Message:  req.Message,
		RunTime:  gtime.Now(),
		CaseId:   req.CaseId,
		CaseDesc: req.CaseDesc,
		Status:   req.Status,
		RunUser:  req.RunUser,
		RunIp:    req.RunIp,
		Id:       req.Id,
	})
	if err != nil {
		res = &v1.CaseReportCreateRes{
			Ok: false,
		}
	} else {
		res = &v1.CaseReportCreateRes{
			Ok: true,
		}
	}
	return res, nil
}

func (c *cCaseReport) Update(ctx context.Context, req *v1.CaseReportUpdateReq) (res *v1.CaseReportUpdateRes, err error) {
	err = service.Report().Update(ctx, &entity.Casereport{
		Message:  req.Message,
		RunTime:  gtime.Now(),
		CaseId:   req.CaseId,
		CaseDesc: req.CaseDesc,
		Status:   req.Status,
		RunUser:  req.RunUser,
		RunIp:    req.RunIp,
		Id:       req.Id,
	})
	r := &v1.CaseReportUpdateRes{Ok: true}
	if err != nil {
		r.Ok = false
	}
	return r, err
}

func (c *cCaseReport) FindById(ctx context.Context, req *v1.CaseReportFindByIdReq) (res *v1.CaseReportGetOneRes, err error) {
	out, err := service.Report().GetById(ctx, req.Id)
	res = &v1.CaseReportGetOneRes{}
	if err != nil {
		return res, err
	}

	res.CaseReport = out.CaseReport
	res.CasesReportOutput.CaseReport = out.CaseReport

	return res, nil
}

func (c *cCaseReport) FindList(ctx context.Context, req *v1.CaseReportGetListReq) (res *v1.CaseReportGetListRes, err error) {
	out, err := service.Report().GetList(ctx, model.CasesReportListInput{
		CaseId:  req.CaseId,
		Status:  req.Status,
		RunTime: req.RunTime,
		Page:    req.Page,
		Limit:   req.Limit,
		Sort:    req.Sort,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.CaseReportGetListRes{}
	res.Row = out.List
	res.Limit = out.Limit
	res.Page = out.Page
	res.Total = out.Total
	return res, nil
}
