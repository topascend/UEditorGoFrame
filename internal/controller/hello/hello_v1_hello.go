package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"

	"UEditorGoFrame/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World!")
	return
}

// Ueditor 接口 测试
func (c *ControllerV1) Ueditor(r *ghttp.Request) {
	_ = r.Response.WriteTpl("ueditor/index.html", g.Map{})
}
