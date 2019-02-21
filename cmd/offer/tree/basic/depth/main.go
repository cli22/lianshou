package main

import (
	"lianshou/cmd/offer/common"
	"fmt"
)

func getDepth(root *common.BinaryTreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := getDepth(root.Left)
	rightDepth := getDepth(root.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func main() {
	root := &common.BinaryTreeNode{1, &common.BinaryTreeNode{2, &common.BinaryTreeNode{1, nil, nil}, nil}, nil}
	depth := getDepth(root)
	fmt.Println(depth)
}
