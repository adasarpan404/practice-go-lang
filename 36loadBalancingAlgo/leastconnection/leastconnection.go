package leastconnection

import "fmt"

type Server struct {
	name            string
	connectionCount int
}

func NewServer(name string) *Server {
	return &Server{name: name, connectionCount: 0}
}

type LeastConnectionLoadBalancer struct {
	servers            []*Server
	currentServerIndex int
}

func NewLeastConnectionLoadBalancer() *LeastConnectionLoadBalancer {
	return &LeastConnectionLoadBalancer{
		currentServerIndex: 0,
	}
}

func (l *LeastConnectionLoadBalancer) addServer(serverName string) {
	server := NewServer(serverName)
	l.servers = append(l.servers, server)
}

func (l *LeastConnectionLoadBalancer) nextServer() string {
	if len(l.servers) == 0 {
		return ""
	}
	minIndex := 0
	minConnections := l.servers[0].connectionCount

	for i := 0; i < len(l.servers); i++ {
		if l.servers[i].connectionCount < minConnections {
			minConnections = l.servers[i].connectionCount
			minIndex = i
		}
	}
	l.servers[minIndex].connectionCount++
	return l.servers[minIndex].name
}

func LeastConnection() {
	leastConnectionLoadBalancer := NewLeastConnectionLoadBalancer()
	leastConnectionLoadBalancer.addServer("Server 1")
	leastConnectionLoadBalancer.addServer("Server 2")
	leastConnectionLoadBalancer.addServer("Server 3")
	for i := 0; i < 10; i++ {
		server := leastConnectionLoadBalancer.nextServer()
		fmt.Printf("Request %d directed to: %s\n", i+1, server)
	}
}
