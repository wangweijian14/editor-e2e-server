package v1

import (
	"github.com/go-rod/rod/lib/proto"
	"github.com/gogf/gf/v2/frame/g"
)

type MonitorPathGetReq struct {
	g.Meta `path:"/monitor/path" tags:"monitor" method:"get" summary:"获取 monitorPath"`
}

type MonitorPathGetRes struct {
	MonitorPath string `json:"monitorPath" dc:"monitor path"`
}

type MonitorPagesReq struct {
	g.Meta `path:"/monitor/pages" tags:"monitor" method:"get" summary:"获取 monitor pages"`
}

type MonitorPagesRes struct {
	Pages []*proto.TargetTargetInfo
}

type MonitorPageOneReq struct {
	g.Meta `path:"/monitor/page/one" tags:"monitor" method:"get" summary:"获取 monitor page by pageId"`
	PageId string `json:"pageId" dc:"page id"`
}

type MonitorPageOneRes struct {
	PageInfo *proto.TargetTargetInfo `json:"pageInfo"`
}

type MonitorScreenshotReq struct {
	g.Meta `path:"/monitor/screenshot" tags:"monitor" method:"get" summary:"获取 monitor page screenshot by pageId"`
	PageId string `json:"pageId" dc:"page id"`
}

type MonitorScreenshotRes struct {
	g.Meta `Content-Type:"image/png;"`
}
