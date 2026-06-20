# 1422 — Maximum Score After Splitting A String

## Deskripsi

**Soal:** [1422. Maximum Score After Splitting A String](https://leetcode.com/problems/maximum-score-after-splitting-a-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxScore(s string) int`

## Solusi Go

```go
package main

// LeetCode #1422: Maximum Score After Splitting a String
// https://leetcode.com/problems/maximum-score-after-splitting-a-string/
// Difficulty: Easy
//
// LeetCode submission: func maxScore(s string) int

import "fmt"

func main() {
	fmt.Println(MaximumScoreAfterSplittingAString("011101")) // 5
	fmt.Println(MaximumScoreAfterSplittingAString("00111"))  // 5
	fmt.Println(MaximumScoreAfterSplittingAString("1111"))   // 3
}

// Time: O(n), Space: O(1)
func MaximumScoreAfterSplittingAString(s string) int {
	ones := 0
	for _, ch := range s {
		if ch == '1' {
			ones++
		}
	}
	zeros, maxScore := 0, 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones--
		}
		if zeros+ones > maxScore {
			maxScore = zeros + ones
		}
	}
	return maxScore
}
```
