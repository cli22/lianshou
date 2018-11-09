package main

import (
	"lianshou/cmd/offer/common"
)

func count(root *common.BinaryTreeNode) int {
	if root == nil{
		return 0
	}

	return count(root.Left)
}
