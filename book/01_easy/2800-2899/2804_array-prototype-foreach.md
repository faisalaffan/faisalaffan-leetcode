# 2804 — Array Prototype Foreach

## Deskripsi

**Soal:** [2804. Array Prototype Foreach](https://leetcode.com/problems/array-prototype-foreach/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2804: Array Prototype ForEach
// https://leetcode.com/problems/array-prototype-foreach/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(1)
// Note: JS problem, adapted to Go. Applies callback to each element.

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	sum := 0
	ArrayPrototypeForeach(nums, func(v int) {
		sum += v
	})
	fmt.Println(sum)
}

func ArrayPrototypeForeach(arr []int, callback func(int)) {
	for _, v := range arr {
		callback(v)
	}
}
```
