package middleware

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

// CORS 跨域处理中间件
func CORS(r *ghttp.Request) {
	g.Log().Info(r.GetCtx(), "🔥 CORS middleware triggered")
	r.Response.CORSDefault()
	if r.Method == "OPTIONS" {
		r.Exit()
	}
}
