# 1389 — Create Target Array In The Given Order

## Deskripsi

**Soal:** [1389. Create Target Array In The Given Order](https://leetcode.com/problems/create-target-array-in-the-given-order/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func createTargetArray(nums []int, index []int) []int`

## Solusi Go

```go
package main

// LeetCode #1389: Create Target Array in the Given Order
// https://leetcode.com/problems/create-target-array-in-the-given-order/
// Difficulty: Easy
//
// LeetCode submission: func createTargetArray(nums []int, index []int) []int

import "fmt"

func main() {
	fmt.Println(CreateTargetArrayInTheGivenOrder([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1})) // [0 4 1 3 2]
	fmt.Println(CreateTargetArrayInTheGivenOrder([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0})) // [0 1 2 3 4]
}

// Time: O(n^2), Space: O(n)
func CreateTargetArrayInTheGivenOrder(nums []int, index []int) []int {
  // Membuat slice untuk menyimpan hasil
	res := make([]int, 0, len(nums))
	for i, idx := range index {
		res = append(res[:idx], append([]int{nums[i]}, res[idx:]...)...)
	}
	return res
}
```
