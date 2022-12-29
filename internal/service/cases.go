// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"context"
	"wiki/internal/model"
	"wiki/internal/model/entity"
)

type ICases interface {
	GetList(ctx context.Context, in model.CasesListInput) (out *model.CasesListOutput, err error)
	GetById(ctx context.Context, in uint64) (out *model.CasesOutput, err error)
	Create(ctx context.Context, in model.CasesCreateInput) (CasesId int64, err error)
	Update(ctx context.Context, in *entity.Cases) (err error)
}

var localCases ICases

func Cases() ICases {
	if localCases == nil {
		panic("implement not found for interface ICases, forgot register?")
	}
	return localCases
}

func RegisterCases(i ICases) {
	localCases = i
}
