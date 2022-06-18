package main

/*
	name: 剑指 Offer II 029. 排序的循环链表
	link: https://leetcode.cn/problems/4ueAj6/
*/

func main() {

}

type Node struct {
	Val  int
	Next *Node
}

func insert(head *Node, insertVal int) *Node {
	node := &Node{Val: insertVal}
	if head == nil {
		node.Next = node
		return node
	}
	if head.Next == head {
		head.Next = node
		node.Next = head
		return head
	}
	curr, next := head, head.Next
	for next != head {
		if insertVal >= curr.Val && insertVal <= next.Val {
			break
		}
		if curr.Val > next.Val {
			if insertVal > curr.Val || insertVal < next.Val {
				break
			}
		}
		curr = curr.Next
		next = next.Next
	}
	curr.Next = node
	node.Next = next
	return head
}
