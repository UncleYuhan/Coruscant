package app

import (
	v1 "cloudphone/api/app/v1"
	"context"

	"cloudphone/internal/service"
)

// AppUninstall 卸载应用
func (c *ControllerV1) AppUninstall(ctx context.Context, req *v1.AppUninstallReq) (res *v1.AppUninstallRes, err error) {
	return service.AppService.AppUninstall(ctx, req)
}
