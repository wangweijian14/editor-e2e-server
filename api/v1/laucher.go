package v1

import "github.com/gogf/gf/v2/frame/g"

type LaucherGetReq struct {
	g.Meta `path:"/laucher/get" tags:"case" method:"get" summary:"获取laucherPath"`
}

type LaucherGetRes struct {
	LautherPath string `json:"laucherPath" dc:"laucher path"`
}
