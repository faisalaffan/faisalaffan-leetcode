# 2053 — Kth Distinct String In An Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func KthDistinctStringInAnArray(arr []string, k int) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2053: Kth Distinct String in an Array
// https://leetcode.com/problems/kth-distinct-string-in-an-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(KthDistinctStringInAnArray([]string{"d", "b", "c", "b", "c", "a"}, 2)) // "a"
	fmt.Println(KthDistinctStringInAnArray([]string{"aaa", "aa", "a"}, 1))              // "aaa"
	fmt.Println(KthDistinctStringInAnArray([]string{"a", "b", "a"}, 3))                 // ""
}

// Time: O(n), Space: O(n)
func KthDistinctStringInAnArray(arr []string, k int) string {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[string]int)
	for _, s := range arr {
		freq[s]++
	}

	idx := 1
	for _, s := range arr {
		if freq[s] == 1 {
			if idx == k {
				return s
			}
			idx++
		}
	}
	return ""
}
```
