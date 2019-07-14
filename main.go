package main

import (
	"fmt"
	"math/rand"
	"time"
)

var arrs = []int{0, 10, 4, 5, 2, 9, 8, 6, 7, 1, 3}

func main() {
	rand.Seed(time.Now().UnixNano())

	skipLists := make([]*node, 0)

	for _, value := range arrs {
		skipLists = insert(skipLists, value)
	}

	scan(skipLists)

	fmt.Print("")
}

func scan(skipLists []*node) {
	if len(skipLists) == 0 {
		return
	}

	for i := len(skipLists) - 1; i > -1; i-- {
		start := skipLists[i]
		for {
			if start == nil {
				break
			}
			fmt.Printf("%v -> ", start.value)
			start = start.right
		}
		fmt.Println()
	}
}

type node struct {
	value int
	right *node
	down  *node
}

func insert(skipLists []*node, value int) []*node {
	if len(skipLists) == 0 {
		skipLists = append(skipLists, newNode(value))
		return skipLists
	}

	level := level()
	if level > len(skipLists)-1 {
		for len(skipLists) < level+1 {
			s := skipLists[len(skipLists)-1]
			skipLists = append(skipLists, &node{
				value: s.value,
				down:  s,
				right: nil,
			})
		}
	}

	var lastnodesaved *node
	start := skipLists[level]
	for {
		for start.right != nil && start.value < value && start.right.value < value {
			start = start.right
		}

		nNode := newNode(value)

		right := start.right
		start.right = nNode
		nNode.right = right

		if lastnodesaved != nil {
			lastnodesaved.down = nNode
		}
		lastnodesaved = nNode

		if start.down == nil {
			break
		}
		start = start.down
	}
	return skipLists
}

func newNode(value int) *node {
	return &node{value: value}
}

func search(root *node, new *node) {
	start := root

	for {
		for start.right != nil && start.value < new.value {
			start = start.right
		}

		if start.down == nil {

		}

		start = start.down
	}
}

func level() int {
	k := 0
	for {
		i := rand.Intn(2)
		if i == 1 {
			k++
		} else {
			return k
		}
	}
}
