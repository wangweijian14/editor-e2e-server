package model

import "wiki/internal/model/entity"

// CasesListInput 获取 Cases 列表
type CasesListInput struct {
	HisId       int    `json:"hisId"`        //
	Description string `json:"description" ` //
	Page        int    // 分页号码
	Limit       int    // 分页数量，最大50
	Sort        int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CasesListOutput 查询Cases结果
type CasesListOutput struct {
	List  []*entity.Cases `json:"list" description:"列表"`
	Page  int             `json:"page" description:"分页码"`
	Limit int             `json:"limit" description:"分页数量"`
	Total int             `json:"total" description:"数据总数"`
}

type CasesOutput struct {
	Cases          *entity.Cases     `json:"cases"`
	CaseStepOutput []*CaseStepOutput `json:"case_step_output"`
}

type CasesCreateInput struct {
	HisId       string `json:"hisId"`       //
	Description string `json:"description"` //
	ExecCount   int    `json:"execCount" `  //
	PassCount   int    `json:"passCount"`   //
	SkipCount   int    `json:"skipCount"`   //
	FailedCount int    `json:"failedCount"` //
	OpenId      int    `json:"openId"`      //
}
