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
	ICaseStep interface {
		GetList(ctx context.Context, in model.CaseStepListInput) (out *model.CaseStepListOutput, err error)
		GetById(ctx context.Context, in uint64) (out *model.CaseStepOutput, err error)
		Create(ctx context.Context, in model.CaseStepCreateInput) (CaseStepId int64, err error)
		Update(ctx context.Context, in *entity.CaseStep) (err error)
		Delete(ctx context.Context, id int) (err error)
	}
)

var (
	localCaseStep ICaseStep
)

func CaseStep() ICaseStep {
	if localCaseStep == nil {
		panic("implement not found for interface ICaseStep, forgot register?")
	}
	return localCaseStep
}

func RegisterCaseStep(i ICaseStep) {
	localCaseStep = i
}
