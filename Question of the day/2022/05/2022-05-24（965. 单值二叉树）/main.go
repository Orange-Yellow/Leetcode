package main

/*
	name: 965. 单值二叉树
	link: https://leetcode.cn/problems/univalued-binary-tree/
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	val := root.Val
	return isSame(root.Right, val) && isSame(root.Left, val)
}

func isSame(root *TreeNode, val int) bool {
	if root == nil {
		return true
	}
	if root.Val != val {
		return false
	}
	return isSame(root.Right, val) && isSame(root.Left, val)
}
