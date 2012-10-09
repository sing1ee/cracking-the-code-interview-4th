package main

import (
	"fmt"
	"container/list"
	"math/rand"
)

/*
*Given a directed graph, design an algorithm to find out whether there is a route between tow nodes
*/

const (
	used = iota
	unused
)

type Graph struct {
	nodes *list.List
}

func (g *Graph) GetAllNodes() *list.List {
	return g.nodes
}

type Node struct {
	neighbours *list.List
	value interface{}
	state int
}

func (n *Node) GetAllNeighbours() *list.List {
	return n.neighbours
}

func Route(g *Graph, start, end *Node) bool {
	for e := g.GetAllNodes().Front(); e != nil; e = e.Next() {
		e.Value.(*Node).state = unused
	}
	stack := make([]*Node, g.GetAllNodes().Len())
	top := 0
	stack[top] = start
	start.state = used

	for ; top >= 0; {
		el := stack[top]
		if el == end {
			return true
		}
		for e := el.GetAllNeighbours().Front(); e != nil; e = e.Next() {
			node := e.Value.(*Node)
			if node.state == unused {
				node.state = used
				top = top + 1
				stack[top] = node
			}
		}
	}
}

func main() {
	
}

