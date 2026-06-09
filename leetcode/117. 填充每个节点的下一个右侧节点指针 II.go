package leetcode

/*
* @author: hjz
* @time: 2026/6/9 20:52
 */

func connect(root *Node) *Node {
	queue := []*Node{root}
	var level int
	for len(queue) > 0 {
		level = len(queue)
		var prev *Node
		for i := 0; i < level; i++ {
			node := queue[i]
			if prev != nil {
				prev.Next = node
			}
			prev = node
			if node == nil {
				continue
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[level:]
	}
	return root
}
