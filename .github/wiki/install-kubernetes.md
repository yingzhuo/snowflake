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
