package leetcode

/*
* @author: hjz
* @time: 2026/5/21 22:39
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
