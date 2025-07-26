package server

import (
	v1 "cloudphone/api/server/v1"
	"cloudphone/internal/service"
	"context"
)

func (c *ControllerV1) ListServer(ctx context.Context, req *v1.ListServerReq) (res *v1.ListServerRes, err error) {
	return service.Server.ListServer(ctx, req)
}
