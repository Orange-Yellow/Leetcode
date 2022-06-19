package main

/*
	name: 508. 出现次数最多的子树元素和
	link: https://leetcode.cn/problems/most-frequent-subtree-sum/
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findFrequentTreeSum(root *TreeNode) []int {
	ans := []int{}
	hash := map[int]int{}
	maxCnt := 0
	// 深度优先搜索，查找所有存在的子树和及其出现的次数
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := node.Val + dfs(node.Left) + dfs(node.Right)
		hash[sum]++
		if hash[sum] > maxCnt {
			maxCnt = hash[sum]
		}
		return sum
	}
	dfs(root)
	// 将出现次数最大的子树和加到结果集中
	for s, c := range hash {
		if c == maxCnt {
			ans = append(ans, s)
		}
	}
	return ans
}
