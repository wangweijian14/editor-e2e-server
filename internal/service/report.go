// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"wiki/internal/model"
	"wiki/internal/model/entity"
)

type (
	IReport interface {
		GetList(ctx context.Context, in model.CasesReportListInput) (out *model.CasesReportListOutput, err error)
		GetById(ctx context.Context, in uint64) (out *model.CasesReportOutput, err error)
		Create(ctx context.Context, in *entity.Casereport) (err error)
		Update(ctx context.Context, in *entity.Casereport) (err error)
	}
)

var (
	localReport IReport
)

func Report() IReport {
	if localReport == nil {
		panic("implement not found for interface IReport, forgot register?")
	}
	return localReport
}

func RegisterReport(i IReport) {
	localReport = i
}
