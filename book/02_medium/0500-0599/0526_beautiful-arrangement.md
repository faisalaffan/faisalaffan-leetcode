# 0526 — Beautiful Arrangement

## Deskripsi

**Soal:** [0526. Beautiful Arrangement](https://leetcode.com/problems/beautiful-arrangement/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(k) where k = number of valid permutations  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #526: Beautiful Arrangement
// https://leetcode.com/problems/beautiful-arrangement/
// Difficulty: Medium
// Time: O(k) where k = number of valid permutations
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(CountArrangement(2))
	fmt.Println(CountArrangement(1))
}

func CountArrangement(n int) int {
  // Membuat slice untuk menyimpan hasil
	used := make([]bool, n+1)
	count := 0

	var backtrack func(pos int)
	backtrack = func(pos int) {
		if pos > n {
			count++
			return
		}
		for i := 1; i <= n; i++ {
			if !used[i] && (i%pos == 0 || pos%i == 0) {
				used[i] = true
				backtrack(pos + 1)
				used[i] = false
			}
		}
	}

	backtrack(1)
	return count
}
```
