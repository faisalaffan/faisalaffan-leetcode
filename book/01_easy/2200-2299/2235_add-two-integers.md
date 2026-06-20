# 2235 — Add Two Integers

## Deskripsi

**Soal:** [2235. Add Two Integers](https://leetcode.com/problems/add-two-integers/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2235: Add Two Integers
// https://leetcode.com/problems/add-two-integers/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(AddTwoIntegers(12, 5))   // 17
	fmt.Println(AddTwoIntegers(-10, 4))  // -6
}

// Time: O(1), Space: O(1)
func AddTwoIntegers(num1 int, num2 int) int {
	return num1 + num2
}
```
