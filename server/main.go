package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// ServiceGoods 商品服务
	ServiceGoods = "goods"
)

// serviceEndpointKeyPrefix 服务入口在 etcd 存储的 key 前缀
var serviceEndpointKeyPrefix = "/service_registry_discovery"

// servicePort 服务端口
var servicePort = 80

func main() {
	InitEndpoint()
	// 服务注册
	go func() {
		ServiceRegistry()
	}()
	http.HandleFunc("/goods/list", GetGoodsList)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil)
}

// GetGoodsList 获取商品列表
func GetGoodsList(w http.ResponseWriter, r *http.Request) {
	var res = map[string]interface{}{
		"message": "get goods list success",
	}
	b, _ := json.Marshal(res)
	_, _ = fmt.Fprint(w, string(b))
}
