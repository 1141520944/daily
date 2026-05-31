package leetcode

/*
* @author: hjz
* @time: 2026/5/31 15:34
 */

func majorityElement(nums []int) int {
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	resc, resn := 0, 0
	for k, v := range m {
		if v > resc {
			resc = v
			resn = k
		}
	}
	return resn
}
