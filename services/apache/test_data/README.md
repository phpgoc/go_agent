##

- 在检验配置文件是否正确时，可以使用`-t`选项
- 下边的命令是在本地linux环境已经安装了apache2的情况下，使用的命令
- 仅仅是检验配置文件是否正确，apache配置文件非常复杂,mod都是动态加载的,进行-t都是基于主配置文件进行追加检验的,也就是说listen 80因为在主配置文件里已有,如果测试文件里有listen 80,则会报错,因为端口已经被占用
- -t是加载主配置文件的,测试用例是单纯使用该配置文件

```shell
source /etc/apache2/envvars
 apache2 -t -c "Include `pwd`/simple.config"
```

```shell
source /etc/apache2/envvars
export INCLUDE_PATH=`pwd`/include
 apache2 -t -c "Include `pwd`/include/apache.config"
```

```shell

source /etc/apache2/envvars
apache2 -t -c "Include `pwd`/inlog_outlog/apache.config"

```

```shell
source /etc/apache2/envvars
export TEST_ROOT=`pwd`/double_include
 apache2 -t -c "Include `pwd`/double_include/apache.config"

```