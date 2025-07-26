package v1

import "github.com/gogf/gf/v2/frame/g"

// 服务器通用字段（你已有）
type ServerInfo struct {
	Name            string `json:"name"         v:"required#服务器名称不能为空"`
	IP              string `json:"ip"           v:"required|ipv4#IP不能为空|IP格式错误"`
	SSHUser         string `json:"ssh_user"     v:"required#SSH 用户不能为空"`
	FlaskPort       int    `json:"flask_port"`
	DockerPort      int    `json:"docker_port"`
	WsScrcpyPort    int    `json:"ws_scrcpy_port"`
	WsScrcpyAdbPort int    `json:"ws_scrcpy_adb_port"`
}

type AddServerReq struct {
	g.Meta `path:"/server/add" method:"post" tags:"服务器管理" summary:"添加服务器"`
	ServerInfo
}
type AddServerRes struct{}

type DeleteServerReq struct {
	g.Meta `path:"/server/delete" method:"post" tags:"Server" summary:"删除服务器"`
	Name   string `json:"name" v:"required#服务器名称不能为空"`
}
type DeleteServerRes struct{}

type UpdateServerReq struct {
	g.Meta `path:"/server/update" method:"post" tags:"服务器管理" summary:"更新服务器配置"`
	ServerInfo
}
type UpdateServerRes struct{}

type ListServerReq struct {
	g.Meta `path:"/server/list" method:"get" tags:"服务器管理" summary:"列出所有服务器"`
}
type ListServerRes struct {
	List []ServerInfo `json:"list"`
}
