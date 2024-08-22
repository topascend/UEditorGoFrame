package cmd

import (
	// 百度编辑器插件
	ueditorController "github.com/topascend/go-gf-ueditor/controller"

	"UEditorGoFrame/internal/controller/hello"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
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

				group.Bind(
					hello.NewV1(),
				)

				// 绑定ueditor插件
				group.Bind(
					ueditorController.NewUEditor(),
				)
			})
			s.Run()
			return nil
		},
	}
)
