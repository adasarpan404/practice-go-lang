package weightedroundrobinbalancing

import "fmt"

type Server struct {
	name   string
	weight int
}

type weightedLoadBalancer struct {
	servers            []Server
	currentServerIndex int
	gcd                int
	currentWeight      int
	maxWeight          int
}

func NewWeightedRoundRobinLoadBalancer() *weightedLoadBalancer {
	return &weightedLoadBalancer{
		currentServerIndex: 0,
		currentWeight:      0,
		gcd:                0,
		maxWeight:          0,
	}
}

func (l *weightedLoadBalancer) getGCD(a, b int) int {
	for b != 0 {
		temp := b
		b = a % b
		a = temp
	}
	return a
}
func (l *weightedLoadBalancer) getMaxWeight() int {
	max := 0
	for _, server := range l.servers {
		if server.weight > max {
			max = server.weight
		}
	}
	return max
}
func (l *weightedLoadBalancer) addServer(stringName string, serverWeight int) {
	server := Server{
		name:   stringName,
		weight: serverWeight,
	}
	l.servers = append(l.servers, server)
	l.gcd = l.getGCD(l.gcd, serverWeight)
	l.maxWeight = l.getMaxWeight()
}

func (l *weightedLoadBalancer) nextServer() string {
	if len(l.servers) == 0 {
		return "" // No servers available
	}

	for {
		l.currentServerIndex = (l.currentServerIndex + 1) % len(l.servers)
		if l.currentServerIndex == 0 {
			l.currentWeight -= l.gcd
			if l.currentWeight <= 0 {
				l.currentWeight = l.maxWeight
				if l.currentWeight == 0 {
					return "" // No servers available
				}
			}
		}

		if l.servers[l.currentServerIndex].weight >= l.currentWeight {
			return l.servers[l.currentServerIndex].name
		}
	}
}

func WeightedRoundRobinLoadBalancer() {
	loadBalancer := NewWeightedRoundRobinLoadBalancer()

	// Add servers with weights
	loadBalancer.addServer("Server 1", 3)
	loadBalancer.addServer("Server 2", 1)
	loadBalancer.addServer("Server 3", 2)

	// Retrieve servers in weighted round-robin fashion
	for i := 0; i < 10; i++ {
		server := loadBalancer.nextServer()
		fmt.Printf("Request %d directed to: %s\n", i+1, server)
	}
}
