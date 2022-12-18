package main

import (
	"errors"
	"fmt"
	"math/rand"
)

// serviceEndpoints 服务入口列表
var serviceEndpoints = map[string]map[string]string{
	ServiceGoods: {},
}

// GetServiceEndpoint 获取服务入口
func GetServiceEndpoint(svc string) (key, endpoint string, err error) {
	endpoints := serviceEndpoints[svc]
	if len(endpoints) == 0 {
		return "", "", errors.New(fmt.Sprintf("%s服务不可用，请稍后再试", svc))
	}
	num := len(endpoints)
	keys := make([]string, num)
	for v := range endpoints {
		keys = append(keys, v)
	}
	randomKey := keys[rand.Intn(len(keys))]
	return randomKey, endpoints[randomKey], nil
}
