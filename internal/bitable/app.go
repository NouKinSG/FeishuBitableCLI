package bitable

import (
	"context"
	"fmt"
	"github.com/NouKinSG/FeishuBitableCLI/internal/config"
	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

// AppInfo 是供 TUI 展示的简化结构体
type AppInfo struct {
	AppToken string
	Name     string
	Url      string
}

// CreateApp 用于创建一个新的多维表格，并返回其信息
func CreateApp(name, folderToken string) (*AppInfo, error) {
	cli := lark.NewClient(config.C.AppID, config.C.AppSecret)
	req := larkbitable.NewCreateAppReqBuilder().
		ReqApp(
			larkbitable.NewReqAppBuilder().
				Name(name).
				FolderToken(folderToken).
				Build(),
		).
		Build()

	resp, err := cli.Bitable.V1.App.Create(
		context.Background(),
		req,
		larkcore.WithUserAccessToken(config.C.UserAccessToken),
	)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	if !resp.Success() {
		return nil, fmt.Errorf("服务端返回错误: code=%d msg=%s", resp.Code, resp.Msg)
	}

	app := resp.Data.App
	var token, nm, url string
	if app.AppToken != nil {
		token = *app.AppToken
	}
	if app.Name != nil {
		nm = *app.Name
	}
	if app.Url != nil {
		url = *app.Url
	}

	return &AppInfo{
		AppToken: token,
		Name:     nm,
		Url:      url,
	}, nil
}
