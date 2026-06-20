# 1018 — Binary Prefix Divisible By 5

## Deskripsi

**Soal:** [1018. Binary Prefix Divisible By 5](https://leetcode.com/problems/binary-prefix-divisible-by-5/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1) (excluding output)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1018: Binary Prefix Divisible By 5
// https://leetcode.com/problems/binary-prefix-divisible-by-5/
// Difficulty: Easy
// Time: O(n) | Space: O(1) (excluding output)

import "fmt"

func main() {
	fmt.Println(prefixesDivBy5([]int{0, 1, 1}))          // [true,false,false]
	fmt.Println(prefixesDivBy5([]int{1, 1, 1}))          // [false,false,false]
	fmt.Println(prefixesDivBy5([]int{0, 1, 1, 1, 1, 1})) // [true,false,false,false,true,false]
}

// LeetCode submission: prefixesDivBy5
func prefixesDivBy5(nums []int) []bool {
  // Membuat slice untuk menyimpan hasil
	ans := make([]bool, len(nums))
	val := 0
	for i, b := range nums {
		val = (val*2 + b) % 5
		ans[i] = val == 0
	}
	return ans
}
```
