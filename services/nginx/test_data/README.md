# 

1. 测试数据
2. 每个文件夹是一个单独的测试用例
3. 可以到那个路径执行 ```nginx -c nginx.conf -p . -t ``` 来测试配置文件是否正确
4. 这些文件都会被 [../get_nginx_info_test.go](../get_nginx_info_test.go) 读取