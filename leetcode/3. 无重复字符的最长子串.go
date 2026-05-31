package leetcode

/*
* @author: hjz
* @time: 2026/5/31 16:23
 */
func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	res := 0
	l, r := 0, 0

	for r < len(s) {
		m[s[r]]++
		for m[s[r]] > 1 {
			m[s[l]]--
			l++
		}
		res = max(res, r-l+1)
		r++
	}
	return res
}
