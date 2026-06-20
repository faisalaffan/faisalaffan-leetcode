# 2775 — Undefined To Null

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func UndefinedToNull(obj NullableObj) NullableObj
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Alokasi slice integer
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
