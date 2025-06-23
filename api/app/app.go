// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package app

import (
	"context"

	"cloudphone/api/app/v1"
)

type IAppV1 interface {
	AppInstall(ctx context.Context, req *v1.AppInstallReq) (res *v1.AppInstallRes, err error)
	AppUninstall(ctx context.Context, req *v1.AppUninstallReq) (res *v1.AppUninstallRes, err error)
}
