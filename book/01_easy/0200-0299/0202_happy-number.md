# 0202 — Happy Number

## Deskripsi

**Soal:** [0202. Happy Number](https://leetcode.com/problems/happy-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func IsHappy(n int) bool`

## Solusi Go

```go
package main

// LeetCode #202: Happy Number
// https://leetcode.com/problems/happy-number/
// Difficulty: Easy

import "fmt"

// Time: O(log n) | Space: O(1)
func IsHappy(n int) bool {
	next := func(x int) int {
		sum := 0
		for x > 0 {
			d := x % 10
			sum += d * d
			x /= 10
		}
		return sum
	}
	slow, fast := n, next(n)
	for fast != 1 && slow != fast {
		slow = next(slow)
		fast = next(next(fast))
	}
	return fast == 1
}

func main() {
	fmt.Println(IsHappy(19))
	fmt.Println(IsHappy(2))
}
```
