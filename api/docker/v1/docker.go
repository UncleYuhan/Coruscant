package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateContainerReq struct {
	g.Meta     `method:"post" path:"/docker/create" tags:"Docker" summary:"创建云手机容器"`
	Name       string `json:"name" v:"required"`
	Image      string `json:"image" v:"required"`
	DockerPort int    `json:"docker-port"  d:"5555"`
	ServerPort int    `json:"server-port"`
}

type CreateContainerRes struct {
	ContainerID string `json:"container_id"`
	// URL         string `json:"url"`  // docker这里返回web url不合适，还得adb，故而注释
}

type ListContainersReq struct {
	g.Meta `method:"get" path:"/docker/list" tags:"Docker" summary:"查询容器列表"`
}

type ListContainersRes struct {
	List []ContainerInfo `json:"list"`
}

type ContainerInfo struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
	Ports  string `json:"ports"`
}
