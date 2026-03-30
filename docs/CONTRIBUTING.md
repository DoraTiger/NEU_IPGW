# 贡献指南

[简体中文](./CONTRIBUTING.md) | [English](.//CONTRIBUTING.en.md)

感谢你对本项目的关注与贡献！为保证代码质量和协作效率，请在提交 Issue 或 Pull Request 前阅读本指南。

## 如何贡献

### 🐞 报告问题

- 在 [Issues](https://github.com/doratiger/neu_ipgw/issues) 中搜索是否已有类似问题。
- 提供清晰的复现步骤、环境信息 (如 Go 版本、操作系统) 和预期/实际行为。

### 💡 建议功能

- 先通过 Issue 讨论新功能的必要性和设计，避免无效开发。
- 描述清楚使用场景和预期行为。

### 🧑‍💻 贡献代码

1. Fork 本仓库
2. 创建特性分支 (如 `feat/qr-login` 或 `fix/timeout-issue`)
3. 编写代码并确保通过测试
4. 遵循下方的 **代码规范** 和 **Commit Message 格式**
5. 提交 Pull Request，并关联相关 Issue(如有)

## 代码规范

### Go 代码风格

本项目遵循 [Go 官方代码审查建议](https://go.dev/wiki/CodeReviewComments) 和 [Uber Go 风格指南](https://github.com/xxjwxc/uber_go_guide_cn)。  
请确保：
- 使用 `go fmt` 格式化代码
- 避免不必要的复杂逻辑
- 函数和变量命名清晰、符合语义

### 中文文案排版

所有中文文档 (包括注释、README、CHANGELOG 等) 请遵循  
[《中文文案排版指北》](https://github.com/sparanoid/chinese-copywriting-guidelines)：
- 中英文/数字之间加空格 (如 `支持 iOS 15`)
- 使用全角中文标点
- 不使用非标准缩写 (如“前端”不要写成“FED”)

## Commit Message 格式

我们采用 [Conventional Commits](https://www.conventionalcommits.org/) 规范，格式如下：

```text
<type>: <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

### `type` 类型说明

-   feat：新增功能
-   fix：修复 Bug
-   docs：文档更新 (含 README、CHANGELOG 等)
-   style：代码样式调整 (如格式化、空格等，不涉及逻辑变更)
-   refactor：代码重构 (不涉及功能变更)
-   perf：性能优化
-   test：增加或修改测试用例
-   chore：修改编译流程，或变更依赖库和工具等
-   ci：持续集成配置变更
-   revert：版本回滚

### `subject` 要求

-   简洁的标题，描述本次提交的概要
-   使用**祈使句**、**现在时** (例如 "fix bug" 而不是 "Fixed bug")
-   首字母**小写**
-   **不加句号**
-   长度建议 <= 50 字符

### `body` 要求 (可选但推荐)

-   解释**为什么**要做这些更改，而不是只描述做了什么
-   每行 <= 72 字符
-   `feat`/`fix`/`refactor` 等建议提供 body

### `footer` 要求 (可选)

-   关联的 Issue 或 PR (如 `Close #123`)

## 示例

1. `feat`/`refactor` 等提交示例

```text
feat(auth): add QR code login

- implement ScanQR in internal/client
- add --qr flag to CLI
- update README usage example

Closes #45
```

2. `docs` 提交示例

```text
docs: update README with logout command
```

3. `fix` 提交示例

```text
fix(client): handle timeout in weak network

Previous implementation would hang indefinitely.
Now use context.WithTimeout to enforce 10s limit.

Closes #38
```

## 文档与测试

-   **新功能必须更新 `README.md`**
-   **用户可见变更需同步更新 `CHANGELOG.md`**
-   **破坏性变更请在 `UPDATING.md` **中说明**
-   鼓励为关键逻辑添加单元测试 (*_test.go)

## 许可证

所有贡献均视为同意在本项目的 [MIT 许可证](./LICENSE) 下发布

## 获取帮助

如有疑问，请在 [Discussions](https://github.com/DoraTiger/NEU_IPGW/discussions) 或相关 Issue 中提出
