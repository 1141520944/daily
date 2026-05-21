package leetcode

import "sort"

/*
* @author: hjz
* @time: 2026/5/21 23:06
 */
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := [][]int{}
	l, r := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > r {
			res = append(res, []int{l, r})
			l, r = intervals[i][0], intervals[i][1]
		} else {
			r = max(r, intervals[i][1])
		}
	}
	res = append(res, []int{l, r})
	return res
}
