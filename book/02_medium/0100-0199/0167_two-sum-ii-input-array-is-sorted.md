# 0167 — Two Sum Ii Input Array Is Sorted

## Deskripsi

**Soal:** [0167. Two Sum Ii Input Array Is Sorted](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func twoSum(numbers []int, target int) []int`

## Solusi Go

```go
package main

// LeetCode #167: Two Sum II - Input Array Is Sorted
// https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func twoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

  // Loop two-pointer: kiri vs kanan
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{2, 3, 4}, 6))
	fmt.Println(twoSum([]int{-1, 0}, -1))
}
```
