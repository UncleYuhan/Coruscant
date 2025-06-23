package service

import (
	"context"
	"fmt"
	"net/url"
	"os/exec"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

type sDocker struct{}

var Service = new(sDocker)

const (
	dockerHost      = "tcp://192.168.1.8:2375" // Docker 宿主机地址
	wsScrcpyHost    = "192.168.1.8:8000"       // ws-scrcpy 服务地址
	wsScrcpyAdbPort = 8890                     // 固定的 remote adb 端口，用于构造控制页面
)

// 获取 Docker 客户端
func getDockerClient() (*client.Client, error) {
	return client.NewClientWithOpts(
		client.WithHost(dockerHost),
		client.WithAPIVersionNegotiation(),
	)
}

// CreateContainer 创建并启动容器，返回控制页面访问地址
func (s *sDocker) CreateContainer(ctx context.Context, name, image string, hostPort, containerPort int) (map[string]string, error) {
	cli, err := getDockerClient()
	if err != nil {
		return nil, fmt.Errorf("create docker client failed: %w", err)
	}
	defer cli.Close()

	exposedPort := fmt.Sprintf("%d/tcp", containerPort)
	portSet := nat.PortSet{
		nat.Port(exposedPort): struct{}{},
	}
	portMap := nat.PortMap{
		nat.Port(exposedPort): []nat.PortBinding{{
			HostIP:   "0.0.0.0",
			HostPort: fmt.Sprintf("%d", hostPort),
		}},
	}

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:        image,
		ExposedPorts: portSet,
	}, &container.HostConfig{
		PortBindings: portMap,
		Privileged:   true,
	}, &network.NetworkingConfig{}, nil, name)
	if err != nil {
		return nil, fmt.Errorf("container create failed: %w", err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, fmt.Errorf("container start failed: %w", err)
	}

	// 使用宿主机端口映射替代容器内 IP
	adbAddr := fmt.Sprintf("127.0.0.1:%d", hostPort)

	// adb connect 宿主机映射端口
	out, err := exec.Command("adb", "connect", adbAddr).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("adb connect failed: %v (%s)", err, string(out))
	}

	// 构造控制页面地址（转义 udid 和 ws 参数）
	escapedUDID := url.QueryEscape(adbAddr)
	wsParam := url.QueryEscape(fmt.Sprintf("ws://%s/?action=proxy-adb&remote=tcp:%d&udid=%s", wsScrcpyHost, wsScrcpyAdbPort, adbAddr))
	controlPageURL := fmt.Sprintf("http://%s/#!action=stream&udid=%s&player=mse&ws=%s", wsScrcpyHost, escapedUDID, wsParam)

	result := map[string]string{
		"container_id": resp.ID[:12],
		"adb_device":   adbAddr,
		"scrcpy_url":   controlPageURL,
	}

	return result, nil
}

// ListContainers 获取所有容器信息
func (s *sDocker) ListContainers(ctx context.Context) ([]map[string]string, error) {
	cli, err := getDockerClient()
	if err != nil {
		return nil, fmt.Errorf("create docker client failed: %w", err)
	}
	defer cli.Close()

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		return nil, fmt.Errorf("list containers failed: %w", err)
	}

	var list []map[string]string
	for _, c := range containers {
		var ports []string
		for _, p := range c.Ports {
			ports = append(ports, fmt.Sprintf("%d->%d/%s", p.PublicPort, p.PrivatePort, p.Type))
		}
		list = append(list, map[string]string{
			"id":     c.ID[:12],
			"name":   strings.TrimPrefix(c.Names[0], "/"),
			"status": c.Status,
			"ports":  strings.Join(ports, ", "),
		})
	}
	return list, nil
}
