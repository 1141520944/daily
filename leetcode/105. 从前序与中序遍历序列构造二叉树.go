package leetcode

/*
* @author: hjz
* @time: 2026/6/9 20:40
 */
//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0 {
//		return nil
//	}
//	root := &TreeNode{Val: preorder[0]}
//	var index int
//	for i := range inorder {
//		if inorder[i] == preorder[0] {
//			index = i
//			break
//		}
//	}
//	root.Left = buildTree(preorder[1:index+1], inorder[:index])
//	root.Right = buildTree(preorder[index+1:], inorder[index+1:])
//	return root
//}
