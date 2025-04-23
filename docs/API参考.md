# 飞书多维表格 CLI 工具 API 参考

本文档提供了飞书多维表格 CLI 工具的详细 API 参考，帮助开发者了解和使用项目中的各种功能。

## 目录

- [多维表格应用操作](#多维表格应用操作)
- [表格操作](#表格操作)
- [字段操作](#字段操作)
- [记录操作](#记录操作)
- [配置管理](#配置管理)

## 多维表格应用操作

### CreateApp

创建一个新的多维表格应用。

**函数签名：**

```go
func CreateApp(name, folderToken string) (*AppInfo, error)
```

**参数：**

- `name` (string): 多维表格的名称
- `folderToken` (string): 文件夹的 Token，如果为空则创建在根目录

**返回值：**

- `*AppInfo`: 包含创建的多维表格信息的结构体
- `error`: 错误信息，如果创建成功则为 nil

**示例：**

```go
info, err := bitable.CreateApp("项目管理表", "")
if err != nil {
    fmt.Printf("创建多维表格失败: %v\n", err)
    return
}
fmt.Printf("创建成功! 表格名称: %s, Token: %s\n", info.Name, info.AppToken)
```

**AppInfo 结构体：**

```go
type AppInfo struct {
    AppToken string // 多维表格应用的唯一标识
    Name     string // 多维表格的名称
    Url      string // 多维表格的访问链接
}
```

## 表格操作

### CreateTable

在指定的多维表格应用中创建一个新表格。

**函数签名：**

```go
func CreateTable(appToken string)
```

**参数：**

- `appToken` (string): 多维表格应用的 Token

**示例：**

```go
bitable.CreateTable("bascnCv7tnLBHJseKTqcP6Xvgzc")
```

**注意：**

此函数目前使用硬编码的方式创建一个包含「索引字段」和「产品状态」两个字段的表格。产品状态字段是单选类型，包含「上线」、「下线」和「开发中」三个选项。

## 字段操作

### CreateFields

创建字段（暂未实现完整功能）。

**函数签名：**

```go
func CreateFields()
```

**示例：**

```go
bitable.CreateFields()
```

**注意：**

此函数目前仅输出调试信息，尚未实现实际的字段创建功能。

## 记录操作

### InsertMock

插入模拟数据（暂未实现完整功能）。

**函数签名：**

```go
func InsertMock()
```

**示例：**

```go
bitable.InsertMock()
```

### QueryRecords

查询记录（暂未实现完整功能）。

**函数签名：**

```go
func QueryRecords()
```

**示例：**

```go
bitable.QueryRecords()
```

## 配置管理

### Load

加载配置文件。

**函数签名：**

```go
func Load(env string) (*Config, error)
```

**参数：**

- `env` (string): 环境名称，如 "local"、"dev"、"prod" 等

**返回值：**

- `*Config`: 配置结构体指针
- `error`: 错误信息，如果加载成功则为 nil

**示例：**

```go
config, err := config.Load("local")
if err != nil {
    fmt.Printf("配置加载失败: %v\n", err)
    return
}
```

**Config 结构体：**

```go
type Config struct {
    AppID           string `yaml:"app_id"`           // 飞书应用 ID
    AppSecret       string `yaml:"app_secret"`       // 飞书应用密钥
    AppToken        string `yaml:"app_token"`        // 多维表格应用 Token
    TableID         string `yaml:"table_id"`         // 表格 ID
    UserAccessToken string `yaml:"user_access_token"` // 用户访问令牌
    Debug           bool   `yaml:"debug"`           // 调试模式开关
}
```

## 终端用户界面

### StartTUI

启动终端用户界面。

**函数签名：**

```go
func StartTUI()
```

**示例：**

```go
tui.StartTUI()
```

### createBitableCmd

创建多维表格的命令。

**函数签名：**

```go
func createBitableCmd(name string, folderToken string) tea.Cmd
```

**参数：**

- `name` (string): 多维表格的名称
- `folderToken` (string): 文件夹的 Token

**返回值：**

- `tea.Cmd`: bubbletea 命令

## 错误处理

项目中的错误处理遵循 Go 语言的惯例，使用多返回值方式返回错误。API 函数通常会返回一个错误值作为最后一个返回参数，调用者应当检查这个错误值是否为 nil。

```go
info, err := bitable.CreateApp("测试表格", "")
if err != nil {
    // 处理错误
    fmt.Printf("创建多维表格失败: %v\n", err)
    return
}
// 继续处理成功的情况
```

## 最佳实践

1. **配置管理**：始终使用 `config.Load()` 函数加载配置，而不是直接读取配置文件。

2. **错误处理**：检查所有返回错误，并提供有意义的错误信息。

3. **资源清理**：使用 defer 语句确保资源被正确释放。

4. **日志记录**：在调试模式下使用 `utils.Logger.Debug()` 记录详细信息。

## 未来计划

以下 API 功能计划在未来版本中实现：

- 完整的字段管理 API
- 记录的增删改查 API
- 数据导入导出功能
- 批量操作支持

## 参考资源

- [飞书开放平台文档](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/bitable-overview)
- [飞书多维表格 API 文档](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app/introduction)
- [Go SDK 文档](https://github.com/larksuite/oapi-sdk-go)