package v1

import "github.com/gogf/gf/v2/frame/g"

// 获取应用包名列表
type AppRequest struct {
	g.Meta `method:"get" path:"/app/list" tags:"App" summary:"获取已安装包名"`
	Target string `json:"target" v:"required"`
}
type AppPackageListRes struct {
	List []string `json:"list"`
}

// 安装 APK
type AppInstallReq struct {
	g.Meta `method:"post" path:"/app/install" tags:"App" summary:"安装应用"`
	Target string `json:"target" v:"required"`
	Apk    string `json:"apk" v:"required"` // 服务器路径
}
type AppInstallRes struct {
	Output string `json:"output"`
}

// 卸载应用
type AppUninstallReq struct {
	g.Meta  `method:"delete" path:"/app/uninstall" tags:"App" summary:"卸载应用"`
	Target  string `json:"target" v:"required"`
	Package string `json:"package" v:"required"`
}
type AppUninstallRes struct {
	Output string `json:"output"`
}
