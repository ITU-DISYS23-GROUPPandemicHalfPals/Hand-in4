package main

import "mutualExclusion/mutualExclusion"

type node struct {
	Id    int
	Port  int
	Ports []int

	Elections    chan *mutualExclusion.ElectionMessage
	Coordinators chan *mutualExclusion.CoordinatorMessage
}

func Node(id int, port int) *node {
	return &node{
		Id:   id,
		Port: port,

		Elections:    make(chan *mutualExclusion.ElectionMessage, 100),
		Coordinators: make(chan *mutualExclusion.CoordinatorMessage, 100),
	}
}

func main() {
	n := Node(1, 5523)

	n.client()
	n.server()
}

func (n *node) client() {

}

func (n *node) server() {

}
