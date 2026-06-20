# 2878 — Get The Size Of A Dataframe

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func GetTheSizeOfADataframe(df [][]int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2878: Get the Size of a DataFrame
// https://leetcode.com/problems/get-the-size-of-a-dataframe/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we return the dimensions [rows, cols].

import "fmt"

func main() {
	// LeetCode name: getDataframeSize
	fmt.Println(GetTheSizeOfADataframe([][]int{{1, 15}, {2, 11}, {3, 11}, {4, 20}})) // [4 2]
	fmt.Println(GetTheSizeOfADataframe([][]int{{5, 25}}))                             // [1 2]
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: getDataframeSize
func GetTheSizeOfADataframe(df [][]int) []int {
	if len(df) == 0 {
		return []int{0, 0}
	}
	return []int{len(df), len(df[0])}
}
```
