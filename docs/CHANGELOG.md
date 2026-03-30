# 更新日志

[简体中文](./CHANGELOG.md) | [English](./CHANGELOG.en.md)

## [v0.3.0] - 2026-03-30

### 新增
- 添加本地加密凭据存储并支持 `--save`
- 添加按用户名的多账号管理与 `accounts list`
- 登录/登出命令新增 `--account` 凭据选择
- 允许通过环境变量覆盖凭据目录与主密钥
- 新增 `install` 子命令，支持在 Linux/macOS/FreeBSD 进行用户级或系统级安装
- Makefile 新增 `install-user`、`install-system`、`uninstall-user` 目标并遵循 XDG 目录布局

### 更改
- 扩展 `Makefile` 的跨平台构建矩阵
- 更新文档，描述 Linux 用户级安装目录与符号链接
- 重组 README 安装章节，按平台覆盖自动/手动流程

## [v0.2.2] - 2026-03-10

### 新增
- 引入 RSA 加密机制以实现新的账户认证机制

### 更改
- 调整账号数据结构以兼容加密凭据
- 更新配置模块以兼容加密凭据

## [v0.2.1] - 2025-09-25

### 更改
- 优化构建流程并改进 `Makefile` 配置
- 在发布流程中加入 SHA256 校验生成

### 新增
- 发布脚本自动生成二进制文件的 SHA256 校验值
- 为每个版本提供统一的 `checksums.txt`

## [v0.2.0] - 2023-04-22

### 更改
- **破坏性变更**：CLI 切换为子命令模式（`login`、`logout`、`version`）

### 新增
- 支持 `logout` 命令
- `login` 增加 `-u`（用户名）与 `-p`（密码）
- `version` 增加 `--verbose` 以输出详细信息

### 移除
- 移除根命令上的扁平化参数（`--username`、`--password`）

## [v0.1.0] - 2023-04-18

### 新增
- 初始发布
- 通过 `--username` 与 `--password` 登录
- 通过 `--version` 查看版本信息
