package v1

import (
	"wiki/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type ElementCreateReq struct {
	g.Meta      `path:"/element/create" tags:"element" method:"post" summary:"create element"`
	Name        string `json:"name" v:"required#请输入名称"   dc:"Element名称"`
	Description string `json:"description" v:"required#请输入Element描述"   dc:"用途描述"`
	PageId      int    `json:"pageId"  v:"required#请输所属页面id"   dc:"所属页面id"` //
	Path        string `json:"path"  v:"required#元素path不能为空"   dc:"元素path"` //
	Attribute   string `json:"attribute"  dc:"包含属性元素" `                     //
	Father      int    `json:"father" dc:"父结点"`
}

type ElementUpdateReq struct {
	g.Meta      `path:"/element/update" tags:"element" method:"post" summary:"update element"`
	Id          int    `json:"id" v:"required#element ID不能为空" dc:"id"`
	Name        string `json:"name"   dc:"element 名称"`
	Description string `json:"description"    dc:"element 用途描述"`
	PageId      int    `json:"pageId"  dc:"element 所属页面id"` //
	Path        string `json:"path"  dc:"元素path"`           //
	Attribute   string `json:"attribute"  dc:"包含属性元素" `     //
	Father      int    `json:"father" dc:"父结点"`
}

type ElementRemoveReq struct {
	g.Meta `path:"/element/remove" tags:"element" method:"post" summary:"remove element"`
	Id     int `json:"id" v:"required#element ID不能为空" dc:"id"`
}

type ElementCreateRes struct {
	PageId int64 `dc:"create element id"`
}

type ElementUpdateRes struct {
	Ok bool `dc:"更新是否成功：true 成功 | false 失败" json:"ok"`
}

type ElementFindByIdReq struct {
	g.Meta `path:"/element/get-one" tags:"element" method:"get" summary:"get one element by id"`
	Id     uint64 `v:"required#id不能是0或空" dc:"element ID"`
}

type ElementGetOneRes struct {
	Page          *entity.Page
	Element       *entity.Element
	FatherElement *entity.Element
}

type ElementGetListReq struct {
	g.Meta      `path:"/element/list" tags:"element" method:"get" summary:"get list of element"`
	Description string `json:"description" dc:"element 描述"`  // page 描述
	ElementName string `json:"element_name" dc:"element 名称"` // page 名称
	PageId      int    `json:"page_id" dc:"关联页面 id不能是0或空"`
	Father      int    `json:"father" dc:"父结点"`
	Page        int    `json:"page" dc:"分页号码" `                     // 分页号码
	Limit       int    `json:"limit" dc:"分页数量"`                     // 分页数量，最大50
	Sort        int    `json:"sort" dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"` // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type ElementGetListRes struct {
	Row   interface{} `dc:"row" json:"row"`
	Total int         `json:"total" dc:"总数"`
	Page  int         `json:"page" dc:"当前页码"`
	Limit int         `json:"limit" dc:"分页大小"`
}
