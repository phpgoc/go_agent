# 目录结构

- [scripts](scripts/README.md) 里的内容一般只能linux环境下使用，一般是脚手架工具

- cmd go代码，可执行文件的入口

- services go代码，各个命令的实现

- utils go代码，一些工具函数

- dotnet_tests dotnet的测试代码

- proto grpc的protoc文件
- 
- bin 编译目标路径
- 
- release 发布二进制文件夹

# 描述

- 如果文件夹名命名为linux或者windows，这两个文件夹里的函数声明无须保持一致，因为他们只会被同架构的文件调用
- 如果同一个文件夹里命名为linux.go, windows.go, 那么这两个文件里的函数声明必须保持一致
- 类似 hello_linux.go, hello_windows.go同上

# 建立开发环境

1. 安装(docker,docker-buildx) or (podman)
2. <div id="step2">执行scripts/build-go_env.sh</div>

    - 可能遇到docker hub网络问题
    - 下载google的go包依赖可能需要翻墙
    - 可以通过传递容器 image的方式完成这步
3. <div id="step3">执行scripts/build_dotnet_env.sh</div>

    - dotnet是测试用的, 不是必须执行
    - 可能遇到docker hub网络问题
    - 可以通过传递容器 image的方式完成这步

# 编译

1. 需要预先完成 [建立开发环境 2.](#step2)
2. 执行 scripts/build_go_agent_exe.sh
    - 生成的文件在bin目录下.默认是 linux amd64 版本的可执行文件
3. 要生成其他的目标可以执行执行 scripts/run_in_go_env_it.sh 进入交互式环境
    - GOOS=windows GOARCH=amd64 go build -o bin/agent_windows_amd64.exe ./cmd/go-agent/main.go
    - CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/agent_darwin_amd64 ./cmd/go-agent/main.go
    - CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/agent_darwin_arm64 ./cmd/go-agent/main.go
    - CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/agent_linux_arm64 ./cmd/go-agent/main.go
    - CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bin/agent_linux_arm ./cmd/go-agent/main.go
    - CGO_ENABLED=0代表不使用cgo, 也就是不使用c语言的库,但是对windows来说,这个参数是没用的,因为windows一定会使用cgo,并且不太可能找不到
4. 这几乎就是最好的到处编译的方式了,因为国内网络问题,go依赖包不好下载,把依赖放到项目里不如这种docker image 提前download的方式

# 开发

1. 在protos/目录下编写proto文件
2. 使用scripts/gen_go_by_protoc.sh生成go代码(必须,统一protoc,protoc-go等的版本)
3. 在services/目录下编写服务实现
4. 在dotnet_tests/GrpcLib/目录下编写dotnet 调用方法
5. 在dotnet_tests/CallEverything/目录下编写方法调用
6. 执行go编译目标运行.或者映射端口到宿主机的方式50051:50051,在go_env 里执行 
   - scripts/run_in_go_env_it.sh -p 50051
   - go run cmd/go-agent/main.go
7. 执行dotnet测试 scripts/run_dotnet_test.sh
   - 需要预先完成 [建立开发环境 3.](#step3)
   - docker caller 的地址需要是172.17.0.1
   - podman caller 的地址需要是localhost
   - [dotnet_tests/CallEverything/Program.cs](dotnet_tests/CallEverything/Program.cs)
   - 机器的第一次执行可能会较慢,需要下载dotnet的依赖包
   - dotnet的依赖下载一般不会有问题