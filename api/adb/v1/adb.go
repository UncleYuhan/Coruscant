package v1

import "github.com/gogf/gf/v2/frame/g"

type ListAdbDevicesReq struct {
	g.Meta `method:"get" path:"/devices/list" tags:"Docker" summary:"查询设备列表"`
}

type ListAdbDevicesRes struct {
	List []DevicesInfo `json:"list"`
}

type DevicesInfo struct {
	Ipport string `json:"ipport"`
	Type   string `json:"type"`
	Url    string `json:"url"`
}
