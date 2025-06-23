package service

import (
	v1 "cloudphone/api/adb/v1"
	"context"
	"fmt"
	"net/url"
	"os/exec"
	"strings"
)

type sAdb struct{}

var AdbService = new(sAdb)

// runRemoteAdb 在远程服务器执行 adb 命令
func runRemoteAdb(cmd string) ([]byte, error) {
	return exec.Command("ssh", "yyy@192.168.1.8", fmt.Sprintf("adb %s", cmd)).CombinedOutput()
}

func (s *sAdb) ListAdbDevices(ctx context.Context, _ *v1.ListAdbDevicesReq) (*v1.ListAdbDevicesRes, error) {
	// 远程断开所有连接，避免干扰
	_, _ = runRemoteAdb("disconnect")

	// 获取远程已连接设备
	output, err := runRemoteAdb("devices")
	if err != nil {
		return nil, fmt.Errorf("remote adb devices failed: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	connectedMap := make(map[string]bool)
	for _, line := range lines[1:] {
		fields := strings.Fields(strings.TrimSpace(line))
		if len(fields) >= 2 && fields[1] == "device" {
			connectedMap[fields[0]] = true
		}
	}

	containers, err := Service.ListContainers(ctx)
	if err != nil {
		return nil, fmt.Errorf("list containers failed: %w", err)
	}

	seen := make(map[string]bool)
	var adbTargets []string
	for _, c := range containers {
		ports := strings.Split(c["ports"], ",")
		for _, p := range ports {
			p = strings.TrimSpace(p)
			if strings.Contains(p, "->") {
				parts := strings.Split(p, "->")
				hostPort := strings.Split(parts[0], "/")[0]
				adbAddr := fmt.Sprintf("127.0.0.1:%s", hostPort)
				if !seen[adbAddr] {
					adbTargets = append(adbTargets, adbAddr)
					seen[adbAddr] = true
				}
			}
		}
	}

	// 尝试远程 adb connect 所有未连接目标
	for _, addr := range adbTargets {
		if !connectedMap[addr] {
			_, _ = runRemoteAdb(fmt.Sprintf("connect %s", addr))
		}
	}

	// 等待连接生效（可选）
	_ = exec.Command("sleep", "1").Run()

	// 重新获取远程设备状态
	output2, _ := runRemoteAdb("devices")
	lines2 := strings.Split(string(output2), "\n")
	finalStatus := make(map[string]string)
	for _, line := range lines2[1:] {
		fields := strings.Fields(strings.TrimSpace(line))
		if len(fields) >= 2 {
			finalStatus[fields[0]] = fields[1]
		}
	}

	// 构造最终响应
	var list []v1.DevicesInfo
	for _, addr := range adbTargets {
		encodedUdid := url.QueryEscape(addr)
		mseUrl := fmt.Sprintf("http://192.168.1.8:8000/#!action=stream&udid=%s&player=mse&ws=ws%%3A%%2F%%2F192.168.1.8%%3A8000%%2F%%3Faction%%3Dproxy-adb%%26remote%%3Dtcp%%253A8890%%26udid%%3D%s", encodedUdid, encodedUdid)

		list = append(list, v1.DevicesInfo{
			Ipport: addr,
			Type:   finalStatus[addr],
			Url:    mseUrl,
		})
	}

	return &v1.ListAdbDevicesRes{List: list}, nil
}
