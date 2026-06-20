# 2794 — Create Object From Two Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CreateObjectFromTwoArrays(keys []string, values []int) map[string]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Membuat map (HashMap) — pencarian O(1)
	result := make(map[string]int, len(keys))
	for i, k := range keys {
		if i < len(values) {
			result[k] = values[i]
		}
	}
	return result
}
```
