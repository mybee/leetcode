package main

import "math"

//Definition for a binary tree node.
 type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }

var  last = -math.MaxFloat64

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isValidBST(root.Left) {
		if last < root.Val {
			last = root.Val
			return isValidBST(root.Right)
		}
	}
	return false
}
