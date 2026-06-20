# 3827 — Count Monobit Integers

## Deskripsi

**Soal:** [3827. Count Monobit Integers](https://leetcode.com/problems/count-monobit-integers/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3827: Count Monobit Integers
// https://leetcode.com/problems/count-monobit-integers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CountMonobitIntegers(1))
	fmt.Println(CountMonobitIntegers(4))
	fmt.Println(CountMonobitIntegers(0))
}

// Time: O(log n)
// Space: O(1)
func CountMonobitIntegers(n int) int {
	ans := 1 // 0 is monobit (all zeros)
	x := 1   // 2^1 - 1 = 1 (all ones)
	for x <= n {
		ans++
		x = x*2 + 1
	}
	return ans
}
```
