package main

import (
	"container/list"
	"fmt"
	"strconv"
)

var visited = make(map[string]*Node)

// Blocker ...
type Blocker interface {
	IsBlocked(n *Node) bool
}

type blocker struct {
	nodes []*Node
}

func (b *blocker) IsBlocked(node *Node) bool {
	for idx := range b.nodes {
		n := b.nodes[idx]

		if compare(node, n) {
			return true
		}
	}
	return false
}

// NewBlocker ...
func NewBlocker(nodes []*Node) Blocker {
	return &blocker{nodes}
}

// Node ...
type Node struct {
	pos1 int
	pos2 int
	pos3 int
	pos4 int

	level int
}

func (n *Node) String() string {
	return strconv.Itoa(n.pos1) + "," + strconv.Itoa(n.pos2) + "," +
		strconv.Itoa(n.pos3) + "," + strconv.Itoa(n.pos4)
}

func (n *Node) increment(x int) int {
	x = x + 1

	if x/9 > 0 {
		return 0
	}

	return x
}

func (n *Node) decrement(x int) int {
	x = x - 1

	if x < 0 {
		return 9
	}

	return x
}

// Children ...
func (n *Node) Children() []*Node {
	return []*Node{
		&Node{n.increment(n.pos1), n.pos2, n.pos3, n.pos4, n.level + 1},
		&Node{n.decrement(n.pos1), n.pos2, n.pos3, n.pos4, n.level + 1},
		&Node{n.pos1, n.increment(n.pos2), n.pos3, n.pos4, n.level + 1},
		&Node{n.pos1, n.decrement(n.pos2), n.pos3, n.pos4, n.level + 1},
		&Node{n.pos1, n.pos2, n.increment(n.pos3), n.pos4, n.level + 1},
		&Node{n.pos1, n.pos2, n.decrement(n.pos3), n.pos4, n.level + 1},
		&Node{n.pos1, n.pos2, n.pos3, n.increment(n.pos4), n.level + 1},
		&Node{n.pos1, n.pos2, n.pos3, n.decrement(n.pos4), n.level + 1},
	}
}

// Queue ...
type Queue struct {
	*list.List
}

// Push ...
func (q *Queue) Push(node *Node) {
	q.PushFront(node)
}

// Pop ...
func (q *Queue) Pop() *Node {
	if q.Len() == 0 {
		return nil
	}

	return q.Remove(q.Back()).(*Node)
}

// NewQueue ...
func NewQueue() *Queue { return &Queue{list.New()} }

func compare(node *Node, n *Node) bool {
	if node.pos1 == n.pos1 &&
		node.pos2 == n.pos2 &&
		node.pos3 == n.pos3 &&
		node.pos4 == n.pos4 {
		return true
	}
	return false
}

func isVisited(n *Node) bool {
	_, ok := visited[n.String()]
	return ok
}

func markVisited(n *Node) {
	visited[n.String()] = n
}

func main() {
	start := &Node{1, 2, 3, 4, 1}
	end := &Node{7, 1, 3, 9, -1}

	blocked := []*Node{
		&Node{1, 2, 3, 5, -1},
		&Node{1, 2, 3, 3, -1},
	}

	blocker := NewBlocker(blocked)

	queue := NewQueue()
	queue.Push(start)

	for queue.Len() != 0 {
		n := queue.Pop()

		if isVisited(n) {
			continue
		}

		fmt.Println(n)

		if compare(n, end) {
			fmt.Println(n.level)
			return
		}

		markVisited(n)

		for _, c := range n.Children() {
			if !blocker.IsBlocked(c) {
				queue.Push(c)
			}
		}
	}

}
