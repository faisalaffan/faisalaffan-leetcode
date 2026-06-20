# 2690 — Infinite Method Object

## Deskripsi

**Soal:** [2690. Infinite Method Object](https://leetcode.com/problems/infinite-method-object/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2690: Infinite Method Object
// https://leetcode.com/problems/infinite-method-object/
// Difficulty: Easy [Paid]
// Time: O(1) | Space: O(1)
// Note: JavaScript Proxy problem, adapted to Go. Returns an object that returns "methodName" for any method.

import "fmt"

func main() {
	obj := InfiniteMethodObject()
	fmt.Println(obj("abc"))
	fmt.Println(obj("xyz"))
}

func InfiniteMethodObject() func(string) string {
	return func(methodName string) string {
		return methodName
	}
}
```
