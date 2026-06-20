# 2293 — Min Max Game

## Deskripsi

**Soal:** [2293. Min Max Game](https://leetcode.com/problems/min-max-game/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2293: Min Max Game
// https://leetcode.com/problems/min-max-game/
// Difficulty: Easy
// Time O(n) | Space O(n)

import "fmt"

func main() {
	fmt.Println(MinMaxGame([]int{1, 3, 5, 2, 4, 8, 2, 2})) // 1
	fmt.Println(MinMaxGame([]int{3}))                        // 3
}

func MinMaxGame(nums []int) int {
	for len(nums) > 1 {
  // Membuat slice untuk menyimpan hasil
		next := make([]int, len(nums)/2)
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(next); i++ {
			if i%2 == 0 {
				next[i] = min(nums[2*i], nums[2*i+1])
			} else {
				next[i] = max(nums[2*i], nums[2*i+1])
			}
		}
		nums = next
	}
	return nums[0]
}
```
