// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package adb

import (
	"context"

	"cloudphone/api/adb/v1"
)

type IAdbV1 interface {
	ListAdbDevices(ctx context.Context, req *v1.ListAdbDevicesReq) (res *v1.ListAdbDevicesRes, err error)
}
