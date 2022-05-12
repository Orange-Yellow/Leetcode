package main

import (
	"strconv"
	"strings"
)

func main() {

}

// 序列化：前序或者后序遍历生成字符串
func (this *Codec) serialize(root *TreeNode) string {
	ans := ""

	var serialization func(node *TreeNode)
	serialization = func(node *TreeNode) {
		// 以，分割字符 因为数字可能占用多为字符
		if node == nil {
			ans += "-,"
			return
		}
		ans += strconv.Itoa(node.Val) + ","
		serialization(node.Left)
		serialization(node.Right)
	}
	serialization(root)
	return ans
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	arr := strings.Split(data, ",")
	var contract func() *TreeNode
	contract = func() *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		if arr[0] == "-" {
			arr = arr[1:]
			return nil
		}
		val, _ := strconv.Atoi(arr[0])
		arr = arr[1:]
		return &TreeNode{Val: val, Left: contract(), Right: contract()}
	}
	return contract()
}
