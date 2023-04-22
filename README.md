# NEU_IPGW

东北大学校园网登录脚本

## 概述

服务器需要在命令行中进行登录，在 github 中找到 [neucn/ipgw](https://github.com/neucn/ipgw) 这个项目，但是多次尝试总是无法登录，故而基于该项目自行构建了一个简易版本。

学习go工程开发，在`v0.2.0`中，将项目基于`cobra`命令行工具进行了重构。

## 安装

以下安装过程以 Ubuntu 为例，其他系统请使用对应版本。

### 二进制文件安装

1. 从[release 页面](https://github.com/doratiger/neu_ipgw/releases)获取最新版本压缩文件
2. 解压缩，并赋予执行权限
3. 部署至`/usr/local/bin`目录

```bash
## example for ubuntu
# download
wget https://github.com/DoraTiger/NEU_IPGW/releases/download/v0.1.0/NEU_IPGW-linux-amd64.tar.gz
# unzip and grant
tar -zxf ./ipgw-linux-amd64.zip
chmod +x ./NEU_IPGW
# move
sudo cp ./NEU_IPGW  /usr/local/bin/
```

### 源码安装

1. 准备 go 语言环境，可参考该[博客](https://www.superheaoz.top/2022/10/1036/)的 2.3 节。
2. 编译项目(考虑到服务器本身无网络的情况，提供了vendor目录，支持离线编译，如不需要，请删除makefile中的`-mod=vendor`)
3. 部署至`/usr/local/bin`目录

```bash
## example for ubuntu
# download
git clone https://github.com/DoraTiger/NEU_IPGW.git
cd NEU_IPGW
# build
make all
# grant
chmod +x ./build/linux-amd64/NEU_IPGW
# move
sudo cp ./build/linux-amd64/NEU_IPGW /usr/local/bin/
```

## 使用

1. 登录校园网

```bash
NEU_IPGW login -u username -p password
```

2. 退出登录

```bash
NEU_IPGW logout
```

## 参考

- [东北大学非官方跨平台校园网关客户端](https://github.com/neucn/ipgw)
- [NEU API](https://github.com/neucn/neugo)

## 存在问题

