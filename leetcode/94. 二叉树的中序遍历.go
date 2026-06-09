package leetcode

/*
* @author: hjz
* @time: 2026/6/9 22:15
 */
func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorderTraversalHelp(root, &res)
	return res
}
func inorderTraversalHelp(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorderTraversalHelp(root.Left, res)
	*res = append(*res, root.Val)
	inorderTraversalHelp(root.Right, res)
}
