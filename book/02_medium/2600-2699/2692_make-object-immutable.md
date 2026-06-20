# 2692 — Make Object Immutable

## Deskripsi

**Soal:** [2692. Make Object Immutable](https://leetcode.com/problems/make-object-immutable/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func NewImmutable(data map[string]any) *ImmutableMap`

## Solusi Go

```go
package main

// LeetCode #2692: Make Object Immutable
// https://leetcode.com/problems/make-object-immutable/
// Difficulty: Medium [Paid] (JS problem)
// Time: O(1) | Space: O(1)

import "fmt"

type ImmutableMap struct {
	data map[string]any
}

func NewImmutable(data map[string]any) *ImmutableMap {
  // Membuat map untuk pencarian O(1): key → value
	copied := make(map[string]any)
	for k, v := range data {
		copied[k] = v
	}
	return &ImmutableMap{data: copied}
}

func (im *ImmutableMap) Get(key string) (any, bool) {
	val, ok := im.data[key]
	return val, ok
}

func (im *ImmutableMap) Set(key string, val any) {
	// Immutable: do nothing (or panic in JS-like implementation)
	// In Go, we just don't modify
}

func main() {
	im := NewImmutable(map[string]any{"a": 1, "b": 2})

	// Test case 1
	val, ok := im.Get("a")
	fmt.Println("Test 1:", val, ok)
	// Expected: 1 true

	// Test case 2
	im.Set("c", 3)
	_, ok2 := im.Get("c")
	fmt.Println("Test 2:", ok2)
	// Expected: false (immutable)

	// Test case 3
	_, ok3 := im.Get("z")
	fmt.Println("Test 3:", ok3)
	// Expected: false
}
```
