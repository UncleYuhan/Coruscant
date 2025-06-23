package adb

import (
	"context"

	v1 "cloudphone/api/adb/v1"
	"cloudphone/internal/service"
)

func (c *ControllerV1) ListAdbDevices(ctx context.Context, req *v1.ListAdbDevicesReq) (*v1.ListAdbDevicesRes, error) {
	return service.AdbService.ListAdbDevices(ctx, req)
}
