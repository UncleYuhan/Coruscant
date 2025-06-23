package docker

import (
	"context"

	v1 "cloudphone/api/docker/v1"
	"cloudphone/internal/service"
)

func (c *ControllerV1) ListContainers(ctx context.Context, req *v1.ListContainersReq) (res *v1.ListContainersRes, err error) {
	list, err := service.Service.ListContainers(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.ListContainersRes{
		List: make([]v1.ContainerInfo, len(list)),
	}
	for i, item := range list {
		res.List[i] = v1.ContainerInfo{
			ID:     item["id"],
			Name:   item["name"],
			Status: item["status"],
			Ports:  item["ports"],
		}
	}
	return
}
