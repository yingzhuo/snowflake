# Snowflake

```bash
docker image pull registry.cn-shanghai.aliyuncs.com/yingzhor/snowflake:latest
```

### 安装

推荐使用`docker-compose`或`kubernetes`运行本服务。

* [docker-compose](.github/wiki/install-dco.md)
* [kubernetes](.github/wiki/install-kubernetes.md)

如果你愿意用`systemctl`管理之，你可以参考如下的`snowflake.service`

```
[Unit]
Description=Snowflake ID generator
After=network.target

[Service]
Type=simple
ExecStart=/opt/snowflake/bin/snowflake --type=json --indent --port=15656 --node-id=512
KillMode=mixed
Restart=on-failure

[Install]
WantedBy=multi-user.target
```

### 使用

本程序只暴露了3个HTTP接口:

* [GET][/id]: ID生成接口，可选参数n。表示要一次性分配ID的数量。
* [GET][/ping]: 简单返回"pong"，只作为`k8s`的 readiness-probe 使用。
* [GET][/metrics]: 只作为 prometheus 监控采集数据的接口使用。

### 客户端

* [Java](https://github.com/yingzhuo/snowflake-java-client)
* [Golang](https://github.com/yingzhuo/snowflake-golang-client)

### 参考

* [https://github.com/yingzhuo/go-cli](https://github.com/yingzhuo/go-cli)
* [https://github.com/bwmarrin/snowflake](https://github.com/bwmarrin/snowflake)
* [https://github.com/golang/protobuf](https://github.com/golang/protobuf)
* [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* [https://github.com/prometheus/client_golang](https://github.com/prometheus/client_golang)
