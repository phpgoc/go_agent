
# 目录结构
- scripts里的内容一般只能linux环境下使用，一般是脚手架工具

- cmd go代码，可执行文件的入口

- command go代码，各个命令的实现

- utils go代码，一些工具函数

- test 测试代码

- proto grpc的protoc文件

# 描述
- linux和windows 子目录文件里的函数声明不需要一致，因为子目录里只会被linux或者windows中的一个引用