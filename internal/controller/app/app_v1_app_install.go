package app

import (
	v1 "cloudphone/api/app/v1"
	"context"

	"cloudphone/internal/service"
)

// AppInstall 安装 APK
func (c *ControllerV1) AppInstall(ctx context.Context, req *v1.AppInstallReq) (res *v1.AppInstallRes, err error) {
	return service.AppService.AppInstall(ctx, req)
}

// AppList 获取设备已安装应用包名
func (c *ControllerV1) AppList(ctx context.Context, req *v1.AppRequest) (res *v1.AppPackageListRes, err error) {
	return service.AppService.AppList(ctx, req)
}
