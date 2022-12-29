package v1

import (
	"wiki/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type PageCreateReq struct {
	g.Meta      `path:"/page/create" tags:"page" method:"post" summary:"create page"`
	Name        string `json:"name" v:"required#请输入名称"   dc:"页面名称"`
	Description string `json:"description" v:"required#请输入页面描述"   dc:"用途描述"`
	Url         string `json:"url" dc:"page.url"`
}

type PageUpdateReq struct {
	g.Meta      `path:"/page/update" tags:"page" method:"post" summary:"update page"`
	Id          int    `json:"id" v:"required#页面ID不能为空" dc:"id"`
	Name        string `json:"name"   dc:"页面名称"`
	Description string `json:"description"    dc:"用途描述"`
	Url         string `json:"url" dc:"page.url"`
}

type PageCreateRes struct {
	PageId int64 `dc:"create page id"`
}

type PageUpdateRes struct {
	Ok bool `dc:"更新是否成功：true 成功 | false 失败" json:"ok"`
}

type PageFindByIdReq struct {
	g.Meta `path:"/page/get-one" tags:"page" method:"get" summary:"get one page by id"`
	Id     uint64 `v:"required#id不能是0或空" dc:"Page ID"`
}

type PageFindPageByIdWithElementReq struct {
	g.Meta `path:"/page/get-one-with-elements" tags:"page" method:"get" summary:"get one page with elements by id"`
	Id     uint64 `v:"required#id不能是0或空" dc:"Page ID"`
}

type PageGetOneRes struct {
	Page *entity.Page
}

type PageGetOneWithElementsRes struct {
	Page     *entity.Page
	Elements []*entity.Element
}

type PageGetListReq struct {
	g.Meta      `path:"/page/list" tags:"page" method:"get" summary:"get list of page"`
	Description string `json:"description" dc:"Page 描述"` // page 描述
	PageName    string `json:"page_name" dc:"Page 名称"`   // page 名称
	Page        int    `dc:"分页号码" `                      // 分页号码
	Limit       int    `dc:"分页数量"`                       // 分页数量，最大50
	Sort        int    `dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"`  // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type PageTreeReq struct {
	g.Meta      `path:"/page/tree" tags:"page" method:"post" summary:"get list of page"`
	Description string `json:"description" dc:"Page 描述"` // page 描述
	PageName    string `json:"page_name" dc:"Page 名称"`   // page 名称
	Page        int    `dc:"分页号码" `                      // 分页号码
	Limit       int    `dc:"分页数量"`                       // 分页数量，最大50
	Sort        int    `dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"`  // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type PageGetListRes struct {
	Row   interface{} `dc:"row" json:"row"`
	Total int         `json:"total" dc:"总数"`
	Page  int         `json:"page" dc:"当前页码"`
	Limit int         `json:"limit" dc:"分页大小"`
}

type PageGetTreeRes struct {
	Row   interface{} `dc:"row" json:"row"`
	Total int         `json:"total" dc:"总数"`
	Page  int         `json:"page" dc:"当前页码"`
	Limit int         `json:"limit" dc:"分页大小"`
}
