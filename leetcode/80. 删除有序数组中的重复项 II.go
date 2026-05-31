package leetcode

/*
* @author: hjz
* @time: 2026/5/31 15:22
 */
func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	s := 2 // 前两个一定合法
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[s-2] { // 和往前两个位置的元素比较
			nums[s] = nums[i]
			s++
		}
	}
	return s
}
