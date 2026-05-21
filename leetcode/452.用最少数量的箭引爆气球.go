package leetcode

import "sort"

/*
* @author: hjz
* @time: 2026/5/21 22:56
 */

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	arrowPos := points[0][1] // 第一支箭射在第一个气球的右端点

	for i := 1; i < len(points); i++ {
		if points[i][0] > arrowPos {
			arrows++
			arrowPos = points[i][1]
		}
	}
	return arrows
}
