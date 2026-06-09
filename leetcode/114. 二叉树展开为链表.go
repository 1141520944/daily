package leetcode

/*
* @author: hjz
* @time: 2026/6/9 21:12
 */

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	r := root.Right
	if root.Left != nil {
		root.Right = root.Left
		root.Left = nil
		l := root.Right
		for l.Right != nil {
			l = l.Right
		}
		l.Right = r
	}
}
