# 3110 — Score Of A String

## Deskripsi

**Soal:** [3110. Score Of A String](https://leetcode.com/problems/score-of-a-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3110: Score of a String
// https://leetcode.com/problems/score-of-a-string/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: scoreOfString
	fmt.Println(ScoreOfAString("hello")) // 13
	fmt.Println(ScoreOfAString("zaz"))   // 50
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: scoreOfString
func ScoreOfAString(s string) int {
	score := 0
	for i := 1; i < len(s); i++ {
		diff := int(s[i]) - int(s[i-1])
		if diff < 0 {
			diff = -diff
		}
		score += diff
	}
	return score
}
```
