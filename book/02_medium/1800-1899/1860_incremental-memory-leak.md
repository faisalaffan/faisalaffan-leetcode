# 1860 — Incremental Memory Leak

## Deskripsi

**Soal:** [1860. Incremental Memory Leak](https://leetcode.com/problems/incremental-memory-leak/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(sqrt(memory1+memory2)), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1860: Incremental Memory Leak
// https://leetcode.com/problems/incremental-memory-leak/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MemLeak(2, 2))
	fmt.Println(MemLeak(8, 11))
	fmt.Println(MemLeak(1, 1))
}

// Time: O(sqrt(memory1+memory2)), Space: O(1)
func MemLeak(memory1 int, memory2 int) []int {
	t := 1
	for memory1 >= t || memory2 >= t {
		if memory1 >= memory2 {
			memory1 -= t
		} else {
			memory2 -= t
		}
		t++
	}
	return []int{t, memory1, memory2}
}
```
