// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package server

import (
	"context"

	v1 "cloudphone/api/server/v1"
)

type IServerV1 interface {
	AddServer(ctx context.Context, req *v1.AddServerReq) (res *v1.AddServerRes, err error)
	DeleteServer(ctx context.Context, req *v1.DeleteServerReq) (res *v1.DeleteServerRes, err error)
	UpdateServer(ctx context.Context, req *v1.UpdateServerReq) (res *v1.UpdateServerRes, err error)
	ListServer(ctx context.Context, req *v1.ListServerReq) (res *v1.ListServerRes, err error)
}
