package v1

import (
	"wiki/internal/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type CaseReportCreateReq struct {
	g.Meta   `path:"/report/create" tags:"c-report" method:"post" summary:"create CaseReport"`
	Id       string      `json:"id" v:"required#CaseReport ID不能为空" dc:"id"`
	Message  string      `json:"message" dc:"报告内容" `                     //
	CaseId   int         `json:"caseId"   dc:"report关联caseid"`           //
	CaseDesc string      `json:"caseDesc" dc:"用例标题"`                     //
	RunTime  *gtime.Time `json:"runTime"  `                              //
	Status   int         `json:"status"   dc:"用例执行是否通过 0:未执行 1:通过 2:失败"` //
	RunUser  string      `json:"runUser"  `                              //
	RunIp    string      `json:"runIp"    `                              //
}

type CaseReportUpdateReq struct {
	g.Meta   `path:"/report/update" tags:"c-report" method:"post" summary:"update CaseReport"`
	Id       string `json:"id" v:"required#CaseReport ID不能为空" dc:"id"`
	Message  string `json:"message" dc:"报告内容" `                     //
	CaseId   int    `json:"caseId"   dc:"report关联caseid"`           //
	CaseDesc string `json:"caseDesc" dc:"用例标题"`                     //
	Status   int    `json:"status"   dc:"用例执行是否通过 0:未执行 1:通过 2:失败"` //
	RunUser  string `json:"runUser"  `                              //
	RunIp    string `json:"runIp"    `                              //
}

type CaseReportCreateRes struct {
	Ok bool `dc:"创建是否成功:true 成功 | false 失败" json:"ok"`
}

type CaseReportUpdateRes struct {
	Ok bool `dc:"更新是否成功:true 成功 | false 失败" json:"ok"`
}

type CaseReportFindByIdReq struct {
	g.Meta `path:"/report/get-one" tags:"c-report" method:"get" summary:"get one CaseReport by id"`
	Id     uint64 `v:"required#id不能是0或空" dc:"CaseReport ID"`
}

type CaseReportGetOneRes struct {
	model.CasesReportOutput
}

type CaseReportGetListReq struct {
	g.Meta  `path:"/report/list" tags:"c-report" method:"get" summary:"get list of CaseReport"`
	CaseId  int         `json:"caseId"`                               //
	Status  int         `json:"status" dc:"用例执行是否通过 0:未执行 1:通过 2:失败"` //
	RunTime *gtime.Time `json:runTime`
	Page    int         `dc:"分页号码" `                     // 分页号码
	Limit   int         `dc:"分页数量"`                      // 分页数量，最大50
	Sort    int         `dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"` // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type CaseReportGetListRes struct {
	Row   interface{} `dc:"row" json:"row"`
	Total int         `json:"total" dc:"总数"`
	Page  int         `json:"page" dc:"当前页码"`
	Limit int         `json:"limit" dc:"分页大小"`
}
