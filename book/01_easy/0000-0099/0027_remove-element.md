# 0027 — Remove Element

## Deskripsi

**Soal:** [0027. Remove Element](https://leetcode.com/problems/remove-element/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func RemoveElement(nums []int, val int) int`

## Solusi Go

```go
package main

// LeetCode #27: Remove Element
// https://leetcode.com/problems/remove-element/
// Difficulty: Easy

import "fmt"

// Time: O(n) | Space: O(1)
func RemoveElement(nums []int, val int) int {
	k := 0
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}

func main() {
	n1 := []int{3, 2, 2, 3}
	fmt.Println(RemoveElement(n1, 3), n1)
	n2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(RemoveElement(n2, 2), n2)
}
```
