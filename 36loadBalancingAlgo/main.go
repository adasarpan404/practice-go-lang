package main

import (
	"fmt"

	"github.com/adasarpan404/loadbalancing/iphash"
	"github.com/adasarpan404/loadbalancing/leastconnection"
	"github.com/adasarpan404/loadbalancing/roundrobinbalancing"
	"github.com/adasarpan404/loadbalancing/weightedroundrobinbalancing"
)

func main() {
	fmt.Println("Round Robin loadbalancing")
	roundrobinbalancing.RoundRobinBalancing()
	fmt.Println("Weighted Round Robin Load Balancing")
	weightedroundrobinbalancing.WeightedRoundRobinLoadBalancer()
	fmt.Println("Least Connection Load Balancing")
	leastconnection.LeastConnection()
	fmt.Println("Ip Hash Load Balancing")
	iphash.IpHash()
}
