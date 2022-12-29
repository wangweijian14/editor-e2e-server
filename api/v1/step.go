package v1

import (
	"wiki/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type StepCreateReq struct {
	g.Meta          `path:"/step/create" tags:"step" method:"post" summary:"create step"`
	Description     string `json:"description" v:"required#请输入step描述"   dc:"用途描述"`
	TargetPage      int    `json:"target_page"  v:"required#请输所属页面id"   dc:"所属页面id"`              //
	TargetElementId int    `json:"targetElementId" v:"required#请输入目标elementId"   dc:"目标 element"` //
	ActionId        int    `json:"actionId"  v:"required#动作 id不能空 "   dc:"动作Id"      `            //
}

type StepUpdateReq struct {
	g.Meta          `path:"/step/update" tags:"step" method:"post" summary:"update step"`
	Id              int    `json:"id" v:"required#step ID不能为空" dc:"id"`
	Description     string `json:"description"  dc:"用途描述"`
	TargetPage      int    `json:"target_page"  dc:"所属页面id"`        //
	TargetElementId int    `json:"targetElementId" dc:"目标 element"` //
	ActionId        int    `json:"actionId" dc:"动作 Id"      `       //
}

type StepRemoveReq struct {
	g.Meta `path:"/step/remove" tags:"step" method:"post" summary:"remove step"`
	Id     int `json:"id" v:"required#step ID不能为空" dc:"stip id 用户删除step" `
}

type StepCreateRes struct {
	Id int64 `dc:"create step id"`
}

type StepUpdateRes struct {
	Ok bool `dc:"更新是否成功：true 成功 | false 失败" json:"ok"`
}

type StepFindByIdReq struct {
	g.Meta `path:"/step/get-one" tags:"step" method:"get" summary:"get one step by id"`
	Id     uint64 `v:"required#id不能是0或空" dc:"step ID"`
}

type StepGetOneRes struct {
	Page    *entity.Page
	Step    *entity.Step
	Element *entity.Element
}

type StepGetListReq struct {
	g.Meta          `path:"/step/list" tags:"step" method:"get" summary:"get list of step"`
	Description     string `dc:"step 描述"`                           // step 描述
	TargetPage      int    `json:"target_page"  dc:"所属页面id"`        //
	TargetElementId int    `json:"targetElementId" dc:"目标 element"` //
	ActionId        int    `json:"actionId" dc:"动作 Id"      `       //
	Ids             string `json:"ids" dc:"ids "`
	Page            int    `dc:"分页号码" `                     // 分页号码
	Limit           int    `dc:"分页数量"`                      // 分页数量，最大50
	Sort            int    `dc:"排序类型(0:最新, 默认。1:活跃, 2:热度)"` // 排序类型(0:最新, 默认。1:活跃, 2:热度)
}

type StepGetListRes struct {
	Row   interface{} `dc:"row" json:"row"`
	Total int         `json:"total" dc:"总数"`
	Page  int         `json:"page" dc:"当前页码"`
	Limit int         `json:"limit" dc:"分页大小"`
}
