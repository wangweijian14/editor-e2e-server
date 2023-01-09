package cmd

import (
	"context"
	"net/http"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"wiki/internal/controller"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				//group.Middleware(MiddlewareHandlerResponse)
				group.Middleware(MiddlewareErrorHandler)
				group.Middleware(MiddlewareCORS)
				group.Bind(
					controller.Page,
					controller.Element,
					controller.Step,
					controller.Cases,
					controller.CaseStep,
					controller.CLaucher,
					controller.CaseReport,
				)
			})
			s.Run()
			return nil
		},
	}
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{Code: r.Response.Status, Message: r.GetError().Error()})
	}
}

// WikiTestHandlerResponse is the default implementation of HandlerResponse.
type WikiTestHandlerResponse struct {
	Code    int         `json:"code"    dc:"Error code"`
	Message string      `json:"message" dc:"Error message"`
	Data    interface{} `json:"data"    dc:"Result data for certain request according API definition"`
	Count   int         `json:"count" dc:"total"`
}
