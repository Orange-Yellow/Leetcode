package main

/*
	name: 面试题 04.06.后继者
	link: https://leetcode.cn/problems/successor-lcci/
*/

func main() {

}

/**
  Definition for a binary tree node.
  type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
  }
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 官方题解
func InorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var ans *TreeNode
	//《左->根->右》
	// 如果p的右节点不等于空则p的【下一个节点在p的右节点的最左节点】
	if p.Right != nil {
		ans = p.Right
		for ans.Left != nil {
			ans = ans.Left
		}
		return ans
	}

	// 如果p的右节点为空, 说明p节点的为其【下一个节点的左节点】
	node := root
	for node != nil {
		// 如果该节点的值大于p节点值就遍历下去
		if node.Val > p.Val {
			ans = node
			node = node.Left
		} else {
			// 否则就说明已经到当前节点，而且其右节点为空，结束遍历
			node = node.Right
		}
	}
	return ans
}

// 根据BFS特性 递归题解
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 如果root小于p 说明，下一节点在root的右节点
	if root.Val <= p.Val {
		return inorderSuccessor(root.Right, p)
	}
	// 如果左节点为空，说明是root本身
	ans := inorderSuccessor(root.Left, p)
	if ans == nil {
		return root
	} else {
		return ans
	}
}
