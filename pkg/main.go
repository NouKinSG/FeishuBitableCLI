package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkcore "github.com/larksuite/oapi-sdk-go/v3/core"
	bitable "github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
)

// —— 配置区，请成自己的凭证 ——
const (
	AppID       = "cli_a88b5b4afe39901c"
	AppSecret   = "uFOs3nKwFTZubV22ATrM1cCAmFuJUaZK"
	FolderToken = "Ns8DfRkkXl9vDedVBcUcM2XUnqh"                    // 多维表格要创建到的飞书文件夹
	UserToken   = "u-djnPiKaYl5PrLnd6QVuC9sh466aM40argwG0k1.GyKz_" // user_access_token 或 tenant_access_token
	viewID      = "vewxEfYatA"                                     // 指定视图
)

// DemoManager 负责整个示例的执行流程
type DemoManager struct {
	client   *lark.Client
	ctx      context.Context
	appToken string
	tableID  string
}

// fundField 定义每个字段的元信息
type fundField struct {
	Name    string   // 字段名
	Type    int      // type 枚举
	UIType  string   // ui_type
	Options []string // 枚举可选值
	Multi   bool     // 是否多选
}

// 私募基金产品表字段定义
var fundFields = []fundField{
	{"产品全称", 1, "Text", nil, false},
	{"产品简称", 1, "Text", nil, false},
	{"托管机构", 3, "SingleSelect", []string{"A机构", "B机构", "C机构"}, false},
	{"管理人", 3, "SingleSelect", []string{"管A", "管B"}, false},
	{"TA", 11, "User", nil, false},
	{"状态", 3, "SingleSelect", []string{"开放", "封闭"}, false},
	{"成立日期", 5, "DateTime", nil, false},
	{"最小追加金额", 2, "Number", nil, false},
	{"产品策略", 4, "MultiSelect", []string{"策略A", "策略B", "策略C"}, true},
	{"备注", 1, "Text", nil, false},
}

func main() {
	rand.Seed(time.Now().UnixNano())
	NewDemoManager().Run()
}

// NewDemoManager 初始化 DemoManager，并创建底层 SDK Client
func NewDemoManager() *DemoManager {
	ctx := context.Background()
	client := lark.NewClient(AppID, AppSecret)
	return &DemoManager{client: client, ctx: ctx}
}

// Run 按顺序执行：创建 App → 创建表 → 插入初始数据 → 随机更新记录
func (m *DemoManager) Run() {
	m.appToken = "Xy0abjzi6a796hs3SsRc0XJSnyf"
	m.tableID = "tblmZMRQQVBHI6Z0"

	////第 1 步：创建 App
	//var err error
	//m.appToken, err = m.createApp()
	//if err != nil {
	//	log.Fatalf("创建 App 失败: %v", err)
	//}
	//fmt.Println("✅ AppToken:", m.appToken)
	//
	////第 2 步：创建数据表
	//m.tableID, err = m.createTable(m.appToken)
	//if err != nil {
	//	log.Fatalf("创建表失败: %v", err)
	//}
	//fmt.Println("✅ TableId:", m.tableID)
	//
	//// 第 3 步：批量插入 10 条随机数据
	//if err = m.insertMockRecords(m.tableID, 10); err != nil {
	//	log.Fatalf("插入初始数据失败: %v", err)
	//}
	//fmt.Println("✅ 初始数据插入完成")

	// 第 4 步：每隔 10 秒随机更新一条记录的“数量”字段
	fmt.Println("🚀 开始随机更新，每 10s 一次...")
	m.randomUpdates()
}

// createApp 调用 App.Create 接口，返回新建的 appToken
func (m *DemoManager) createApp() (string, error) {
	// 创建请求对象
	req := bitable.NewCreateAppReqBuilder().
		ReqApp(bitable.NewReqAppBuilder().
			Name(`一篇新的多维表格`).
			FolderToken(``).
			Build()).
		Build()

	// 发起请求
	resp, err := m.client.Bitable.V1.App.Create(m.ctx, req, larkcore.WithUserAccessToken(UserToken))
	if err != nil {
		return "", err
	}
	if !resp.Success() {
		return "", fmt.Errorf("App.Create 返回错误: %s", larkcore.Prettify(resp.CodeError))
	}
	return *resp.Data.App.AppToken, nil
}

