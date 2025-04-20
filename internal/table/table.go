package create

import (
	"context"
	"fmt"
	"os"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	larkbitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

func CreateTable(appToken string) {
	appID := os.Getenv("APP_ID")
	appSecret := os.Getenv("APP_SECRET")
	userAccessToken := os.Getenv("USER_ACCESS_TOKEN")

	if appID == "" || appSecret == "" || userAccessToken == "" {
		fmt.Println("❌ 缺少环境变量：APP_ID、APP_SECRET 或 USER_ACCESS_TOKEN")
		return
	}

	client := lark.NewClient(appID, appSecret)

	req := larkbitable.NewCreateAppTableReqBuilder().
		AppToken(appToken).
		Body(larkbitable.NewCreateAppTableReqBodyBuilder().
			Table(larkbitable.NewReqTableBuilder().
				Name("产品数据表").
				DefaultViewName("默认视图").
				Fields([]*larkbitable.AppTableCreateHeader{
					larkbitable.NewAppTableCreateHeaderBuilder().
						FieldName("索引字段").
						Type(1). // 文本
						Build(),
					larkbitable.NewAppTableCreateHeaderBuilder().
						FieldName("产品状态").
						Type(3). // 单选
						UiType("SingleSelect").
						Property(larkbitable.NewAppTableFieldPropertyBuilder().
							Options([]*larkbitable.AppTableFieldPropertyOption{
								larkbitable.NewAppTableFieldPropertyOptionBuilder().
									Name("上线").Color(0).Build(),
								larkbitable.NewAppTableFieldPropertyOptionBuilder().
									Name("下线").Color(1).Build(),
								larkbitable.NewAppTableFieldPropertyOptionBuilder().
									Name("开发中").Color(2).Build(),
							}).Build()).
						Build(),
				}).Build()).Build()).Build()

	resp, err := client.Bitable.V1.AppTable.Create(context.Background(), req, larkcore.WithUserAccessToken(userAccessToken))
	if err != nil {
		fmt.Println("❌ 请求失败:", err)
		return
	}
	if !resp.Success() {
		fmt.Printf("❌ 接口调用失败：%+v\n", resp.CodeError)
		return
	}

	fmt.Println("✅ 数据表创建成功！")
	fmt.Println("🆔 表格 ID：", resp.Data.TableId)         // 新创建的数据表 ID
	fmt.Println("🔍 默认视图 ID：", resp.Data.DefaultViewId) // 表格的默认视图 ID
	fmt.Println("📋 字段 ID 列表：", resp.Data.FieldIdList)  // 创建表格时包含的字段 ID 集合
}
