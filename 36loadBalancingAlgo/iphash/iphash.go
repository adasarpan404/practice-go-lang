package iphash

import "fmt"

type IpHashLoadBalancer struct {
	serverMap map[string]string
}

func NewIpHashLoadBalancer() *IpHashLoadBalancer {
	return &IpHashLoadBalancer{
		serverMap: make(map[string]string),
	}
}

func (l *IpHashLoadBalancer) addServer(ipAddress, serverName string) {
	l.serverMap[ipAddress] = serverName
}

func (l *IpHashLoadBalancer) getServer(ipAddress string) string {
	if serverName, found := l.serverMap[ipAddress]; found {
		return serverName
	}
	return ""
}

func IpHash() {
	loadBalancer := NewIpHashLoadBalancer()

	loadBalancer.addServer("192.168.0.1", "Server 1")
	loadBalancer.addServer("192.168.0.2", "Server 2")
	loadBalancer.addServer("192.168.0.3", "Server 3")
	fmt.Println("Request from 192.168.0.1 directed to:", loadBalancer.getServer("192.168.0.1"))
	fmt.Println("Request from 192.168.0.2 directed to:", loadBalancer.getServer("192.168.0.2"))
	fmt.Println("Request from 192.168.0.3 directed to:", loadBalancer.getServer("192.168.0.3"))
	fmt.Println("Request from 192.168.0.4 directed to:", loadBalancer.getServer("192.168.0.4"))
}
