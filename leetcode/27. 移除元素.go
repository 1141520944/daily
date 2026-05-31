package leetcode

/*
* @author: hjz
* @time: 2026/5/31 15:06
* @function:
* 类似移动0 当前数字不是0和慢指针交换
 */
func removeElement(nums []int, val int) int {
	s, f := 0, 0
	for f < len(nums) {
		if nums[f] != val {
			nums[s] = nums[f]
			s++
		}
		f++
	}
	return s
}
