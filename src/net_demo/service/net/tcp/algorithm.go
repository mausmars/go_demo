package tcp

import (
	"fmt"
	"strings"
)

type LinkedNode struct {
	pre  *LinkedNode
	next *LinkedNode
	data interface{}
}

type DoubleLinkedQueue struct {
	head *LinkedNode
	last *LinkedNode
}

func (q *DoubleLinkedQueue) Push(v interface{}) {
	var nNode = new(LinkedNode)
	nNode.data = v
	if (q.last == nil) {
		nNode.next = nNode
		nNode.pre = nNode
		q.last = nNode
		q.head = nNode
	} else {
		nNode.pre = q.last
		nNode.next = q.last.next
		q.last.next.pre = nNode
		q.last.next = nNode
		q.last = nNode
	}
	//fmt.Println("添加lst节点 ", nNode.data)
}

func (q *DoubleLinkedQueue) Pop() *LinkedNode {
	if (q.head == nil) {
		return nil
	} else if (q.head == q.last) {
		nNode := q.head
		q.last = nil
		q.head = nil
		nNode.next = nil
		nNode.pre = nil
		//fmt.Println("移除fist节点 ", nNode.data)
		return nNode
	} else {
		q.head.next.pre = q.head.pre
		q.head.pre.next = q.head.next
		nNode := q.head
		q.head = q.head.next
		nNode.next = nil
		nNode.pre = nil
		//fmt.Println("移除fist节点 ", nNode.data)
		return nNode
	}
}

func (q *DoubleLinkedQueue) Peek() *LinkedNode {
	if (q.head == nil) {
		return nil
	} else {
		return q.head
	}
}

func (q *DoubleLinkedQueue) Print() {
	var sb strings.Builder
	fmt.Fprint(&sb, "顺序打印")
	fmt.Fprint(&sb, " [")

	c := q.head
	for ; ; {
		if (c == nil) {
			break
		}
		fmt.Fprint(&sb, c.data)
		if (c.next == q.head) {
			break
		}
		fmt.Fprint(&sb, ",")
		c = c.next
	}
	fmt.Fprint(&sb, "]")
	fmt.Println(sb.String())
}

func TestDoubleLinkedQueue() {
	queue := new(DoubleLinkedQueue)

	for i := 1; i <= 3; i++ {
		queue.Push(i)
	}
	queue.Print()

	for i := 1; i <= 4; i++ {
		node := queue.Pop()
		if (node == nil) {
			fmt.Println("Pop ", node)
		} else {
			fmt.Println("Pop ", node.data)
		}
		queue.Print()
	}
}
