# 脚本概述

- 脚本都是用于开发的辅助脚本
- 只能在 Linux 系统上运行
- private开头的脚本是工具脚本，不要直接调用
- 脚本不一定能够直接使用，需要自己安装依赖，比如docker，docker-buildx，protobuf-compiler等

# 特定脚本说明

- build_go_env.sh: docker构建go环境

- gen_go_by_protoc.sh 通过protos目录下的proto文件生成go代码，脚本还没有docker化，需要本地安装protobuf-compiler，protoc-gen-go，protoc-gen-go-grpc

- build_dotnet_env.sh: docker构建dotnet环境

- build_go_agent_eve.sh: docker生成最终目标exe，生成在bin目录

- run_dotnet_test.sh docker运行dotnet测试
    - 默认会调用docker宿主机上的agent，如果要改，手动修改 [dotnet_tests/GrpcTest/Utils.cs](../dotnet_tests/GrpcTest/Utils.cs)
  
- run_go_unit_test.sh docker运行go单元测试



