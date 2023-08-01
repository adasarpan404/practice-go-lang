package dynamicloadbalancing

import (
	"fmt"
	"math/rand"
	"time"
)

var seed int64 = 12345

func lcgRand() float64 {
	const a, c, m = 1664525, 10139004223, 100
	currentSeed := seed
	currentSeed = (a*currentSeed + c) % m
	normalizedValue := float64(currentSeed) / float64(m)
	return normalizedValue*(2.0-0.1) + 0.1
}

type Server struct {
	name string
	load float64
}

type DynamicLoadBalancer struct {
	servers []Server
}

func (dlb *DynamicLoadBalancer) addServer(serverName string) {
	server := Server{name: serverName, load: 0.0}
	dlb.servers = append(dlb.servers, server)
}

func (dlb *DynamicLoadBalancer) getNextServer(load float64) string {
	if len(dlb.servers) == 0 {
		return ""
	}

	minLoad := dlb.servers[0].load
	minIndex := 0

	for i := 1; i < len(dlb.servers); i++ {
		if dlb.servers[i].load < minLoad {
			minLoad = dlb.servers[i].load
			minIndex = i
		}
	}

	dlb.servers[minIndex].load += load

	return dlb.servers[minIndex].name
}

func DynamicLoad() {
	var loadBalancer DynamicLoadBalancer

	loadBalancer.addServer("Server 1")
	loadBalancer.addServer("Server 2")
	loadBalancer.addServer("Server 3")

	for i := 0; i < 10; i++ {
		server := loadBalancer.getNextServer(lcgRand())
		fmt.Printf("Request %d directed to: %s\n", i+1, server)
	}
}
