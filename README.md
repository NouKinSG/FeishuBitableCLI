# FeishuBitableCLI

[![Go Version](https://img.shields.io/badge/Go-1.23-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

FeishuBitableCLI 是一个基于飞书多维表格 API 的交互式终端管理工具，帮助用户在命令行环境中高效管理飞书多维表格。

## 功能特点

- 🚀 **交互式终端界面**：基于 [bubbletea](https://github.com/charmbracelet/bubbletea) 构建的美观易用的 TUI 界面
- 📊 **多维表格管理**：创建、查看和管理飞书多维表格
- 🔧 **表格结构操作**：管理字段、记录和表格结构
- 🔄 **数据同步**：支持本地数据与飞书多维表格的同步操作
- 🔐 **安全认证**：基于飞书开放平台的安全认证机制

## 安装说明

### 前置条件

- Go 1.23 或更高版本
- 飞书开发者账号和应用凭证

### 安装步骤

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

在项目根目录创建 `configs/local.yaml` 文件，填入以下内容：

```yaml
app_id: "你的飞书应用ID"
app_secret: "你的飞书应用密钥"
user_access_token: "你的用户访问令牌"
```

## 使用方法

### 启动应用

```bash
./feishu-bitable-cli
```

### 主要功能

- **创建多维表格**：通过交互式界面创建新的多维表格
- **查看多维表格**：浏览已有的多维表格（开发中）
- **删除多维表格**：删除不需要的多维表格（开发中）
- **管理表格结构**：添加、修改字段和记录

## 项目结构

```
.
├── cmd/                    # 命令行入口
├── configs/                # 配置文件
├── internal/               # 内部包
│   ├── bitable/           # 飞书多维表格 API 封装
│   ├── cli/               # 命令行界面
│   ├── config/            # 配置管理
│   └── utils/             # 工具函数
└── pkg/                    # 公共包
    └── bitable/           # 多维表格客户端
```

## 技术栈

- [Go](https://golang.org/) - 编程语言
- [bubbletea](https://github.com/charmbracelet/bubbletea) - 终端 UI 框架
- [飞书开放平台 API](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/bitable-overview) - 多维表格 API

## 贡献指南

1. Fork 本仓库
2. 创建您的特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交您的更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 打开一个 Pull Request

## 许可证

本项目采用 MIT 许可证 - 详见 [LICENSE](LICENSE) 文件

## 联系方式

如有问题或建议，请提交 [Issue](https://github.com/yourusername/FeishuBitableCLI/issues)
