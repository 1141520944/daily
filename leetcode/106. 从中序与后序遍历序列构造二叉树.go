package leetcode

/*
* @author: hjz
* @time: 2026/6/9 20:50
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[len(postorder)-1]}
	var index int
	for i := range inorder {
		if inorder[i] == root.Val {
			index = i
			break
		}
	}
	root.Left = buildTree(inorder[:index], postorder[:index])
	root.Right = buildTree(inorder[index+1:], postorder[index:len(postorder)-1])
	return root
}
