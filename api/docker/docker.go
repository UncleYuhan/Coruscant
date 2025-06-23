// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package docker

import (
	"context"

	"cloudphone/api/docker/v1"
)

type IDockerV1 interface {
	CreateContainer(ctx context.Context, req *v1.CreateContainerReq) (res *v1.CreateContainerRes, err error)
	ListContainers(ctx context.Context, req *v1.ListContainersReq) (res *v1.ListContainersRes, err error)
}
