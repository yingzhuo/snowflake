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

### 客户端

* [Java](https://github.com/yingzhuo/snowflake-java-client)
* [Golang](https://github.com/yingzhuo/snowflake-golang-client)

### 参考

* [https://github.com/bwmarrin/snowflake](https://github.com/bwmarrin/snowflake)
* [https://github.com/golang/protobuf](https://github.com/golang/protobuf)
* [https://github.com/sirupsen/logrus](https://github.com/sirupsen/logrus)
* [https://github.com/subchen/go-cli](https://github.com/subchen/go-cli)
