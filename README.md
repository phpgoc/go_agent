# 目录结构

- scripts里的内容一般只能linux环境下使用，一般是脚手架工具

- cmd go代码，可执行文件的入口

- command go代码，各个命令的实现

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