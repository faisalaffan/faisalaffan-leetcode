# 2666 — Allow One Function Call

## Deskripsi

**Soal:** [2666. Allow One Function Call](https://leetcode.com/problems/allow-one-function-call/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2666: Allow One Function Call
// https://leetcode.com/problems/allow-one-function-call/
// Difficulty: Easy
// Time: O(1) | Space: O(1)
// Note: JavaScript problem, adapted to Go. Ensures fn is called at most once.

import "fmt"

func main() {
	fn := func(x int) int { return x * 2 }
	onceFn := AllowOneFunctionCall(fn)
	fmt.Println(onceFn(5))
	fmt.Println(onceFn(10))
}

func AllowOneFunctionCall(fn func(int) int) func(int) int {
	called := false
	return func(x int) int {
		if called {
			return 0
		}
		called = true
		return fn(x)
	}
}
```
