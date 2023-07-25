package roundrobinbalancing

import "fmt"

type roundRobinLoadBalancer struct {
	server             []string
	currentServerIndex int
}

func NewRoundRobinLoadBalancer() *roundRobinLoadBalancer {
	return &roundRobinLoadBalancer{
		server:             make([]string, 0),
		currentServerIndex: 0,
	}
}
func (r *roundRobinLoadBalancer) addServer(server string) {
	r.server = append(r.server, server)
}

func (r *roundRobinLoadBalancer) nextServer() string {
	if len(r.server) == 0 {
		return "" // if server is empty
	}

	nextServer := r.server[r.currentServerIndex]
	r.currentServerIndex = (r.currentServerIndex + 1) % len(r.server)
	return nextServer
}

func RoundRobinBalancing() {
	loadBalancer := NewRoundRobinLoadBalancer()
	loadBalancer.addServer("Server 1")
	loadBalancer.addServer("Server 2")
	loadBalancer.addServer("Server 3")
	for i := 0; i < 10; i++ {
		server := loadBalancer.nextServer()
		fmt.Printf("Request %d directed to: %s\n", i+1, server)
	}
}
