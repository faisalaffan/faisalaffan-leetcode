# 2755 — Deep Merge Of Two Objects

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DeepMergeOfTwoObjects(obj1, obj2 JSONObj) JSONObj
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2755: Deep Merge of Two Objects
// https://leetcode.com/problems/deep-merge-of-two-objects/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type JSONObj map[string]interface{}

func DeepMergeOfTwoObjects(obj1, obj2 JSONObj) JSONObj {
	result := make(JSONObj)
	for k, v := range obj1 {
		result[k] = v
	}
	for k, v2 := range obj2 {
		if v1, ok := result[k]; ok {
			m1, ok1 := v1.(map[string]interface{})
			m2, ok2 := v2.(map[string]interface{})
			if ok1 && ok2 {
				result[k] = DeepMergeOfTwoObjects(JSONObj(m1), JSONObj(m2))
			} else {
				result[k] = v2
			}
		} else {
			result[k] = v2
		}
	}
	return result
}

func main() {
	merged := DeepMergeOfTwoObjects(
		JSONObj{"a": 1, "b": JSONObj{"c": 2}},
		JSONObj{"b": JSONObj{"d": 3}, "e": 4},
	)
	fmt.Println(merged)

	merged2 := DeepMergeOfTwoObjects(
		JSONObj{"x": 1},
		JSONObj{"x": 2},
	)
	fmt.Println(merged2)
}
```
