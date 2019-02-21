package main

import (
	"fmt"

	"lianshou/cmd/offer/common"
)

func postorder(root *common.BinaryTreeNode) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Println(root.Value)

}

func main() {
	root := &common.BinaryTreeNode{1, &common.BinaryTreeNode{2, &common.BinaryTreeNode{3, nil, nil}, nil}, nil}
	postorder(root)
}
