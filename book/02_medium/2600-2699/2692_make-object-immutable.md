# 2692 — Make Object Immutable

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func NewImmutable(data map[string]any) *ImmutableMap`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
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
