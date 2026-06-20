# 2495 — Number Of Subarrays Having Even Product

## Deskripsi

**Soal:** [2495. Number Of Subarrays Having Even Product](https://leetcode.com/problems/number-of-subarrays-having-even-product/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2495: Number of Subarrays Having Even Product
// https://leetcode.com/problems/number-of-subarrays-having-even-product/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Product is even if at least one even element.
// Count subarrays from last even position.

import "fmt"

func main() {
	fmt.Println(evenProduct([]int{1, 2, 3})) // 3 ([2], [1,2], [2,3], [1,2,3])
	fmt.Println(evenProduct([]int{1, 3, 5})) // 0
}

func evenProduct(nums []int) int64 {
	var ans int64
	lastEven := -1
	for i, v := range nums {
		if v%2 == 0 {
			lastEven = i
		}
		if lastEven != -1 {
			ans += int64(lastEven + 1)
		}
	}
	return ans
}
```
