# NEU_IPGW
东北大学校园网登录脚本

服务器需要在命令行中进行登录，在github中找到[东北大学非官方跨平台校园网关客户端](https://github.com/neucn/ipgw)这个项目，但是多次尝试总是无法登录，故而自行构建了一个简易版本。

# 用法

1. 准备go语言环境，可参考该[博客](https://www.superheaoz.top/2022/10/1036/)的2.3节。

2. 执行仓库中的`main.go`文件

    ```shell
    git clone https://github.com/DoraTiger/NEU_IPGW.git
    cd NEU_IPGW
    go run -mod=vendor ./cmd/app/main.go --username username --password password
    ```

3. 编译项目（可选）

    ```shell
    go build -mod=vendor ./cmd/app/main.go
    ```

# 参考

- [东北大学非官方跨平台校园网关客户端](https://github.com/neucn/ipgw)
- [NEU API](https://github.com/neucn/neugo)