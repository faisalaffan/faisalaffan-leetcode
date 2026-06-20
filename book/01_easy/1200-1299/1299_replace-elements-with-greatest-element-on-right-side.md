# 1299 — Replace Elements With Greatest Element On Right Side

## Deskripsi

**Soal:** [1299. Replace Elements With Greatest Element On Right Side](https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) excluding output

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1299: Replace Elements with Greatest Element on Right Side
// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/
// Difficulty: Easy
// Time: O(n) | Space: O(1) excluding output

import "fmt"

func main() {
	fmt.Println(replaceElements([]int{17, 18, 5, 4, 6, 1})) // [18,6,6,6,1,-1]
	fmt.Println(replaceElements([]int{400}))                 // [-1]
}

// LeetCode submission: replaceElements
func replaceElements(arr []int) []int {
  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(arr))
	maxRight := -1
	for i := len(arr) - 1; i >= 0; i-- {
		ans[i] = maxRight
		if arr[i] > maxRight {
			maxRight = arr[i]
		}
	}
	return ans
}
```
