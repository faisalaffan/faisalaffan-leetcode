# 2775 — Undefined To Null

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func UndefinedToNull(obj NullableObj) NullableObj`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2775: Undefined to Null
// https://leetcode.com/problems/undefined-to-null/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type NullableObj map[string]interface{}

func UndefinedToNull(obj NullableObj) NullableObj {
	result := make(NullableObj)
	for k, v := range obj {
		if v == nil {
			result[k] = nil
		} else if m, ok := v.(map[string]interface{}); ok {
			result[k] = UndefinedToNull(m)
		} else if arr, ok := v.([]interface{}); ok {
  // Alokasi slice
			newArr := make([]interface{}, len(arr))
			for i, item := range arr {
				if m2, ok := item.(map[string]interface{}); ok {
					newArr[i] = UndefinedToNull(m2)
				} else {
					newArr[i] = item
				}
			}
			result[k] = newArr
		} else {
			result[k] = v
		}
	}
	return result
}

func main() {
	obj := NullableObj{"a": nil, "b": 42, "c": NullableObj{"d": nil}}
	fmt.Println(UndefinedToNull(obj))

	obj2 := NullableObj{"x": 1}
	fmt.Println(UndefinedToNull(obj2))
}
```
