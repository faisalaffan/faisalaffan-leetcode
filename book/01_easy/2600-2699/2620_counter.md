# 2620 — Counter

## Deskripsi

**Soal:** [2620. Counter](https://leetcode.com/problems/counter/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2620: Counter
// https://leetcode.com/problems/counter/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript closure problem, adapted to Go. Returns a counter function.

import "fmt"

func main() {
	counter := counter(10)
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func counter(n int) func() int {
	return func() int {
		n++
		return n - 1
	}
}
```
