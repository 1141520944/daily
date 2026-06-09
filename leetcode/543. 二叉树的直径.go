package leetcode

/*
* @author: hjz
* @time: 2026/6/9 22:17
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	diameterOfBinaryTreeHelp(root, &res)
	return res - 1 // 现在 root 不为空，res >= 1，返回 >=0
}

func diameterOfBinaryTreeHelp(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := diameterOfBinaryTreeHelp(root.Left, res)
	r := diameterOfBinaryTreeHelp(root.Right, res)
	*res = max(*res, l+r+1) // 更新最大节点数路径
	return max(l, r) + 1    // 返回当前子树的高度（节点数）
}
