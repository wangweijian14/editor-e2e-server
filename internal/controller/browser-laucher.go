package controller

import (
	"context"
	v1 "wiki/api/v1"
	"wiki/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	CMonitor = cMonitor{}
)

type cMonitor struct{}

func (c *cMonitor) GetServeMonitorPath(ctx context.Context, req *v1.MonitorPathGetReq) (res *v1.MonitorPathGetRes, err error) {
	return &v1.MonitorPathGetRes{
		MonitorPath: service.Action().GetServeMonitorPath(),
	}, nil
}

func (c *cMonitor) GetMonitorPages(ctx context.Context, req *v1.MonitorPagesReq) (res *v1.MonitorPagesRes, err error) {
	pages, err := service.Action().GetMonitorPages()
	if err != nil {
		return nil, err
	}
	return &v1.MonitorPagesRes{
		Pages: pages,
	}, nil
}

func (c *cMonitor) GetMonitorPageInfo(ctx context.Context, req *v1.MonitorPageOneReq) (res *v1.MonitorPageOneRes, err error) {
	p, err := service.Action().GetMonitorPageInfo(req.PageId)
	if err != nil {
		return nil, err
	}

	return &v1.MonitorPageOneRes{
		PageInfo: p,
	}, nil
}

func (c *cMonitor) GetScreenshot(ctx context.Context, req *v1.MonitorScreenshotReq) (res *v1.MonitorScreenshotRes, err error) {
	picByte := service.Action().GetMonitorPageScreenshot(req.PageId)
	g.RequestFromCtx(ctx).Response.Write(picByte)
	return &v1.MonitorScreenshotRes{}, nil
}
