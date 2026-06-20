# 3731 — Find Missing Elements

## Deskripsi

**Soal:** [3731. Find Missing Elements](https://leetcode.com/problems/find-missing-elements/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3731: Find Missing Elements
// https://leetcode.com/problems/find-missing-elements/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMissingElements([]int{1, 4, 2, 5}))
	fmt.Println(FindMissingElements([]int{7, 8, 6, 9}))
	fmt.Println(FindMissingElements([]int{5, 1}))
}

// Time: O(n)
// Space: O(n)
func FindMissingElements(nums []int) []int {
  // Membuat map untuk pencarian O(1): key → value
	has := make(map[int]bool)
	mn, mx := nums[0], nums[0]
	for _, v := range nums {
		has[v] = true
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, 0)
	for x := mn + 1; x < mx; x++ {
		if !has[x] {
			ans = append(ans, x)
		}
	}
	return ans
}
```
