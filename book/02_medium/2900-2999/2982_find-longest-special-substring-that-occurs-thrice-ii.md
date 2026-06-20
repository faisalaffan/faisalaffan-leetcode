# 2982 — Find Longest Special Substring That Occurs Thrice Ii

## Deskripsi

**Soal:** [2982. Find Longest Special Substring That Occurs Thrice Ii](https://leetcode.com/problems/find-longest-special-substring-that-occurs-thrice-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2982: Find Longest Special Substring That Occurs Thrice II
// https://leetcode.com/problems/find-longest-special-substring-that-occurs-thrice-ii/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maximumLength2982("aaaa"))
	fmt.Println(maximumLength2982("abcdef"))
	fmt.Println(maximumLength2982("abcaba"))
}

func maximumLength2982(s string) int {
	n := len(s)
	l, r := 0, n
	check := func(x int) bool {
		cnt := [26]int{}
		for i := 0; i < n; {
			j := i + 1
			for j < n && s[j] == s[i] {
				j++
			}
			k := s[i] - 'a'
			add := j - i - x + 1
			if add > 0 {
				cnt[k] += add
			}
			if cnt[k] >= 3 {
				return true
			}
			i = j
		}
		return false
	}
	for l < r {
		mid := (l + r + 1) >> 1
		if check(mid) {
			l = mid
		} else {
			r = mid - 1
		}
	}
	if l == 0 {
		return -1
	}
	return l
}
```
