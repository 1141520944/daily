package leetcode

/*
* @author: hjz
* @time: 2026/5/31 14:07
* @function:
* 倒比较每个数字，大的放到末尾，剩余的放到nums1的末尾
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	r1, r2 := m-1, n-1
	index := len(nums1) - 1
	for r1 >= 0 && r2 >= 0 {
		if nums1[r1] > nums2[r2] {
			nums1[index] = nums1[r1]
			r1--
		} else {
			nums1[index] = nums2[r2]
			r2--
		}
		index--
	}
	for r2 >= 0 {
		nums1[index] = nums2[r2]
		r2--
		index--
	}
}
