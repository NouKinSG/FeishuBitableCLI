# 1. 飞书多维表格 CLI 工具

[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## 1.1. 项目概述

飞书多维表格 CLI 是一个基于飞书多维表格 API 的命令行工具，提供交互式终端界面，帮助用户在命令行环境中高效管理飞书多维表格。该工具使用 Go 语言开发，基于 [bubbletea](https://github.com/charmbracelet/bubbletea) 构建美观易用的 TUI 界面。

## 1.2. 功能特点

- 🚀 **交互式终端界面**：基于 bubbletea 构建的美观易用的 TUI 界面
- 📊 **多维表格管理**：创建、查看和管理飞书多维表格
- 🔧 **表格结构操作**：管理字段、记录和表格结构
- 🔄 **数据同步**：支持本地数据与飞书多维表格的同步操作
- 🔐 **安全认证**：基于飞书开放平台的安全认证机制

## 1.3. 安装说明

### 1.3.1. 前置条件

- Go 1.23 或更高版本
- 飞书开发者账号和应用凭证

### 1.3.2. 安装步骤

1. 克隆仓库

```bash
git clone https://github.com/yourusername/FeishuBitableCLI.git
cd FeishuBitableCLI
```

2. 安装依赖

```bash
go mod download
```

3. 编译项目

```bash
go build -o feishu-bitable-cli ./cmd/feishu-bitable-cli
```

4. 配置应用凭证

在项目根目录创建 `configs/config.local.yaml` 文件，填入以下内容：

```yaml
app_id: "你的飞书应用ID"
app_secret: "你的飞书应用密钥"
user_access_token: "你的用户访问令牌"
debug: false
```

## 1.4. 使用指南

### 1.4.1. 启动应用

```bash
./feishu-bitable-cli
```

### 1.4.2. 基本操作

启动后，您将看到交互式终端界面，可以使用以下功能：

- 创建多维表格
- 管理表格结构
- 操作字段和记录
- 数据导入导出

## 1.5. API 参考

### 1.5.1. 多维表格操作

#### 1.5.1.1. 创建多维表格

```go
info, err := bitable.CreateApp("表格名称", "文件夹Token")
```

#### 1.5.1.2. 创建表格

```go
bitable.CreateTable("应用Token")
```

### 1.5.2. 字段操作

```go
bitable.CreateFields()
```

### 1.5.3. 记录操作

```go
// 插入模拟数据
bitable.InsertMock()

// 查询记录
bitable.QueryRecords()
```

## 1.6. 配置说明

配置文件位于 `configs/config.local.yaml`，包含以下字段：

| 字段 | 说明 |
|------|------|
| app_id | 飞书应用 ID |
| app_secret | 飞书应用密钥 |
| app_token | 多维表格应用 Token |
| table_id | 表格 ID |
| user_access_token | 用户访问令牌 |
| debug | 调试模式开关 |

## 1.7. 项目结构

```
.
├── cmd/                    # 命令行入口
│   └── feishu-bitable-cli/ # 主程序
├── configs/                # 配置文件目录
├── internal/               # 内部包
│   ├── bitable/           # 多维表格操作相关
│   │   ├── app.go         # 应用操作
│   │   ├── field.go       # 字段操作
│   │   ├── record.go      # 记录操作
│   │   └── table.go       # 表格操作
│   ├── cli/               # 命令行界面
│   │   └── tui/           # 终端用户界面
│   ├── config/            # 配置管理
│   └── utils/             # 工具函数
├── pkg/                    # 公共包
│   └── bitable/           # 多维表格客户端
└── test/                   # 测试代码
```

## 1.8. 开发指南

### 1.8.1. 添加新功能

1. 在 `internal/bitable/` 目录下实现相关 API 调用
2. 在 `internal/cli/tui/` 目录下添加用户界面处理逻辑
3. 更新配置和文档

### 1.8.2. 调试技巧

启用调试模式：

```yaml
# configs/config.local.yaml
debug: true
```

## 1.9. 常见问题

### 1.9.1. 认证失败

- 检查 `app_id` 和 `app_secret` 是否正确
- 确认 `user_access_token` 是否有效且未过期

### 1.9.2. 权限问题

- 确保应用已获得必要的权限范围
- 检查用户是否有操作目标表格的权限

## 1.10. 贡献指南

欢迎贡献代码、报告问题或提出改进建议。请遵循以下步骤：

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 1.11. 许可证

本项目采用 MIT 许可证 - 详情请参阅 [LICENSE](LICENSE) 文件

## 1.12. 联系方式

如有问题或建议，请通过 Issues 或以下方式联系我们：

- 邮箱：your.email@example.com
- 项目主页：https://github.com/yourusername/FeishuBitableCLI