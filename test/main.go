package main

import (
	"context"
	"fmt"
	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/drive/v1"
)

// SDK 使用文档：https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/server-side-sdk/golang-sdk-guide/preparations
// 复制该 Demo 后, 需要将 "YOUR_APP_ID", "YOUR_APP_SECRET" 替换为自己应用的 APP_ID, APP_SECRET.
// 以下示例代码默认根据文档示例值填充，如果存在代码问题，请在 API 调试台填上相关必要参数后再复制代码使用
func main() {
	// 创建 Client
	client := lark.NewClient("cli_a88b5b4afe39901c", "uFOs3nKwFTZubV22ATrM1cCAmFuJUaZK")
	// 创建请求对象
	req := larkdrive.NewTransferOwnerPermissionMemberReqBuilder().
		Token(`LtDkbzUFHa7vsHs6JmDc2la9nlg`).
		Type(`bitable`).
		NeedNotification(true).
		RemoveOldOwner(false).
		StayPut(false).
		OldOwnerPerm(`full_access`).
		Owner(larkdrive.NewOwnerBuilder().
			MemberType(`openid`).
			MemberId(`ou_b0a1468d3bb7b754e64d3b6a657490df`).
			Build()).
		Build()

	// 发起请求
	resp, err := client.Drive.V1.PermissionMember.TransferOwner(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Printf("logId: %s, error response: \n%s", resp.RequestId(), larkcore.Prettify(resp.CodeError))
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
