package main

import (
	"fmt"

	"lianshou/cmd/offer/common"
)

func preorder(root *common.BinaryTreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	preorder(root.Left)
	preorder(root.Right)
}

func main() {
	root := &common.BinaryTreeNode{1, &common.BinaryTreeNode{2, &common.BinaryTreeNode{3, nil, nil}, nil}, nil}
	preorder(root)
}
