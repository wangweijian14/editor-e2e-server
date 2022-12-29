package model

import "wiki/internal/model/entity"

// CaseStepListInput 获取 CaseStep 列表
type CaseStepListInput struct {
	StepId       int    `json:"stepId"       ` //
	Input        string `json:"input"        ` //
	AssertExpect string `json:"assertExpect" ` //
	CaseId       int    `json:"caseId" `       //
	Page         int    // 分页号码
	Limit        int    // 分页数量，最大50
	Sort         int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CaseStepListOutput 查询CaseStep结果
type CaseStepListOutput struct {
	List  []*entity.CaseStep `json:"list" description:"列表"`
	Page  int                `json:"page" description:"分页码"`
	Limit int                `json:"limit" description:"分页数量"`
	Total int                `json:"total" description:"数据总数"`
}

type CaseStepOutput struct {
	CaseStep *entity.CaseStep `json:"case_step"`
	Step     *StepOutput      `json:"step"`
}

type CaseStepCreateInput struct {
	StepId       int    `json:"stepId"       `       //
	Input        string `json:"input"        `       //
	Description  string `json:"description"        ` //
	AssertExpect string `json:"assertExpect" `       //
	Score        int    `json:"score"   `            //
	CaseId       int    `json:"caseId" `             //
}
