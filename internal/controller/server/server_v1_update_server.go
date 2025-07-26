package server

import (
	"context"

	v1 "cloudphone/api/server/v1"
	"cloudphone/internal/service"
)

func (c *ControllerV1) UpdateServer(ctx context.Context, req *v1.UpdateServerReq) (res *v1.UpdateServerRes, err error) {
	return service.Server.UpdateServer(ctx, req)
}
