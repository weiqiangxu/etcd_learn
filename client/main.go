package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

const (
	// ServiceGoods 商品服务
	ServiceGoods = "goods"
)

// 全局服务锁
var serviceLocker = sync.Mutex{}

// servicePort 服务端口
var servicePort = 90

func main() {
	// 监听服务入口
	go func() {
		ServiceDiscovery()
	}()
	http.HandleFunc("/goods/list", GetGoodsList)
	_ = http.ListenAndServe(fmt.Sprintf(":%d", servicePort), nil)
}

// GetGoodsList 获取商品列表
func GetGoodsList(w http.ResponseWriter, r *http.Request) {
	client, endpoint, err := GetServiceEndpoint(ServiceGoods)
	if err != nil {
		_, _ = fmt.Fprint(w, err.Error())
		return
	}
	log.Printf("client=%s endpoint=%s\n", client, endpoint)
	url := fmt.Sprintf("http://%s/goods/list", endpoint)
	resp, err := http.Post(url, "application/json", http.NoBody)
	if err != nil {
		_, _ = fmt.Fprint(w, err.Error())
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	b, _ := io.ReadAll(resp.Body)
	_, _ = fmt.Fprint(w, string(b))
}
