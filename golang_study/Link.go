package main

import (
	"fmt"
)

type node struct {
	data int
	next *node
}

//链表相关
func main() {
	length := 5
	head := &node{0, nil}
	head = createLink(head, length)
	printLink(head)
	reversedHead := reverseLink(head)
	printLink(reversedHead)
	merge := mergeTwoLists(reversedHead, reversedHead)
	printLink(merge)
}

func createLink(head *node, length int) *node {
	if length <= 0 {
		return head
	}
	for i := length - 1; i > 0; i-- {
		p := &node{i, nil}
		p.next = head.next
		head.next = p
	}
	return head
}

func printLink(head *node) {
	for p := head; p != nil; p = p.next {
		fmt.Print(p.data)
		fmt.Print(" ")
	}
	fmt.Println()
}

func reverseLink(head *node) *node {
	if head == nil || head.next == nil {
		return head
	}
	var reversedHead *node
	var p *node

	reversedHead = head
	head = head.next
	reversedHead.next = nil
	p = head.next
	for head != nil {
		head.next = reversedHead
		reversedHead = head
		head = p
		if p != nil {
			p = p.next
		}
	}
	return reversedHead
}

func mergeTwoLists(node1 *node, node2 *node) *node {
	if node1 == nil && node2 == nil {
		return nil
	}
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	var head *node
	var current *node
	if node1.data <= node2.data {
		head = node1
		current = node1
		node1 = node1.next
	} else {
		head = node2
		current = node2
		node2 = node2.next
	}

	for node1 != nil && node2 != nil {
		if node1.data <= node2.data {
			current.next = node1
			current = current.next
			node1 = node1.next
		} else {
			current.next = node2
			current = current.next
			node2 = node2.next
		}
	}
	//if node1 != nil {
	//	current.next = node1
	//}
	//if node2 != nil {
	//	current.next = node2
	//}
	return head

}
