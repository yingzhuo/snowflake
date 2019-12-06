### 启动

编辑文件`docker-compose.yml`，然后启动之。

```yaml
version: "3.7"

services:
  snowflake:
    image: "registry.cn-shanghai.aliyuncs.com/yingzhor/snowflake:latest"
    container_name: "snowflake"
    restart: "always"
    ports:
      - "8080:8080"
    environment:
      - "SNOWFLAKE_TYPE=json"
      - "SNOWFLAKE_INDENT=yes"
```

```
$ docker-compose up -d
```

### 使用

查看暴露的服务:

```
$ docker-compose ps
  Name                 Command               State           Ports
---------------------------------------------------------------------------
snowflake   snowflake --host=0.0.0.0 - ...   Up      0.0.0.0:8080->8080/tcp
```

测试该服务:

```
$ curl -X'GET' "http://localhost:8080/id?n=1"
[1171555234682916864]
```

**注意: 其中请求参数n为一次生成的ID数量，默认为1。**

### 定制

镜像`quay.io/yingzhuo/snowflake:latest`启动为容器时，可通过可执行文件`/bin/snowflake`的命令行选项设置。

命令行选项        | 默认值          | 意义
----------------|----------------|-----------------------------------------------
--node-id       | 512            | 程序使用的NodeID (0 ~ 1023)
--host          | 0.0.0.0        | 程序允许访问的IP地址
--port          | 8080           | 程序监听的端口号
--response-type | protobuf       | 程序response的数据格式 (protobuf | json)

**注意: 当部署`Snowflake`集群时需要为不同节点指定不同的`NodeID`防止生成的ID冲突。**
