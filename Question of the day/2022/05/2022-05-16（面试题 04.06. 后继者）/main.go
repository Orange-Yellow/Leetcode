package main

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

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
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
