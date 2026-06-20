# 2026 — Low Quality Problems

## Deskripsi

**Soal:** [2026. Low Quality Problems](https://leetcode.com/problems/low-quality-problems/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2026: Low-Quality Problems
// https://leetcode.com/problems/low-quality-problems/
// Difficulty: Easy [Paid] (SQL)

import (
	"fmt"
	"sort"
)

func main() {
	// Problems: (problem_id, likes, dislikes)
	problems := [][3]string{{"1", "10", "5"}, {"2", "3", "10"}, {"3", "2", "2"}}
	fmt.Println(LowQualityProblems(problems)) // [2]
}

// Time: O(n log n), Space: O(n)
func LowQualityProblems(problems [][3]string) []int {
	var result []int
	for _, p := range problems {
		id := 0
		for _, c := range p[0] {
			id = id*10 + int(c-'0')
		}
		likes := 0
		for _, c := range p[1] {
			likes = likes*10 + int(c-'0')
		}
		dislikes := 0
		for _, c := range p[2] {
			dislikes = dislikes*10 + int(c-'0')
		}

		// Low-quality: likes / (likes + dislikes) < 0.6
		total := likes + dislikes
		if total > 0 && likes*5 < total*3 {
			result = append(result, id)
		}
	}
	sort.Ints(result)
	return result
}
```
