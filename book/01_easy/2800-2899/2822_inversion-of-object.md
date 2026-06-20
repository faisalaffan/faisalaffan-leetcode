# 2822 — Inversion Of Object

## Deskripsi

**Soal:** [2822. Inversion Of Object](https://leetcode.com/problems/inversion-of-object/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2822: Inversion of Object
// https://leetcode.com/problems/inversion-of-object/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Swaps keys and values of a map.

import "fmt"

func main() {
	input := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println(InversionOfObject(input))
}

func InversionOfObject(obj map[string]int) map[int]string {
  // Membuat map untuk pencarian O(1): key → value
	result := make(map[int]string, len(obj))
	for k, v := range obj {
		result[v] = k
	}
	return result
}
```
