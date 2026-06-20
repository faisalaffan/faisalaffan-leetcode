# 0495 — Teemo Attacking

## Deskripsi

**Soal:** [0495. Teemo Attacking](https://leetcode.com/problems/teemo-attacking/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func TeemoAttacking(timeSeries []int, duration int) int`

## Solusi Go

```go
package main

// LeetCode #495: Teemo Attacking
// https://leetcode.com/problems/teemo-attacking/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func TeemoAttacking(timeSeries []int, duration int) int {
	total := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(timeSeries)-1; i++ {
		total += min(duration, timeSeries[i+1]-timeSeries[i])
	}
	if len(timeSeries) > 0 {
		total += duration
	}
	return total
}

func main() {
	fmt.Println(TeemoAttacking([]int{1, 4}, 2))
	fmt.Println(TeemoAttacking([]int{1, 2}, 2))
}
```
