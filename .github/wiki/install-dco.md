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
