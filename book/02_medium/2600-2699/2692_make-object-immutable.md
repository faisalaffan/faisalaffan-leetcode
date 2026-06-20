# 2692 — Make Object Immutable

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan tabel database dan diminta untuk menulis query SQL. Karena repo ini menggunakan Go, query SQL disimulasikan dengan struktur data Go (map untuk grouping, slice untuk sorting, struct untuk representasi row).

Soal tipe ini menguji kemampuanmu menganalisis data relasional — seperti yang kamu lakukan dengan SQL di pekerjaan backend sehari-hari.

**Konsep kunci:** GROUP BY, JOIN, aggregate (SUM, COUNT, AVG), window function (RANK, ROW_NUMBER), HAVING.

**Fungsi yang perlu kamu implementasikan:**
```go
func NewImmutable(data map[string]any) *ImmutableMap
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
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
