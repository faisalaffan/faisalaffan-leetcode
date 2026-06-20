# 2794 — Create Object From Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func CreateObjectFromTwoArrays(keys []string, values []int) map[string]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2794: Create Object from Two Arrays
// https://leetcode.com/problems/create-object-from-two-arrays/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Creates a map from keys and values arrays.

import "fmt"

func main() {
	fmt.Println(CreateObjectFromTwoArrays([]string{"a", "b", "c"}, []int{1, 2, 3}))
	fmt.Println(CreateObjectFromTwoArrays([]string{"x"}, []int{10}))
}

func CreateObjectFromTwoArrays(keys []string, values []int) map[string]int {
  // HashMap: O(1) lookup
	result := make(map[string]int, len(keys))
	for i, k := range keys {
		if i < len(values) {
			result[k] = values[i]
		}
	}
	return result
}
```
