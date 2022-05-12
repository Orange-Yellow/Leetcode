package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	name：297. 二叉树的序列化与反序列化
	link: https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
*/

func main() {
	val := strconv.Itoa(-7)
	fmt.Println(val, len(val))
}

// 以下为官方题解
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() (_ Codec) {
	return
}

// 序列化：前序或者后序遍历生成字符串
func (this *Codec) Serialize(root *TreeNode) string {
	ans := &strings.Builder{}
	var serialization func(node *TreeNode)
	serialization = func(node *TreeNode) {
		// 以，分割字符 因为数字可能占用多为字符
		if node == nil {
			ans.WriteString("-,")
			return
		}
		ans.WriteString(strconv.Itoa(node.Val))
		ans.WriteByte(',')
		serialization(node.Left)
		serialization(node.Right)
	}
	serialization(root)
	return ans.String()
}

// 反序列化
func (this *Codec) Deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	index := 0
	var contract func() *TreeNode
	contract = func() *TreeNode {
		if arr[index] == "-" {
			index++
			return nil
		}
		val, _ := strconv.Atoi(arr[index])
		index++
		return &TreeNode{Val: val, Left: contract(), Right: contract()}
	}
	return contract()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
