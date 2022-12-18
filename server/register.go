package main

import (
	"context"
	"encoding/json"
	"fmt"
	clientV3 "go.etcd.io/etcd/client/v3"
	"log"
	"os"
	"time"
)

// etcdCfg Etcd配置
var etcdCfg = clientV3.Config{
	Endpoints: []string{
		"http://127.0.0.1:2379",
	},
	DialTimeout:          time.Second * 30,
	DialKeepAliveTimeout: time.Second * 30,
	Username:             "root",
	Password:             "123456",
}

// ServiceRegistry 服务注册
func ServiceRegistry() {
	hostname, _ := os.Hostname()
	cli, err := clientV3.New(etcdCfg)
	if err != nil {
		log.Fatal(err)
	}
	key := fmt.Sprintf("%s/%s/%s", serviceEndpointKeyPrefix, ServiceGoods, hostname)
	log.Printf("key=%s\n", key)
	ctx := context.Background()
	// 过期时间: 3秒钟
	ttl := 3
	// 创建租约
	lease, err := cli.Grant(ctx, int64(ttl))
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(lease)
	log.Printf("grant lease suucess: %s\n", string(b))
	// put kv
	res, err := cli.Put(ctx, key, endpoint, clientV3.WithLease(lease.ID))
	if err != nil {
		log.Fatal(err)
	}
	b, _ = json.Marshal(res)
	log.Printf("put kv with lease suucess: %s\n", string(b))
	// 保持租约不过期
	klRes, err := cli.KeepAlive(ctx, lease.ID)
	if err != nil {
		log.Fatal(err)
	}
	// 监听续约情况
	for v := range klRes {
		b, _ = json.Marshal(v)
		log.Printf("keep lease alive suucess: %s\n", string(b))
	}
	log.Println("stop keeping lease alive")
}
