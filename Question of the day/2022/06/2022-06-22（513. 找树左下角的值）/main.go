package main

/*
	name: 513. 找树左下角的值
	link: https://leetcode.cn/problems/find-bottom-left-tree-value/
*/

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// DFS 深度优先搜索
func findBottomLeftValue(root *TreeNode) int {
	curHeight := 0
	curVal := 0
	var dfs func(node *TreeNode, height int)
	dfs = func(node *TreeNode, height int) {
		if node == nil {
			return
		}
		height++
		dfs(node.Left, height)
		dfs(node.Right, height)
		// 因为先遍历的是左子树，所以同一高度，最左结点是最先遍历到的
		if height > curHeight {
			curHeight = height
			curVal = node.Val
		}
	}
	dfs(root, 0)
	return curVal
}

// BFS 广度优先搜索
func FindBottomLeftValue(root *TreeNode) int {
	q := []*TreeNode{root}
	ans := 0
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		//在遍历一个节点时，需要先把它的非空右子节点放入队列，然后再把它的非空左子节点放入队列，这样才能保证从右到左遍历每一层的节点。
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		// 遍历的最后一个节点就是最底层的最后左节点
		ans = node.Val
	}
	return ans
}
