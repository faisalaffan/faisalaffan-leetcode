# 1243 — Array Transformation

## Deskripsi

**Soal:** [1243. Array Transformation](https://leetcode.com/problems/array-transformation/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n^2) worst case  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1243: Array Transformation
// https://leetcode.com/problems/array-transformation/
// Difficulty: Easy [Paid]
// Time: O(n^2) worst case | Space: O(n)

import "fmt"

func main() {
	fmt.Println(transformArray([]int{6, 2, 3, 4})) // [6,3,3,4]
	fmt.Println(transformArray([]int{1, 6, 3, 4, 3, 5})) // [1,4,4,4,4,5]
}

// LeetCode submission: transformArray
func transformArray(arr []int) []int {
	if len(arr) <= 2 {
		return append([]int{}, arr...)
	}
	for {
		changed := false
  // Membuat slice untuk menyimpan hasil
		next := make([]int, len(arr))
		copy(next, arr)
		for i := 1; i < len(arr)-1; i++ {
			if arr[i] < arr[i-1] && arr[i] < arr[i+1] {
				next[i]++
				changed = true
			} else if arr[i] > arr[i-1] && arr[i] > arr[i+1] {
				next[i]--
				changed = true
			}
		}
		if !changed {
			break
		}
		arr = next
	}
	return arr
}
```
