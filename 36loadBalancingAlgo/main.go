package main

import (
	"fmt"

	"github.com/adasarpan404/loadbalancing/roundrobinbalancing"
	"github.com/adasarpan404/loadbalancing/weightedroundrobinbalancing"
)

func main() {
	fmt.Println("Round Robin loadbalancing")
	roundrobinbalancing.RoundRobinBalancing()
	fmt.Println("Weighted Round Robin Load Balancing")
	weightedroundrobinbalancing.WeightedRoundRobinLoadBalancer()
}
