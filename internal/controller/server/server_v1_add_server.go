package server

import (
	"context"

	v1 "cloudphone/api/server/v1"
	"cloudphone/internal/service"
)

func (c *ControllerV1) AddServer(ctx context.Context, req *v1.AddServerReq) (res *v1.AddServerRes, err error) {
	return service.Server.AddServer(ctx, req)
}
