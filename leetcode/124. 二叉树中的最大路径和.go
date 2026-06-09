package leetcode

/*
* @author: hjz
* @time: 2026/6/9 22:11
 */
func maxPathSum(root *TreeNode) int {

}

func maxPathSumHelp(root *TreeNode, res *int) {
	if root == nil {
		return
	}
	maxPathSumHelp(root.Left, res)
	maxPathSumHelp(root.Right, res)
}
