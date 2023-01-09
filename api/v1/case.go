package v1

import (
	"wiki/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type CasesExecuteReq struct {
	g.Meta   `path:"/case/execute" tags:"case" method:"post" summary:"case execute"`
	Id       string `json:"case_id" dc:"用例iD""`
	MustPage string `json:"must_page" dc:"baseUrl 初始打开的链接"`
}

type CasesExecuteResBck struct {
	MustPageSnapshot   string            `json:"must_page_snapshot"`
	ErrPageSnapshot    string            `json:"err_page_snapshot"`
	FinishPageSnapShot string            `json:"finish_page_snap_shot"`
	StepMsg            []*StepExecuteRes `json:"step_msg"`
}

type CasesExecuteRes struct {
	IsSuccess       bool           `json:"is_success" dc:"执行结果,true === pass"`
	ErrPageSnapshot string         `json:"err_page_snapshot"`
	StepResults     []*StepResults `json:"step_results"`
	Duration        int64          `json:"duration"`
	StartTime       string         `json:"startTime"`
	HisId           string         `json:"hisId"`
	CaseDescription string         `json:"caseDescription"`
	Error           string         `json:"error"`
}

type StepResults struct {
	IsSuccess           bool         `json:"isSuccess" dc:"步骤执行成功 ture===成功"`
	Error               error        `json:"error" dc:"执行异常信息"`
	StepId              int          `json:"stepId" dc:"步骤ID"`
	CaseStepDescription string       `json:"caseStepDescription"`
	ElementName         string       `json:"elementName"`
	ElementPath         string       `json:"elementPath"`
	UnixMilli           int64        `json:unixMilli`
	AssertRes           []*AssertRes `json:"assertRes"`
}

type AssertRes struct {
	Method    string `json:"method" dc:"断言方法"`
	IsSuccess bool   `json:"isSuccess" dc:"断言是否通过"`
	Error     string `json:"error"`
	Message   string `json:"message" dc: "断言详情"`
}

type StepExecuteRes struct {
	R         interface{}
	AssertRes interface{}
}

type StepOpenPageRes struct {
	R interface{}
}

type StepExecuteReq struct {
	g.Meta `path:"/case/step-test" tags:"case" method:"post" summary:"step execute test"`
	Id     uint64 `json:"stepId" v:"required#id不能是0或空`
	Input  string `json:"input" dc:"遇到input输入框要输入的字段"`
	// AssertExpect *AssertExpect `json:"assertExpect" dc:"期望的断言"`
	AssertExpect string `json:"assertExpect" dc:"期望的断言"`
}

// type AssertExpect struct {
// 	Description *Description `json:"description"`
// 	Text        *Text        `json:"text"`
// }

// type Description struct {
// 	Rule   string
// 	Expect string
// }

// type Text struct {
// 	Rule   string
// 	Expect string
// }

type StepOpenPageReq struct {
	g.Meta `path:"/case/page-open" tags:"case" method:"post" summary:"open page test use for debug"`
	Url    string `json:"url" v:"required#url 不能空" dc:"page url"`
}

type StepOpenPageByIdReq struct {
	g.Meta `path:"/case/page-open-id" tags:"case" method:"post" summary:"open page test use for debug"`
	PageId int `json:"pageId" v:"required#page.id 不能空" dc:"page id"`
}

type CasesCreateReq struct {
	g.Meta      `path:"/case/create" tags:"case" method:"post" summary:"create Cases"`
	HisId       string `json:"hisId"  dc:"步骤结合，多个步骤使用英文逗号 , 分割"      ` //
	Description string `json:"description" dc:"用例描述，一般为测试目的"`          //
}

type CasesUpdateReq struct {
	g.Meta      `path:"/case/update" tags:"case" method:"post" summary:"update Cases"`
	Id          int    `json:"id" v:"required#Cases ID不能为空" dc:"id"`
	Description string `json:"description" dc:"用例描述，一般为测试目的"` //
	OpenId      int    `json:"openId" v:"required#Cases 初始pageId 不能为空" dc:"初始pageId"`
}

type CasesCreateRes struct {
	CaseId int64 `json:"caseId" dc:"create Cases id"`
}

type CasesUpdateRes struct {
	Ok bool `dc:"更新是否成功：true 成功 | false 失败" json:"ok"`
}

type CasesFindByIdReq struct {
	g.Meta `path:"/case/get-one" tags:"case" method:"get" summary:"get one Cases by id"`
	Id     uint64 `v:"required#id不能是0或空" dc:"Cases ID"`
}

type CasesGetOneRes struct {
	model.CasesOutput
}

type CasesGetListReq struct {
	g.Meta      `path:"/case/list" tags:"case" method:"get" summary:"get list of Cases"`
	CsId        string `json:"csId"  dc:"步骤结合，多个步骤使用英文逗号 , 分割"      ` //
	Description string `json:"description" dc:"用例描述，一般为测试目的"`         //
	Page        int    `dc:"分页号码" `                                   // 分页号码
	Limit       int    `dc:"分页数量"`                                    // 分页数量，最大50
	Sort        int    `dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"`               // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type CasesGetListRes struct {
	Row   interface{} `dc:"row" json:"row"`
	Total int         `json:"total" dc:"总数"`
	Page  int         `json:"page" dc:"当前页码"`
	Limit int         `json:"limit" dc:"分页大小"`
}
