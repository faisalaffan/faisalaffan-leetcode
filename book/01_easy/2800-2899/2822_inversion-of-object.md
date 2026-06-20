# 2822 — Inversion Of Object

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func InversionOfObject(obj map[string]int) map[int]string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // HashMap: O(1) lookup
	result := make(map[int]string, len(obj))
	for k, v := range obj {
		result[v] = k
	}
	return result
}
```
