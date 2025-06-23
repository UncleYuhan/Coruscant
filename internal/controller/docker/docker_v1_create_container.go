package docker

import (
	"context"

	v1 "cloudphone/api/docker/v1"
	"cloudphone/internal/service"
)

func (c *ControllerV1) CreateContainer(ctx context.Context, req *v1.CreateContainerReq) (*v1.CreateContainerRes, error) {
	result, err := service.Service.CreateContainer(ctx, req.Name, req.Image, req.ServerPort, req.DockerPort)
	if err != nil {
		return nil, err
	}
	return &v1.CreateContainerRes{
		ContainerID: result["container_id"],
		// URL:         result["scrcpy_url"],
	}, nil
}
