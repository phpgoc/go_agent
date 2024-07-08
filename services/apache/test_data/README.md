##

- 在检验配置文件是否正确时，可以使用`-t`选项
```shell
source /etc/apache2/envvars
 apache2 -t -c "Include `pwd`/simple.config"
```
