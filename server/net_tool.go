package main

import (
	"fmt"
	"log"
	"net"
)

// endpoint 访问入口
var endpoint string

// InitEndpoint 获取当前的服务的IP地址和端口号
func InitEndpoint() {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrList {
		// 检查ip地址判断是否回环地址
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				endpoint = fmt.Sprintf("%s:%d", ipNet.IP.String(), servicePort)
				log.Printf("endpoint=%s\n", endpoint)
				break
			}
		}
	}
}
