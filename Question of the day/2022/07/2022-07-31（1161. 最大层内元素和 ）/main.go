package main

/*
	name: 1161. 最大层内元素和
	link: https://leetcode.cn/problems/maximum-level-sum-of-a-binary-tree/
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) (ans int) {
	sum := []int{}
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, level int) {
		if level == len(sum) {
			sum = append(sum, node.Val)
		} else {
			sum[level] += node.Val
		}
		if node.Left != nil {
			dfs(node.Left, level+1)
		}
		if node.Right != nil {
			dfs(node.Right, level+1)
		}
	}
	dfs(root, 0)
	for i, s := range sum {
		if s > sum[ans] {
			ans = i
		}
	}
	return ans + 1 // 层号从 1 开始
}
