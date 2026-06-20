# 1287 — Element Appearing More Than 25 In Sorted Array

## Deskripsi

**Soal:** [1287. Element Appearing More Than 25 In Sorted Array](https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1287: Element Appearing More Than 25% In Sorted Array
// https://leetcode.com/problems/element-appearing-more-than-25-in-sorted-array/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(findSpecialInteger([]int{1, 2, 2, 6, 6, 6, 6, 7, 10})) // 6
	fmt.Println(findSpecialInteger([]int{1, 1}))                       // 1
}

// LeetCode submission: findSpecialInteger
func findSpecialInteger(arr []int) int {
	target := len(arr) / 4
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(arr)-target; i++ {
		if arr[i] == arr[i+target] {
			return arr[i]
		}
	}
	return arr[0]
}
```
