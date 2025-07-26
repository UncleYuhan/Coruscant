package server

import (
	"context"

	v1 "cloudphone/api/server/v1"
	"cloudphone/internal/service"
)

func (c *ControllerV1) DeleteServer(ctx context.Context, req *v1.DeleteServerReq) (res *v1.DeleteServerRes, err error) {
	return service.Server.DeleteServer(ctx, req)
}
