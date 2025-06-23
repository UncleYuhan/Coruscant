package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"cloudphone/internal/controller/adb"    // 👈 记得引入
	"cloudphone/internal/controller/app"    // 👈 记得引入
	"cloudphone/internal/controller/docker" // 👈 记得引入
	"cloudphone/internal/controller/hello"
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

				// 注册多个 controller
				group.Bind(
					hello.NewV1(),
					docker.NewV1(), // 👈 注册 docker 接口
					adb.NewV1(),    // 👈 注册 adb 接口
					app.NewV1(),    // 👈 注册 app 接口
				)
			})

			s.Run()
			return nil
		},
	}
)
