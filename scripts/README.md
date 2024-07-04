# 脚本概述

- 脚本都是用于开发的辅助脚本
- 只能在 Linux 系统上运行
- private开头的脚本是工具脚本，不要直接调用
- 脚本不一定能够直接使用，需要自己安装依赖，比如docker，docker-buildx，protobuf-compiler等
- 脚本如果有参数，可以通过-h查看帮助,如果没有-h参数说明不需要参数

# 特定脚本说明

- [build_go_env.sh](build_go_env.sh): docker构建go环境
  - 如果go.mod文件有变化，最好重新执行一次

- [gen_go_by_protoc.sh](gen_go_by_protoc.sh) 
   * 需要本机[build_go_env.sh](build_go_env.sh)执行成功过
   * 生成proto文件对应的go代码
   * 可以使用-i参数 执行只执行某些/个proto文件 使用grep匹配的 -i net 会匹配到network.proto

- [build_dotnet_env.sh](build_dotnet_env.sh): docker构建dotnet环境

- [build_go_agent_exe.sh](build_go_agent_exe.sh): docker生成最终目标exe，生成在bin目录

- [run_dotnet_test.sh](run_dotnet_test.sh) docker运行dotnet测试
    - 默认会调用docker宿主机上的agent，如果要改，手动修改 [dotnet_tests/GrpcTest/Utils.cs](../dotnet_tests/GrpcTest/Utils.cs)
  
- [run_go_unit_test.sh](run_go_unit_test.sh) docker运行go单元测试

- [run_in_go_env_it.sh](run_in_go_env_it.sh) 交互式进入go环境
  * -p 参数可以指定端口映射，比如 -p 50051 则会映射50051:50051
  * 可以手动执行一些生成动作
  * CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/agent_windows_amd64.exe ./cmd/go-agent/main.go

