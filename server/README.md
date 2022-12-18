# server

1. 启动一个http服务
2. 服务注册

### run etcd in docker [arm不支持]

[https://etcd.io/docs/v2.3/docker_guide/](https://etcd.io/docs/v2.3/docker_guide/)

[https://hub.docker.com/r/bitnami/etcd](https://hub.docker.com/r/bitnami/etcd/)


### 安装在本机

```
https://github.com/etcd-io/etcd/releases/tag/v3.5.6

https://etcd.io/docs/v3.5/quickstart/
```

### 运行

> ./etcd

```
{"level":"info","ts":"2022-12-18T23:13:30.027+0800",
"caller":"embed/etcd.go:375","msg":"closed etcd server",
"name":"default","data-dir":"default.etcd",
"advertise-peer-urls":["http://localhost:2380"],
"advertise-client-urls":["http://localhost:2379"]}
```

```
➜  etcd-v3.5.6-darwin-arm64 ./etcdctl put greeting "Hello, etcd"
OK
➜  etcd-v3.5.6-darwin-arm64 etcdctl get greeting
greeting
Hello, etcd
➜  etcd-v3.5.6-darwin-arm64 
```

