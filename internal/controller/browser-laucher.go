package controller

import (
	"context"
	v1 "wiki/api/v1"
	"wiki/internal/service"
)

var (
	CLaucher = cLaucher{}
)

type cLaucher struct{}

func (c *cLaucher) GetServeMonitorPath(ctx context.Context, req *v1.LaucherGetReq) (res *v1.LaucherGetRes, err error) {
	return &v1.LaucherGetRes{
		LautherPath: service.Action().GetServeMonitorPath(),
	}, nil
}
