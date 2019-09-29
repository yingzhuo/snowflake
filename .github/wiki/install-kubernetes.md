### 启动

```bash
kubectl apply -f https://raw.githubusercontent.com/yingzhuo/snowflake/master/.github/kubernetes/snowflake.yaml
```

配置清单文件请参考[这里](https://github.com/yingzhuo/snowflake/blob/master/.github/kubernetes/snowflake.yaml)。

### 使用

查看暴露的服务:

```
$ kubectl --namespace=snowflake get service -o wide
NAME        TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)    AGE     SELECTOR
snowflake   ClusterIP   10.99.116.153   <none>        8080/TCP   8m38s   app=snowflake
```

测试该服务:

```
$ curl -X'GET' "http://10.99.116.153:8080/id?n=1"
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
--response-type | json           | 程序response的数据格式 (json | protobuf)

**注意: 当部署`Snowflake`集群时需要为不同节点指定不同的`NodeID`防止生成的ID冲突。**