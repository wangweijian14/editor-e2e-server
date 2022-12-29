package v1

import (
	"wiki/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CaseStepCreateReq struct {
	g.Meta       `path:"/case-step/create" tags:"caseStep" method:"post" summary:"create CaseStep"`
	StepId       int    `json:"stepId"  v:"required#请输步骤id" dc:"对应的步骤ID" `                          //
	Input        string `json:"input"    dc:"对应的输入内容，input会自动检查Focus、WaitEnabled、WaitWritable"    ` //
	Description  string `json:"description"    dc:"步骤描述"    `                                       //
	Score        int    `json:"score"   dc:"score越大优先级越高" `                                         //
	CaseId       int    `json:"caseId"   dc:"关联的用例ID" `
	AssertExpect string `json:"assertExpect" dc:"检查点，计划是json 格式，做的时候在定义"` //
}

type CaseStepUpdateReq struct {
	g.Meta       `path:"/case-step/update" tags:"caseStep" method:"post" summary:"update CaseStep"`
	Id           int    `json:"id" v:"required#CaseStep ID不能为空" dc:"id"`
	StepId       int    `json:"stepId"   dc:"对应的步骤ID" `       //
	Description  string `json:"description"    dc:"步骤描述"    ` //
	Score        int    `json:"score"   dc:"score越大优先级越高" `   //
	CaseId       int    `json:"caseId"   dc:"关联的用例ID" `
	Input        string `json:"input"    dc:"对应的输入内容，input会自动检查Focus、WaitEnabled、WaitWritable"    ` //
	AssertExpect string `json:"assertExpect" dc:"检查点，计划是json 格式，做的时候在定义"`                           //
}

type CaseStepRemoveReq struct {
	g.Meta `path:"/case-step/remove" tags:"caseStep" method:"post" summary:"remove CaseStep"`
	Id     int `json:"id" v:"required#CaseStep ID不能为空" dc:"id"`
}

type CaseStepCreateRes struct {
	PageId int64 `dc:"create CaseStep id"`
}

type CaseStepUpdateRes struct {
	Ok bool `dc:"更新是否成功：true 成功 | false 失败" json:"ok"`
}

type CaseStepFindByIdReq struct {
	g.Meta `path:"/case-step/get-one" tags:"caseStep" method:"post" summary:"get one CaseStep by id"`
	Id     uint64 `v:"required#id不能是0或空" dc:"CaseStep ID"`
}

type CaseStepGetOneRes struct {
	model.CaseStepOutput
}

type CaseStepGetListReq struct {
	g.Meta       `path:"/case-step/list" tags:"caseStep" method:"get" summary:"get list of CaseStep"`
	StepId       int    `json:"stepId"   dc:"对应的步骤ID" `                                             //
	Input        string `json:"input"    dc:"对应的输入内容，input会自动检查Focus、WaitEnabled、WaitWritable"    ` //
	AssertExpect string `json:"assertExpect" dc:"检查点，计划是json 格式，做的时候在定义"`                           //
	CaseId       int    `json:"caseId"   dc:"关联的用例ID" `
	Page         int    `dc:"分页号码" `                                         // 分页号码
	Limit        int    `dc:"分页数量"`                                          // 分页数量，最大50
	Sort         int    `json:"sort" dc:"排序类型(0:最新, 默认。1:活跃, 2:热度 3 score)"` // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type CaseStepGetListRes struct {
	Row   interface{} `dc:"row" json:"row"`
	Total int         `json:"total" dc:"总数"`
	Page  int         `json:"page" dc:"当前页码"`
	Limit int         `json:"limit" dc:"分页大小"`
}
