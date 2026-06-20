# 3945 — Digit Frequency Score

## Deskripsi

**Soal:** [3945. Digit Frequency Score](https://leetcode.com/problems/digit-frequency-score/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3945: Digit Frequency Score
// https://leetcode.com/problems/digit-frequency-score/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(DigitFrequencyScore(122))
	fmt.Println(DigitFrequencyScore(101))
}

// Time: O(log n)
// Space: O(1)
func DigitFrequencyScore(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}
```
