package main

import (
	"lianshou/cmd/offer/common"
	"fmt"
)

func count(root *common.BinaryTreeNode) int {
	if root == nil {
		return 0
	}

	leftCount := count(root.Left)
	rightCount := count(root.Right)
	if leftCount > rightCount {
		return leftCount + 1
	} else {
		return rightCount + 1
	}
}

func main() {
	//root := &common.BinaryTreeNode{1, &common.BinaryTreeNode{2, nil, nil}, nil}
	//root := &common.BinaryTreeNode{1, nil, nil}
	cnt := count(nil)
	fmt.Println(cnt)
}
