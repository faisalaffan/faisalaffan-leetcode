# 1446 — Consecutive Characters

## Deskripsi

**Soal:** [1446. Consecutive Characters](https://leetcode.com/problems/consecutive-characters/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxPower(s string) int`

## Solusi Go

```go
package main

// LeetCode #1446: Consecutive Characters
// https://leetcode.com/problems/consecutive-characters/
// Difficulty: Easy
//
// LeetCode submission: func maxPower(s string) int

import "fmt"

func main() {
	fmt.Println(ConsecutiveCharacters("leetcode")) // 2
	fmt.Println(ConsecutiveCharacters("abbcccddddeeeeedcba")) // 5
	fmt.Println(ConsecutiveCharacters("triplepillooooow")) // 5
}

// Time: O(n), Space: O(1)
func ConsecutiveCharacters(s string) int {
	ans, cur := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
			if cur > ans {
				ans = cur
			}
		} else {
			cur = 1
		}
	}
	return ans
}
```
