package bitable

import (
	"Base/internal/config"
	"context"
	"fmt"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

// CreateBitable 用于创建一个新的飞书多维表格
func CreateBitable() {
	// 从配置中获取 AppID、AppSecret 和用户访问令牌
	appID := config.C.AppID
	appSecret := config.C.AppSecret
	userAccessToken := config.C.UserAccessToken

	// 如果配置不完整，给出提示
	if appID == "" || appSecret == "" || userAccessToken == "" {
		fmt.Println("❌ 配置文件缺少必要字段 app_id, app_secret 或 user_access_token")
		return
	}

	// 初始化飞书 SDK 客户端
	client := lark.NewClient(appID, appSecret)

	// 构造创建多维表格的请求（放在根目录下）
	req := larkbitable.NewCreateAppReqBuilder().
		ReqApp(larkbitable.NewReqAppBuilder().
			Name("CLI 创建的多维表格").
			FolderToken("").Build()).
		Build()

	// 调用 API 接口，使用用户身份进行创建
	resp, err := client.Bitable.V1.App.Create(
		context.Background(),
		req,
		larkcore.WithUserAccessToken(userAccessToken),
	)

	// 网络或请求失败处理
	if err != nil {
		fmt.Println("❌ 请求失败:", err)
		return
	}

	// 飞书服务端返回错误处理
	if !resp.Success() {
		fmt.Printf("❌ 接口调用失败（logId: %s）：%+v\n", resp.RequestId(), resp.CodeError)
		return
	}

	// 成功后输出结果信息
	fmt.Println("✅ 多维表格创建成功！")
	fmt.Println("📄 表格名称：", resp.Data.App.Name)
	fmt.Println("🔗 表格链接：", resp.Data.App.Url)
	fmt.Println("🆔 App Token：", resp.Data.App.AppToken)
}
