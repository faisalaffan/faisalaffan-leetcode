# 0881 — Boats To Save People

## Deskripsi

**Soal:** [0881. Boats To Save People](https://leetcode.com/problems/boats-to-save-people/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(log n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #881: Boats to Save People
// https://leetcode.com/problems/boats-to-save-people/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(BoatsToSavePeople([]int{1, 2}, 3))
	fmt.Println(BoatsToSavePeople([]int{3, 2, 2, 1}, 3))
	fmt.Println(BoatsToSavePeople([]int{3, 5, 3, 4}, 5))
}

// Time: O(n log n) | Space: O(log n)
func BoatsToSavePeople(people []int, limit int) int {
	sort.Ints(people)
	left, right := 0, len(people)-1
	ans := 0

	for left <= right {
		if people[left]+people[right] <= limit {
			left++
		}
		right--
		ans++
	}

	return ans
}
```
