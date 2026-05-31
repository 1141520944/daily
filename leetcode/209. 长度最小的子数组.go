package leetcode

import "math"

/*
* @author: hjz
* @time: 2026/5/31 15:42
 */
// [2,3,1,2,4,3]
func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := 0
	res := math.MaxInt32

	for r < len(nums) {
		sum += nums[r] // 扩大窗口

		for sum >= target { // 缩小窗口
			res = min(res, r-l+1)
			sum -= nums[l]
			l++
		}
		r++
	}

	if res == math.MaxInt32 {
		return 0
	}
	return res
}