// createTable 用 fundFields 创建一个新表，返回 tableId
func (m *DemoManager) createTable(appToken string) (string, error) {
	headers := make([]*bitable.AppTableCreateHeader, len(fundFields))
	for i, f := range fundFields {
		b := bitable.NewAppTableCreateHeaderBuilder().
			FieldName(f.Name).
			Type(f.Type).
			UiType(f.UIType)
		if len(f.Options) > 0 {
			opts := make([]*bitable.AppTableFieldPropertyOption, len(f.Options))
			for j, opt := range f.Options {
				opts[j] = bitable.NewAppTableFieldPropertyOptionBuilder().Name(opt).Color(j % 6).Build()
			}
			b.Property(
				bitable.NewAppTableFieldPropertyBuilder().Options(opts).Build(),
			)
		}
		headers[i] = b.Build()
	}

	resp, err := m.client.Bitable.V1.AppTable.Create(
		context.Background(),
		bitable.NewCreateAppTableReqBuilder().
			AppToken(appToken).
			Body(
				bitable.NewCreateAppTableReqBodyBuilder().
					Table(
						bitable.NewReqTableBuilder().
							Name("私募基金样例表").
							DefaultViewName("表格视图").
							Fields(headers).
							Build(),
					).
					Build(),
			).Build(),
	)
	if err != nil {
		return "", err
	}
	if !resp.Success() {
		return "", fmt.Errorf("创建表错误: %s", larkcore.Prettify(resp.CodeError))
	}
	return *resp.Data.TableId, nil
}

