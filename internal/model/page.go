package model

import (
	"wiki/internal/model/entity"

	_ "github.com/gogf/gf/v2/os/gtime"
)

// PageGetListInput 获取内容列表
type PageGetListInput struct {
	Description string // page 描述
	PageName    string // page 名称
	Url         string // page.url
	Page        int    // 分页号码
	Limit       int    // 分页数量，最大50
	Sort        int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// PageGetListOutput 查询页面结果
type PageGetListOutput struct {
	List  PageGetListOutputItem `json:"list" description:"列表"`
	Page  int                   `json:"page" description:"分页码"`
	Limit int                   `json:"limit" description:"分页数量"`
	Total int                   `json:"total" description:"数据总数"`
}

type PageGetListOutputItem struct {
	Page     []*entity.Page    `json:"page"`
	Elements []*entity.Element `json:"elements"`
}

type PageOutputItem struct {
	Page     *entity.Page      `json:"page"`
	Elements []*entity.Element `json:"elements"`
}

type PageCreateInput struct {
	Description string // page 描述
	Name        string // page 名称
	Url         string //page.url
}

type PageUpdateInput struct {
	Id          int
	Description string // page 描述
	Name        string // page 名称
	Url         string //page.url
}

type PageCreateOutput struct {
	PageId int64 `json:"page_id"`
}
