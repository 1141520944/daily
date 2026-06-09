package leetcode

/*
* @author: hjz
* @time: 2026/6/9 20:34
 */

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)
	return &TreeNode{
		Val:   root.Val,
		Left:  root.Right,
		Right: root.Left,
	}
}
