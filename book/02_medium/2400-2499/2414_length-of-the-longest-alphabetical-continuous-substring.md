# 2414 — Length Of The Longest Alphabetical Continuous Substring

## Deskripsi

**Soal:** [2414. Length Of The Longest Alphabetical Continuous Substring](https://leetcode.com/problems/length-of-the-longest-alphabetical-continuous-substring/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2414: Length of the Longest Alphabetical Continuous Substring
// https://leetcode.com/problems/length-of-the-longest-alphabetical-continuous-substring/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Scan, count consecutive chars where s[i] == s[i-1] + 1.

import "fmt"

func main() {
	fmt.Println(longestContinuousSubstring("abacaba")) // 2 ("ab")
	fmt.Println(longestContinuousSubstring("abcde"))   // 5
}

func longestContinuousSubstring(s string) int {
	ans, cur := 0, 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i] == s[i-1]+1 {
			cur++
		} else {
			cur = 1
		}
		if cur > ans {
			ans = cur
		}
	}
	return ans
}
```
