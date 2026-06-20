# 3084 — Count Substrings Starting And Ending With Given Character

## Deskripsi

**Soal:** [3084. Count Substrings Starting And Ending With Given Character](https://leetcode.com/problems/count-substrings-starting-and-ending-with-given-character/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3084: Count Substrings Starting and Ending with Given Character
// https://leetcode.com/problems/count-substrings-starting-and-ending-with-given-character/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countSubstringsStartingEnding("abada", 'a'))
	fmt.Println(countSubstringsStartingEnding("zzz", 'z'))
}

func countSubstringsStartingEnding(s string, c byte) int64 {
	cnt := int64(0)
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			cnt++
		}
	}
	return cnt * (cnt + 1) / 2
}
```
