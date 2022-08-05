package main

/*
	name: 623. 在二叉树中增加一行
	link: https://leetcode.cn/problems/add-one-row-to-tree/
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}
	nodes := []*TreeNode{root}
	for i := 1; i < depth-1; i++ {
		tmp := nodes
		nodes = nil
		for _, node := range tmp {
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
	}
	for _, node := range nodes {
		node.Left = &TreeNode{val, node.Left, nil}
		node.Right = &TreeNode{val, nil, node.Right}
	}
	return root
}
