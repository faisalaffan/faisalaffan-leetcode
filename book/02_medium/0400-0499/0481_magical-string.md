# 0481 — Magical String

## Deskripsi

**Soal:** [0481. Magical String](https://leetcode.com/problems/magical-string/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #481: Magical String
// https://leetcode.com/problems/magical-string/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(MagicalString(6))
	fmt.Println(MagicalString(1))
}

func MagicalString(n int) int {
	if n <= 0 {
		return 0
	}
	if n <= 3 {
		return 1
	}

  // Membuat slice untuk menyimpan hasil
	s := make([]int, n)
	s[0], s[1], s[2] = 1, 2, 2
	count := 1
	writeIdx := 3
	readIdx := 2

	for writeIdx < n {
		val := 3 - s[writeIdx-1] // toggle between 1 and 2
		countTimes := s[readIdx]
		for i := 0; i < countTimes && writeIdx < n; i++ {
			s[writeIdx] = val
			if val == 1 {
				count++
			}
			writeIdx++
		}
		readIdx++
	}

	return count
}
```
