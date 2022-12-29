// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"wiki/internal/model"
)

type (
	IValidatorI interface {
		GetValidator(k string) func(string, string) *model.ValidatorResult
	}
)

var (
	localValidatorI IValidatorI
)

func ValidatorI() IValidatorI {
	if localValidatorI == nil {
		panic("implement not found for interface IValidatorI, forgot register?")
	}
	return localValidatorI
}

func RegisterValidatorI(i IValidatorI) {
	localValidatorI = i
}