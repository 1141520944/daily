package leetcode

/*
* @author: hjz
* @time: 2026/5/21 22:39
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
