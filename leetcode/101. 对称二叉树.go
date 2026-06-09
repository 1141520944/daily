package leetcode

/*
* @author: hjz
* @time: 2026/6/9 20:38
 */

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelp(root.Left, root.Right)
}

func isSymmetricHelp(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSymmetricHelp(p.Left, q.Right) && isSymmetricHelp(p.Right, q.Left)
}
