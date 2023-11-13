package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"me/me"
	"net"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var name = flag.String("name", "John Doe", "The name of the node")
var port = flag.Int("port", 5000, "The port of the node")

type node struct {
	Name            string
	Port            int
	CoordinatorPort int
	Ports           []int

	Elections     chan *me.ElectionMessage
	Clients       map[int]me.MutualExclusionClient
	BiggerClients map[int]me.MutualExclusionClient
	me.UnimplementedMutualExclusionServer
}

func Node(name string, port int) *node {
	var ports []int

	for i := 5000; i <= 5002; i++ {
		if i == port {
			continue
		}

		ports = append(ports, i)
	}

	return &node{
		Name:  name,
		Port:  port,
		Ports: ports,

		Elections:     make(chan *me.ElectionMessage, 10),
		Clients:       make(map[int]me.MutualExclusionClient),
		BiggerClients: make(map[int]me.MutualExclusionClient),
	}
}

func main() {
	flag.Parse()

	n := Node(*name, *port)

	go n.server()
	n.client()
}

func (n *node) server() {
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
	n.Elections <- &me.ElectionMessage{
		Port: int32(n.Port),
	}

	return &me.Response{}, nil
}

func (n *node) Coordinator(_ context.Context, request *me.CoordinatorMessage) (*me.Response, error) {
	n.CoordinatorPort = int(request.Port)

	return &me.Response{}, nil
}

func (n *node) client() {
	ctx := context.Background()

	go n.printCoordinator()
	go n.dialServers()
	go n.broadcaseElection(ctx)
	n.startElection(ctx)
}

func (n *node) printCoordinator() {
	for {
		time.Sleep(time.Second * 5)

		if n.CoordinatorPort != 0 {
			log.Printf("Coordinator is %d", n.CoordinatorPort)
		}
	}
}

func (n *node) startElection(ctx context.Context) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if scanner.Scan() {
			electionMessage := &me.ElectionMessage{
				Port: int32(n.Port),
			}

			response := false
			for _, client := range n.BiggerClients {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				_, error := client.Election(ctx, electionMessage)

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

				n.CoordinatorPort = int(n.Port)
			}
		}
	}
}

func (n *node) broadcaseElection(ctx context.Context) {
	for {
		election := <-n.Elections
		response := false
		for _, client := range n.BiggerClients {
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			_, error := client.Election(ctx, election)

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

			n.CoordinatorPort = int(n.Port)
		}
	}
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
