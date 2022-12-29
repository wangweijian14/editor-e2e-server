package model

import "wiki/internal/model/entity"

type StepCreateInput struct {
	Description     string `json:"description"     ` //
	TargetElementId int    `json:"target_element_id"` //
	ActionId        int    `json:"action_id"      ` //
	TargetPage      int    `json:"target_page"      ` //
}

type StepOutput struct {
	Step    * entity.Step	`json:"step"`
	Page    * entity.Page	`json:"page"`
	Element * ElementOutput	`json:"element"`
}

// StepListInput 获取 step 列表
type StepListInput struct {
	Description     string `json:"description"     ` //
	TargetElementId int    `json:"target_element_id"`
	ActionId        int    `json:"action_id"`
	TargetPage      int    `json:"target_page"`
	Ids string	`json:"ids"`
	Page       int    // 分页号码
	Limit       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// StepListOutput 查询  step  结果
type StepListOutput struct {
	List  []*entity.Step `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Limit  int                        `json:"limit" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}