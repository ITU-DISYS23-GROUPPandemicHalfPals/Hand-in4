package main

import (
	"flag"
	"mutualExclusion/mutualExclusion"
)

var name = flag.String("name", "John Doe", "The name of the node")
var id = flag.Int("id", 1, "The id of the node")
var port = flag.Int("port", 5000, "The port of the node")

type node struct {
	Name  string
	Id    int
	Port  int
	Ports []int

	Elections    chan *mutualExclusion.ElectionMessage
	Coordinators chan *mutualExclusion.CoordinatorMessage
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

		Elections:    make(chan *mutualExclusion.ElectionMessage, 100),
		Coordinators: make(chan *mutualExclusion.CoordinatorMessage, 100),
	}
}

func main() {
	flag.Parse()

	n := Node(*name, *id, *port)

	println(n.Name)
	println(n.Id)
	println(n.Port)

	go n.server()
	n.client()
}

func (n *node) server() {

}

func (n *node) client() {

}
