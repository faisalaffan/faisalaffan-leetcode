# 2823 — Deep Object Filter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DeepObjectFilter(obj FilterObj, predicate func(string, interface{}) bool) FilterObj
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2823: Deep Object Filter
// https://leetcode.com/problems/deep-object-filter/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type FilterObj map[string]interface{}

func DeepObjectFilter(obj FilterObj, predicate func(string, interface{}) bool) FilterObj {
	result := make(FilterObj)
	for k, v := range obj {
		if !predicate(k, v) {
			continue
		}
		if m, ok := v.(map[string]interface{}); ok {
			filtered := DeepObjectFilter(m, predicate)
			if len(filtered) > 0 {
				result[k] = filtered
			}
		} else {
			result[k] = v
		}
	}
	return result
}

func main() {
	obj := FilterObj{
		"a": 1,
		"b": 0,
		"c": FilterObj{"d": 2, "e": 0},
	}
	// Filter out values that are 0
	result := DeepObjectFilter(obj, func(k string, v interface{}) bool {
		if num, ok := v.(int); ok && num == 0 {
			return false
		}
		return true
	})
	fmt.Println(result)

	// Filter keys starting with 'a'
	result2 := DeepObjectFilter(obj, func(k string, v interface{}) bool {
		return k >= "b"
	})
	fmt.Println(result2)
}
```
