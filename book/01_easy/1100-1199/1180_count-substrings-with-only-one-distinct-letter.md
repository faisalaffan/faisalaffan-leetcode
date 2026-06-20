# 1180 — Count Substrings With Only One Distinct Letter

## Deskripsi

**Soal:** [1180. Count Substrings With Only One Distinct Letter](https://leetcode.com/problems/count-substrings-with-only-one-distinct-letter/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1180: Count Substrings with Only One Distinct Letter
// https://leetcode.com/problems/count-substrings-with-only-one-distinct-letter/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countLetters("aaaba")) // 8
	fmt.Println(countLetters("aaaaaaaaaa")) // 55
}

// LeetCode submission: countLetters
func countLetters(s string) int {
	ans := 0
	for i, n := 0, len(s); i < n; {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		L := j - i
		ans += L * (L + 1) / 2
		i = j
	}
	return ans
}
```
