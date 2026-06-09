package leetcode

/*
* @author: hjz
* @time: 2026/6/9 21:05
 */
func levelOrder(root *TreeNode) [][]int {
	queue := []*TreeNode{root}
	var level int
	res := [][]int{}
	for len(queue) > 0 {
		level = len(queue)
		item := []int{}
		for i := 0; i < level; i++ {
			node := queue[i]
			if node == nil {
				continue
			}
			item = append(item, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if len(item) > 0 {
			res = append(res, item)
		}
		queue = queue[level:]
	}
	return res
}
