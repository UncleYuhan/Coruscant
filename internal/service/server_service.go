package service

import (
	"context"
	"fmt"
	"os"

	v1 "cloudphone/api/server/v1"

	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"gopkg.in/yaml.v3"
)

type sServer struct{}

var Server = new(sServer)

const configFile = "manifest/config/config.yaml"

// 添加服务器
func (s *sServer) AddServer(ctx context.Context, req *v1.AddServerReq) (res *v1.AddServerRes, err error) {
	cfgData, err := loadYamlConfig()
	if err != nil {
		return nil, err
	}
	servers := getServersMap(cfgData)
	if _, exists := servers[req.Name]; exists {
		return nil, fmt.Errorf("服务器 [%s] 已存在", req.Name)
	}
	servers[req.Name] = toServerMap(&req.ServerInfo)
	cfgData["custom"].(map[string]interface{})["servers"] = servers
	err = saveYamlConfig(cfgData)
	return &v1.AddServerRes{}, err
}

// 删除服务器
func (s *sServer) DeleteServer(ctx context.Context, req *v1.DeleteServerReq) (res *v1.DeleteServerRes, err error) {
	cfgData, err := loadYamlConfig()
	if err != nil {
		return nil, err
	}
	servers := getServersMap(cfgData)
	if _, exists := servers[req.Name]; !exists {
		return nil, fmt.Errorf("服务器 [%s] 不存在", req.Name)
	}
	delete(servers, req.Name)
	cfgData["custom"].(map[string]interface{})["servers"] = servers
	err = saveYamlConfig(cfgData)
	return &v1.DeleteServerRes{}, err
}

// 更新服务器
func (s *sServer) UpdateServer(ctx context.Context, req *v1.UpdateServerReq) (res *v1.UpdateServerRes, err error) {
	cfgData, err := loadYamlConfig()
	if err != nil {
		return nil, err
	}
	servers := getServersMap(cfgData)
	if _, exists := servers[req.Name]; !exists {
		return nil, fmt.Errorf("服务器 [%s] 不存在", req.Name)
	}
	servers[req.Name] = toServerMap(&req.ServerInfo)
	cfgData["custom"].(map[string]interface{})["servers"] = servers
	err = saveYamlConfig(cfgData)
	return &v1.UpdateServerRes{}, err
}

// 列出服务器（注意方法名是单数：ListServer）
func (s *sServer) ListServer(ctx context.Context, req *v1.ListServerReq) (res *v1.ListServerRes, err error) {
	cfgData, err := loadYamlConfig()
	if err != nil {
		return nil, err
	}
	servers := getServersMap(cfgData)

	var list []v1.ServerInfo
	for name, v := range servers {
		item := gconv.Map(v)
		list = append(list, v1.ServerInfo{
			Name:            name,
			IP:              gconv.String(item["ip"]),
			SSHUser:         gconv.String(item["ssh_user"]),
			FlaskPort:       gconv.Int(item["flask_port"]),
			DockerPort:      gconv.Int(item["docker_port"]),
			WsScrcpyPort:    gconv.Int(item["ws_scrcpy_port"]),
			WsScrcpyAdbPort: gconv.Int(item["ws_scrcpy_adb_port"]),
		})
	}
	return &v1.ListServerRes{List: list}, nil
}

// =====================
// 工具函数
// =====================

func loadYamlConfig() (map[string]interface{}, error) {
	if !gfile.Exists(configFile) {
		return nil, fmt.Errorf("配置文件 [%s] 不存在", configFile)
	}
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("读取配置失败: %v", err)
	}
	var data map[string]interface{}
	if err := yaml.Unmarshal(bytes, &data); err != nil {
		return nil, fmt.Errorf("YAML 解析失败: %v", err)
	}
	if _, ok := data["custom"]; !ok {
		data["custom"] = make(map[string]interface{})
	}
	return data, nil
}

func saveYamlConfig(data map[string]interface{}) error {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("配置编码失败: %v", err)
	}
	if err := os.WriteFile(configFile, bytes, 0644); err != nil {
		return fmt.Errorf("保存配置失败: %v", err)
	}
	return nil
}

func getServersMap(data map[string]interface{}) map[string]interface{} {
	custom := data["custom"].(map[string]interface{})
	if servers, ok := custom["servers"].(map[string]interface{}); ok {
		return servers
	}
	servers := make(map[string]interface{})
	custom["servers"] = servers
	return servers
}

func toServerMap(req *v1.ServerInfo) map[string]interface{} {
	return map[string]interface{}{
		"ip":                 req.IP,
		"ssh_user":           req.SSHUser,
		"flask_port":         req.FlaskPort,
		"docker_port":        req.DockerPort,
		"ws_scrcpy_port":     req.WsScrcpyPort,
		"ws_scrcpy_adb_port": req.WsScrcpyAdbPort,
	}
}
