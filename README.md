# NEU_IPGW

东北大学校园网登录脚本

## 概述

服务器需要在命令行中进行登录，在 github 中找到 [neucn/ipgw](https://github.com/neucn/ipgw) 这个项目，但是多次尝试总是无法登录，故而基于该项目自行构建了一个简易版本。

学习 go 工程开发，在`v0.2.0`中，将项目基于`cobra`命令行工具进行了重构。

## 安装

### 二进制文件安装

1. 从[release 页面](https://github.com/doratiger/neu_ipgw/releases)获取匹配你平台的压缩文件
2. 验证文件完整性 (optional)
3. 解压缩，并赋予执行权限
4. 选择以下方式之一放置可执行文件：
	- **自动安装（Linux/macOS/FreeBSD）**：

```bash
# 将当前目录下的 NEU_IPGW 安装到 ~/.local/share/... 并链接 ~/.local/bin
./NEU_IPGW install

# 安装到 /usr/local/bin（需要 sudo）
sudo ./NEU_IPGW install --system
```

	- **手动拷贝（所有平台通用）**：

```bash
## example for linux(x86_64)
# download
wget https://github.com/DoraTiger/NEU_IPGW/releases/latest/NEU_IPGW-linux-amd64.tar.gz

# check (optional)
wget https://github.com/DoraTiger/NEU_IPGW/releases/latest/NEU_IPGW-linux-amd64.tar.gz.sha256
sha256sum -c NEU_IPGW-linux-amd64.tar.gz.sha256

# unzip and grant
tar -zxf ./NEU_IPGW-linux-amd64.tar.gz
chmod +x ./NEU_IPGW

# move (optional)
sudo cp ./NEU_IPGW  /usr/local/bin/
```

> 自动安装会在不支持的平台（如 Windows）中止，此时请使用手动方式。请确保 `~/.local/bin` 已加入 PATH。

### 源码安装

源码安装方式自适应系统架构

1. 准备 go 语言环境，可参考该[博客](https://www.superheaoz.top/2022/10/1036/)的 2.3 节
2. 获取源码并编译
3. 在 Linux/macOS/FreeBSD 上可使用 Makefile 安装目标，其他平台手动拷贝即可

```bash
# download
git clone https://github.com/DoraTiger/NEU_IPGW.git
cd NEU_IPGW
# build
make build
# 自动安装到用户目录（~/.local/share/... + ~/.local/bin）
make install-user

# 安装到系统目录（需要 sudo）
sudo make install-system

# 移除用户级符号链接
make uninstall-user

# grant
chmod +x ./build/NEU_IPGW
# move (optional)
sudo cp ./build/NEU_IPGW /usr/local/bin/
```

## 使用

1. 登录校园网

```bash
NEU_IPGW login -u username -p password
```

登录成功后会自动展示账号、已用流量、时长、余额以及当前在线 IP，便于确认网关状态。

2. 保存账号密码到本地（同名自动覆盖）

```bash
NEU_IPGW login -u username -p password --save
```

3. 使用已保存账号登录

```bash
# 使用最近一次成功账号
NEU_IPGW login

# 指定某个已保存账号
NEU_IPGW login --account username
```

4. 查看当前登录信息

```bash
NEU_IPGW info
```

若当前终端尚未登录，命令会提示先执行 `NEU_IPGW login`。

默认输出中文，可通过 `--lang en` 或环境变量 `NEU_IPGW_LANG=en` 切换到英文：

```bash
NEU_IPGW --lang en info

export NEU_IPGW_LANG=en
NEU_IPGW info
```

5. 查看已保存账号列表

```bash
NEU_IPGW accounts list
```

6. 退出登录

```bash
NEU_IPGW logout
```

7. 退出登录并删除本地凭据

```bash
# 删除最近一次成功账号的本地凭据
NEU_IPGW logout --forget

# 删除指定账号的本地凭据
NEU_IPGW logout --forget --account username
```

8. 查询电费

```bash
# 使用已保存账号查询电费
NEU_IPGW power

# 手动指定宿舍号查询
NEU_IPGW power -r <宿舍编号>

# 只查询宿舍信息（不查电费）
NEU_IPGW power -i

# 输出全量信息（含房间描述、读表时间等）
NEU_IPGW power -a

# 手动指定宿舍号 + 全量输出
NEU_IPGW power -r <宿舍编号> -a

# 使用指定账号查询
NEU_IPGW power --account username

# 直接传入凭据查询
NEU_IPGW power -u username -p password

# 传入凭据并保存
NEU_IPGW power -u username -p password --save
```

电费输出统一格式：剩余度数 + 电费余额（按 0.51 元/度换算）。

宿舍编号规则：宿舍楼编号(3位) + 寝室号(4位，不足4位前面补0)。

宿舍编号对照表详见：[东北大学各宿舍电费缴费编号](https://hqglc.neu.edu.cn/2021/1020/c2465a204768/page.htm)

9. 环境变量覆盖

```bash
# 覆盖本地凭据目录
export NEU_IPGW_CONFIG_DIR="$HOME/.config/doratiger/neu-ipgw"

# 覆盖本地加密密钥
export NEU_IPGW_MASTER_KEY="your-custom-master-key"

# 覆盖命令行输出语言（zh / en，默认 zh）
export NEU_IPGW_LANG="en"
```

本地凭据默认目录为 `~/.config/doratiger/neu-ipgw/`。
默认内置密钥是公开值，仅用于降低凭据被直接阅读的风险，不等价于高强度安全；建议自行生成并覆盖 `NEU_IPGW_MASTER_KEY`。
如果你是自行编译，可以直接修改 [config/config.go](config/config.go) 中的 `DefaultMasterKey` 再编译。

## 更新日志

[简体中文](docs/CHANGELOG.md) | [English](docs/CHANGELOG.en.md)


## 参考

- [东北大学非官方跨平台校园网关客户端](https://github.com/neucn/ipgw)
- [NEU API](https://github.com/neucn/neugo)

## 存在问题
