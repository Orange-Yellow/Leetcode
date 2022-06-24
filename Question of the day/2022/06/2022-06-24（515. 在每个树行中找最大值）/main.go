package main

/*
	name: 515. 在每个树行中找最大值
	link: https://leetcode.cn/problems/find-largest-value-in-each-tree-row/
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	ans := []int{}
	var dfs func(node *TreeNode, height int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		if len(ans) == height {
			ans = append(ans, node.Val)
		} else {
			if ans[height] < node.Val {
				ans[height] = node.Val
			}

		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
	}
	dfs(root, 0)
	return ans
}
