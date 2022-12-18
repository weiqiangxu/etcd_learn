package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientV3 "go.etcd.io/etcd/client/v3"
	"log"
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
	Password:             "",
}

// serviceEndpointKeyPrefix 服务入口在 etcd 存储的 key 前缀
var serviceEndpointKeyPrefix = "/service_registry_discovery"

// ServiceDiscovery 服务发现
func ServiceDiscovery() {
	cli, err := clientV3.New(etcdCfg)
	if err != nil {
		log.Fatal(err)
	}
	for k, _ := range serviceEndpoints {
		go func(svc string) {
			ctx := context.Background()
			serviceKey := fmt.Sprintf("%s/%s", serviceEndpointKeyPrefix, svc)
			// 获取当前所有服务入口
			response, _ := cli.Get(ctx, serviceKey, clientV3.WithPrefix())
			serviceLocker.Lock()
			for _, v := range response.Kvs {
				serviceEndpoints[svc][string(v.Key)] = string(v.Value)
				log.Printf("key=%s endpoint=%s\n", string(v.Key), string(v.Value))
			}
			serviceLocker.Unlock()
			watchChan := cli.Watch(ctx, serviceKey, clientV3.WithPrefix(), clientV3.WithPrevKV())
			for item := range watchChan {
				for _, event := range item.Events {
					key := string(event.Kv.Key)
					endpoint := string(event.Kv.Value)
					switch event.Type {
					case mvccpb.PUT:
						serviceLocker.Lock()
						serviceEndpoints[svc][key] = endpoint
						serviceLocker.Unlock()
					case mvccpb.DELETE:
						serviceLocker.Lock()
						delete(serviceEndpoints[svc], key)
						serviceLocker.Unlock()
					}
				}
			}
		}(k)
	}
}
