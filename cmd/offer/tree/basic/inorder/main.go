package main

import (
	"fmt"

	"lianshou/cmd/offer/common"
)

func inorder(root *common.BinaryTreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Println(root.Value)
	inorder(root.Right)
}

func main() {
	root := &common.BinaryTreeNode{1, &common.BinaryTreeNode{2, &common.BinaryTreeNode{3, nil, nil}, nil}, nil}
	inorder(root)
}
