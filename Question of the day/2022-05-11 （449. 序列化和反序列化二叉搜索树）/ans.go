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

// 以下为自己题解的前序遍历
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 定义二叉树

// 序列化：将搜索树转化为字符串
func (this *Codec) serialize(root *TreeNode) string {
	arr := []string{}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 前序遍历
		arr = append(arr, strconv.Itoa(node.Val))
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	// 每个节点中间插入空格，方便反序列化时转化为数组
	return strings.Join(arr, " ")
}

// 反序列化： 将字符串转化为搜索树
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	var contract func(left, right int) *TreeNode
	contract = func(left, right int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		// 因为序列化时是前序遍历，所以数组的首部为根节点
		val, _ := strconv.Atoi(arr[0])
		if val < left || val > right {
			return nil
		}
		arr = arr[1:]
		return &TreeNode{Val: val, Left: contract(left, val), Right: contract(val, right)}
	}
	return contract(math.MinInt32, math.MaxInt32)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
