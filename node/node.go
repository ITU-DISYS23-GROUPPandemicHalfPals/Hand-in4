package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"me/me"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = flag.Int("port", 5000, "The port of the node")

type node struct {
	Port            int
	CoordinatorPort int
	Ports           []int

	Token            bool
	CoordinatorToken bool

	TokenRequestChannel chan int
	TokenChannel        chan bool
	ElectionChannel     chan bool

	Clients       map[int]me.MutualExclusionClient
	BiggerClients map[int]me.MutualExclusionClient

	me.UnimplementedMutualExclusionServer
}

func Node(port int) *node {
	var ports []int

	for i := 5000; i <= 5002; i++ {
		ports = append(ports, i)
	}

	return &node{
		Port:            port,
		CoordinatorPort: 0,
		Ports:           ports,

		Token:            false,
		CoordinatorToken: false,

		TokenRequestChannel: make(chan int, 10),
		TokenChannel:        make(chan bool, 1),
		ElectionChannel:     make(chan bool, 10),

		Clients:       make(map[int]me.MutualExclusionClient),
		BiggerClients: make(map[int]me.MutualExclusionClient),
	}
}

func main() {
	flag.Parse()

	n := Node(*port)

	ctx := context.Background()

	go n.server(ctx)
	n.client(ctx)
}

func (n *node) server(ctx context.Context) {
	server := grpc.NewServer()
	me.RegisterMutualExclusionServer(server, n)

	listener, error := net.Listen("tcp", ":"+strconv.Itoa(n.Port))
	if error != nil {
		log.Fatalf("Failed to listen: %s", error)
	}

	error = server.Serve(listener)
	if error != nil {
		log.Fatalf("Failed to serve: %s", error)
	}
}

func (n *node) Election(_ context.Context, request *me.ElectionMessage) (*me.Response, error) {
	n.ElectionChannel <- true

	return &me.Response{}, nil
}

func (n *node) Coordinator(_ context.Context, request *me.CoordinatorMessage) (*me.Response, error) {
	n.CoordinatorPort = int(request.Port)

	if n.CoordinatorPort == n.Port {
		n.CoordinatorToken = true
	} else {
		n.CoordinatorToken = false
	}

	return &me.Response{}, nil
}

func (n *node) RequestToken(_ context.Context, request *me.TokenRequest) (*me.Response, error) {
	n.TokenRequestChannel <- int(request.Port)

	return &me.Response{}, nil
}

func (n *node) GrantToken(_ context.Context, request *me.TokenMessage) (*me.Response, error) {
	n.TokenChannel <- true

	return &me.Response{}, nil
}

func (n *node) ReleaseToken(_ context.Context, request *me.TokenMessage) (*me.Response, error) {
	if n.CoordinatorPort == n.Port {
		n.CoordinatorToken = true
	}

	return &me.Response{}, nil
}

func (n *node) client(ctx context.Context) {
	go n.dialServers()
	go n.broadcastElection(ctx)
	go n.tokenDistributer(ctx)
	n.run(ctx)
}

func (n *node) dialServers() {
	for {
		time.Sleep(time.Second)

		for _, port := range n.Ports {
			_, ok := n.Clients[port]

			if ok {
				continue
			}

			connection, error := grpc.Dial(":"+strconv.Itoa(port), grpc.WithTransportCredentials(insecure.NewCredentials()))
			if error != nil {
				continue
			}

			client := me.NewMutualExclusionClient(connection)

			n.Clients[port] = client

			if port > n.Port {
				n.BiggerClients[port] = client
			}
		}
	}
}

func (n *node) tokenDistributer(ctx context.Context) {
	for {
		port := <-n.TokenRequestChannel

		for !n.CoordinatorToken {

		}

		n.CoordinatorToken = false
		client := n.Clients[port]
		client.GrantToken(ctx, &me.TokenMessage{})
	}
}

func (n *node) broadcastElection(ctx context.Context) {
	for {
		<-n.ElectionChannel
		n.startElection(ctx)
	}
}

func (n *node) startElection(ctx context.Context) {
	response := false
	for _, client := range n.BiggerClients {
		ctx, cancel := context.WithTimeout(ctx, time.Second)
		_, error := client.Election(ctx, &me.ElectionMessage{})

		if error == nil {
			response = true
		}

		cancel()
	}

	if !response {
		coordinatorMessage := &me.CoordinatorMessage{
			Port: int32(n.Port),
		}

		for _, client := range n.Clients {
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			_, _ = client.Coordinator(ctx, coordinatorMessage)
			cancel()
		}
	}
}

func (n *node) run(ctx context.Context) {
	for {
		client, ok := n.Clients[n.CoordinatorPort]
		if !ok {
			log.Print("No coordinator found: Starting new election")
			n.startElection(ctx)
			n.sleep()
			continue
		}

		println()
		log.Printf("Requesting the token - Using Coordinator: %d", n.CoordinatorPort)
		tokenRequest := &me.TokenRequest{
			Port: int32(n.Port),
		}

		_, error := client.RequestToken(ctx, tokenRequest)
		if error != nil {
			log.Print("Coordinator crashed: Starting new election")
			n.startElection(ctx)
			n.sleep()
			continue
		}

		n.Token = <-n.TokenChannel
		log.Print("Granted token")

		log.Print("Entering critical section")

		n.criticalSection()

		log.Printf("Left critical section")

		n.Token = false
		client.ReleaseToken(ctx, &me.TokenMessage{})
		log.Print("Released token")

		n.sleep()
	}
}

func (n *node) sleep() {
	wait := rand.Intn(3)
	time.Sleep(time.Second*time.Duration(wait) + time.Second*3)
}

func (n *node) criticalSection() {
	n.sleep()
	log.Printf("Using the critical section")
	n.sleep()
}
