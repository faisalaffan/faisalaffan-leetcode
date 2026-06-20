# 0231 — Power Of Two

## Deskripsi

**Soal:** [0231. Power Of Two](https://leetcode.com/problems/power-of-two/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func IsPowerOfTwo(n int) bool`

## Solusi Go

```go
package main

// LeetCode #231: Power of Two
// https://leetcode.com/problems/power-of-two/
// Difficulty: Easy

import "fmt"

// Time: O(1) | Space: O(1)
func IsPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func main() {
	fmt.Println(IsPowerOfTwo(1))
	fmt.Println(IsPowerOfTwo(16))
	fmt.Println(IsPowerOfTwo(3))
}
```
