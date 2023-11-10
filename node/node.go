package main

import (
	"context"
	"flag"
	"log"
	"me/me"
	"net"
	"strconv"

	"google.golang.org/grpc"
)

var name = flag.String("name", "John Doe", "The name of the node")
var id = flag.Int("id", 1, "The id of the node")
var port = flag.Int("port", 5000, "The port of the node")

type node struct {
	Name  string
	Id    int
	Port  int
	Ports []int

	Elections    chan *me.ElectionMessage
	Coordinators chan *me.CoordinatorMessage
	me.UnimplementedMutualExclusionServer
}

func Node(name string, id int, port int) *node {
	var ports []int

	for i := 5000; i <= 5002; i++ {
		if i == port {
			continue
		}

		ports = append(ports, i)
	}

	return &node{
		Name:  name,
		Id:    id,
		Port:  port,
		Ports: ports,

		Elections:    make(chan *me.ElectionMessage, 100),
		Coordinators: make(chan *me.CoordinatorMessage, 100),
	}
}

func main() {
	flag.Parse()

	n := Node(*name, *id, *port)

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

}

func (n *node) Coordinator(_ context.Context, request *me.CoordinatorMessage) (*me.Response, error) {

}

func (n *node) client() {

}
