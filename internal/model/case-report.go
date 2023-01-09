package model

import (
	"wiki/internal/model/entity"

	"github.com/gogf/gf/v2/os/gtime"
)

// CasesListInput 获取 Cases 列表
type CasesReportListInput struct {
	CaseId  int         `json:"caseId"`  //
	Status  int         `json:"status" ` //
	RunTime *gtime.Time `json:runTime`
	Page    int         // 分页号码
	Limit   int         // 分页数量，最大50
	Sort    int         // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// CasesListOutput 查询Cases结果
type CasesReportListOutput struct {
	List  []*entity.Casereport `json:"list" description:"列表"`
	Page  int                  `json:"page" description:"分页码"`
	Limit int                  `json:"limit" description:"分页数量"`
	Total int                  `json:"total" description:"数据总数"`
}

type CasesReportOutput struct {
	CaseReport *entity.Casereport `json:"caseReport"`
}

type CasesReportCreateInput struct {
	CaseReport *entity.Casereport `json:"caseReport"`
}
