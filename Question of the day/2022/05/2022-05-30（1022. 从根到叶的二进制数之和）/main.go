package main

import "fmt"

/*
	name: 1022. 从根到叶的二进制数之和
	link: https://leetcode.cn/problems/sum-of-root-to-leaf-binary-numbers/
*/

func main() {
	x := 1
	y := 0
	fmt.Println(x<<1 | y)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 迭代
func SumRootToLeaf(root *TreeNode) int {
	var que []map[string]interface{}
	que = append(que, map[string]interface{}{"n": root, "v": 0})
	res := 0
	for len(que) > 0 {
		node, total := que[0]["n"], que[0]["v"]
		que = que[1:]
		n, _ := node.(*TreeNode)
		t, _ := total.(int)

		t = t*2 + n.Val

		if n.Left == nil && n.Right == nil {
			res += t
		}

		if n.Left != nil {
			que = append(que, map[string]interface{}{"n": n.Left, "v": t})
		}

		if n.Right != nil {
			que = append(que, map[string]interface{}{"n": n.Right, "v": t})
		}
	}
	return res
}

// 递归
func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	val = val<<1 | node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}
	return dfs(node.Left, val) + dfs(node.Right, val)
}
