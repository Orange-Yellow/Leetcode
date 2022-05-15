package main

/*
	name: 449. 序列化和反序列化二叉搜索树
	link: https://leetcode.cn/problems/serialize-and-deserialize-bst/
*/

import (
	"math"
	"strconv"
	"strings"
)

func main() {

}

// 以下为官方题解的后序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 定义二叉树
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (this *Codec) Serialize(root *TreeNode) string {
	ans := []string{}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 后续遍历
		traversal(node.Left)
		traversal(node.Right)
		ans = append(ans, strconv.Itoa(node.Val))
	}
	traversal(root)
	return strings.Join(ans, " ")
}

// Deserializes your encoded data to tree.
func (this *Codec) Deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	var contruct func(left, right int) *TreeNode
	contruct = func(left, right int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}

		val, _ := strconv.Atoi(arr[len(arr)-1])
		if val < left || val > right {
			return nil
		}

		// 移除最后一个元素
		arr = arr[:len(arr)-1]

		return &TreeNode{Val: val, Right: contruct(val, right), Left: contruct(left, val)}
	}
	return contruct(math.MinInt32, math.MaxInt32)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
