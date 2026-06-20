# 0412 — Fizz Buzz

## Deskripsi

**Soal:** [0412. Fizz Buzz](https://leetcode.com/problems/fizz-buzz/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func FizzBuzz(n int) []string`

## Solusi Go

```go
package main

// LeetCode #412: Fizz Buzz
// https://leetcode.com/problems/fizz-buzz/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

// Time: O(n), Space: O(n)
func FizzBuzz(n int) []string {
  // Membuat slice untuk menyimpan hasil
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}

func main() {
	fmt.Println(FizzBuzz(3))
	fmt.Println(FizzBuzz(5))
	fmt.Println(FizzBuzz(15))
}
```
