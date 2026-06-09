package leetcode

/*
* @author: hjz
* @time: 2026/6/9 22:05
 */
func sumNumbers(root *TreeNode) int {
	var num []int
	sumNumbersHelp(root, 0, &num)
	var sum int
	for _, v := range num {
		sum += v
	}
	return sum
}
func sumNumbersHelp(root *TreeNode, i int, num *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*num = append(*num, i*10+root.Val)
	}
	sumNumbersHelp(root.Left, i*10+root.Val, num)
	sumNumbersHelp(root.Right, i*10+root.Val, num)
}