// insertMockRecords 插入 n 条随机 mock 数据到指定 table
func (m *DemoManager) insertMockRecords(tableID string, n int) error {
	for i := 0; i < n; i++ {
		fields := randomFundRecord()
		req := bitable.NewCreateAppTableRecordReqBuilder().
			AppToken(m.appToken).
			TableId(tableID).
			AppTableRecord(
				bitable.NewAppTableRecordBuilder().Fields(fields).Build(),
			).
			Build()

		resp, err := m.client.Bitable.V1.AppTableRecord.Create(
			context.Background(), req,
			larkcore.WithUserAccessToken(UserToken),
		)
		if err != nil {
			return err
		}
		if !resp.Success() {
			return fmt.Errorf("第 %d 条创建失败: %s", i+1, larkcore.Prettify(resp.CodeError))
		}
		fmt.Printf("第 %d 条创建成功: record_id=%s\n", i+1, *resp.Data.Record.RecordId)
		// 控制速率
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

// randomFundRecord 随机生成一条符合 fundFields 定义的记录
func randomFundRecord() map[string]interface{} {
	now := time.Now().UnixNano() / int64(time.Millisecond)
	record := make(map[string]interface{}, len(fundFields))
	for _, f := range fundFields {
		switch f.Type {
		case 1:
			record[f.Name] = fmt.Sprintf("%s-%d", f.Name, rand.Intn(1000))
		case 2:
			record[f.Name] = rand.Intn(10000)
		case 3:
			record[f.Name] = f.Options[rand.Intn(len(f.Options))]
		case 4:
			n := rand.Intn(len(f.Options)) + 1
			slice := make([]interface{}, n)
			for i := 0; i < n; i++ {
				slice[i] = f.Options[rand.Intn(len(f.Options))]
			}
			record[f.Name] = slice
		case 5:
			record[f.Name] = now - int64(rand.Intn(365*24*3600*1000))
		case 7:
			record[f.Name] = rand.Intn(2) == 0
		case 11:
			record[f.Name] = []interface{}{} // 可按需填充用户 id 列表
		default:
			record[f.Name] = nil
		}
	}
	return record
}

// randomUpdates 每隔 interval 从表中查询所有 record_id，随机更新“数量”字段
func (m *DemoManager) randomUpdates() {
	pageToken := ""
	var recordIDs []string

	// 2. 拉取所有 record_id
	for {
		items, nextToken, err := m.SearchRecords(viewID, nil, 20, pageToken)
		if err != nil {
			log.Fatalf("查询记录失败: %v", err)
		}
		for _, rec := range items {
			recordIDs = append(recordIDs, *rec.RecordId)
		}
		if nextToken == "" {
			break
		}
		pageToken = nextToken
	}

	// 准备可更新字段列表和候选值
	statuses := []string{"开放", "封闭"}
	fields := []string{"备注", "状态"}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// 3. 每隔 10s 随机更新一条记录的一个字段
	for range ticker.C {
		if len(recordIDs) == 0 {
			log.Println("没有可更新的记录，退出")
			return
		}
		// 随机选 record_id 和 field
		rid := recordIDs[rand.Intn(len(recordIDs))]
		field := fields[rand.Intn(len(fields))]

		// 随机生成一个新值
		var val interface{}
		switch field {
		case "数量":
			val = rand.Intn(1000)
		case "备注":
			val = fmt.Sprintf("随机备注@%d", time.Now().Unix())
		case "状态":
			val = statuses[rand.Intn(len(statuses))]
		}

		// 发起更新
		err := m.UpdateRecord(rid, map[string]interface{}{field: val})
		if err != nil {
			log.Printf("更新 record_id=%s 字段 %s 失败: %v\n", rid, field, err)
		} else {
			log.Printf("更新 record_id=%s 字段 %s=%v 成功\n", rid, field, val)
		}
	}

}

// SearchRecords 查询记录，支持分页、指定视图、字段名和过滤
func (m *DemoManager) SearchRecords(viewID string, fieldNames []string, pageSize int, pageToken string) ([]*bitable.AppTableRecord, string, error) {
	// 构建请求 body
	bodyBuilder := bitable.NewSearchAppTableRecordReqBodyBuilder().
		AutomaticFields(false)

	if viewID != "" {
		bodyBuilder = bodyBuilder.ViewId(viewID)
	}
	if len(fieldNames) > 0 {
		bodyBuilder = bodyBuilder.FieldNames(fieldNames)
	}
	body := bodyBuilder.Build()

	// 构建请求
	req := bitable.NewSearchAppTableRecordReqBuilder().
		AppToken(m.appToken).
		TableId(m.tableID).
		UserIdType("user_id").
		PageSize(pageSize).
		PageToken(pageToken).
		Body(body).
		Build()

	// 带上用户访问令牌进行查询
	resp, err := m.client.Bitable.V1.AppTableRecord.Search(
		m.ctx,
		req,
		larkcore.WithUserAccessToken(UserToken),
	)
	if err != nil {
		return nil, "", fmt.Errorf("调用 SearchRecords 失败: %w", err)
	}
	if !resp.Success() {
		return nil, "", fmt.Errorf("SearchRecords 返回错误: %s", larkcore.Prettify(resp.CodeError))
	}
	// 获取下一页 token
	nextToken := ""
	if resp.Data.PageToken != nil {
		nextToken = *resp.Data.PageToken
	}
	return resp.Data.Items, nextToken, nil
}

// UpdateRecord 根据 recordId 和变更字段打补丁
func (m *DemoManager) UpdateRecord(recordId string, fields map[string]interface{}) error {
	// 给日期类字段一个默认时间（如需）
	if ts, ok := fields["更新时间"]; ok && ts == nil {
		fields["更新时间"] = time.Now().UnixNano() / int64(time.Millisecond)
	}

	req := bitable.NewUpdateAppTableRecordReqBuilder().
		AppToken(m.appToken).
		TableId(m.tableID).
		RecordId(recordId).
		AppTableRecord(
			bitable.NewAppTableRecordBuilder().
				Fields(fields).
				Build(),
		).
		Build()

	resp, err := m.client.Bitable.V1.AppTableRecord.Update(m.ctx, req,
		larkcore.WithUserAccessToken(UserToken),
	)
	if err != nil {
		return fmt.Errorf("调用 Update API 失败: %w", err)
	}
	if !resp.Success() {
		return fmt.Errorf("Update 返回错误: %m", larkcore.Prettify(resp.CodeError))
	}
	return nil
}
