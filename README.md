# snowflake

```bash
docker image pull registry.cn-shanghai.aliyuncs.com/yingzhor/snowflake:latest
```

### 安装

推荐使用`docker-compose`或`kubernetes`运行本服务。

* [docker-compose](.github/docker-compose.yaml)
* [kubernetes](.github/kubernetes.md)

### 使用

本程序只暴露了1个HTTP接口:

* [GET][/id]: ID生成接口，可选参数n。表示要一次性分配ID的数量。

### 客户端

* [Java](https://github.com/yingzhuo/snowflake-java-client)
* [Golang](https://github.com/yingzhuo/snowflake-golang-client)
