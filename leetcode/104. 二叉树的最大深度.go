package leetcode

/*
* @author: hjz
* @time: 2026/6/9 20:33
 */

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
