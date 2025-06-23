package service

import (
	"bytes"
	v1 "cloudphone/api/app/v1"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type sApp struct{}

var AppService = new(sApp)

// 通用响应格式
// { "code": 0, "message": "OK", "data": { "output": "..." }}
type commonOutput struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data"`
}

type installData struct {
	Output string `json:"output"`
}

type uninstallData struct {
	Output string `json:"output"`
}

type packageListData struct {
	Packages []string `json:"packages"`
}

// AppInstall 安装 APK
func (s *sApp) AppInstall(ctx context.Context, req *v1.AppInstallReq) (*v1.AppInstallRes, error) {
	url := "http://192.168.1.8:8888/install"
	body, _ := json.Marshal(map[string]string{
		"target": req.Target,
		"apk":    req.Apk,
	})
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var common commonOutput
	if err := json.NewDecoder(resp.Body).Decode(&common); err != nil {
		return nil, err
	}

	var data installData
	if err := json.Unmarshal(common.Data, &data); err != nil {
		return nil, err
	}

	return &v1.AppInstallRes{Output: data.Output}, nil
}

// AppUninstall 卸载应用
func (s *sApp) AppUninstall(ctx context.Context, req *v1.AppUninstallReq) (*v1.AppUninstallRes, error) {
	url := "http://192.168.1.8:8888/uninstall"
	body, _ := json.Marshal(map[string]string{
		"target":  req.Target,
		"package": req.Package,
	})
	reqHttp, _ := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer(body))
	reqHttp.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(reqHttp)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var common commonOutput
	if err := json.NewDecoder(resp.Body).Decode(&common); err != nil {
		return nil, err
	}

	var data uninstallData
	if err := json.Unmarshal(common.Data, &data); err != nil {
		return nil, err
	}

	return &v1.AppUninstallRes{Output: data.Output}, nil
}

// AppList 获取设备已安装应用包名
func (s *sApp) AppList(ctx context.Context, req *v1.AppRequest) (*v1.AppPackageListRes, error) {
	url := fmt.Sprintf("http://192.168.1.8:8888/packages?target=%s", req.Target)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var common commonOutput
	if err := json.NewDecoder(resp.Body).Decode(&common); err != nil {
		return nil, err
	}

	var data packageListData
	if err := json.Unmarshal(common.Data, &data); err != nil {
		return nil, err
	}

	return &v1.AppPackageListRes{List: data.Packages}, nil
}
