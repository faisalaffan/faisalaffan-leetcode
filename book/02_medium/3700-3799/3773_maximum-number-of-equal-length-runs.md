# 3773 — Maximum Number Of Equal Length Runs

## Deskripsi

**Soal:** [3773. Maximum Number Of Equal Length Runs](https://leetcode.com/problems/maximum-number-of-equal-length-runs/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func maximumNumberOfEqualLengthRuns(s string) int`

## Solusi Go

```go
package main

// LeetCode #3773: Maximum Number of Equal Length Runs
// https://leetcode.com/problems/maximum-number-of-equal-length-runs/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func maximumNumberOfEqualLengthRuns(s string) int {
  // Membuat map untuk pencarian O(1): key → value
	cnt := make(map[int]int)
	maxCount := 0
	n := len(s)
	i := 0
	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		runLen := j - i
		cnt[runLen]++
		if cnt[runLen] > maxCount {
			maxCount = cnt[runLen]
		}
		i = j
	}
	return maxCount
}

func main() {
	fmt.Println(maximumNumberOfEqualLengthRuns("hello"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aaabaaa"))
	fmt.Println(maximumNumberOfEqualLengthRuns("aabbcc"))
}
```
