package model

import "wiki/internal/model/entity"

// ElementListInput 获取 element 列表
type ElementListInput struct {
	PageId       int // 关联PageID
	ElementName   string   // element 名称
	Description       string // element 描述
	Father int	//父结点
	Page       int    // 分页号码
	Limit       int    // 分页数量，最大50
	Sort       int    // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

// ElementListOutput 查询Element结果
type ElementListOutput struct {
	List  []*entity.Element `json:"list" description:"列表"`
	Page  int                        `json:"page" description:"分页码"`
	Limit  int                        `json:"limit" description:"分页数量"`
	Total int                        `json:"total" description:"数据总数"`
}

type ElementOutput struct {
	Element *entity.Element
	FatherElement *entity.Element
	Page *entity.Page
}

type ElementCreateInput struct {
	Name        string `json:"name"        ` //
	Description string `json:"description" ` //
	PageId      int    `json:"pageId"      ` //
	Path	string `json:"path"`
	Attribute   string `json:"attribute"` // 包含属性元素
	Father int	`json:"father"` //父结点
}




