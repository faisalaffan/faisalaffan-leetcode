# 2880 — Select Data

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func SelectData(df [][]int, studentID int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2880: Select Data
// https://leetcode.com/problems/select-data/
// Difficulty: Easy
//
// Note: This is a Pandas-only problem on LeetCode (Python).
// In Go, we select rows by student_id and return the age values.

import "fmt"

func main() {
	// LeetCode name: selectData
	fmt.Println(SelectData([][]int{{101, 15}, {102, 11}, {103, 11}}, 101)) // [15]
	fmt.Println(SelectData([][]int{{101, 20}, {102, 22}, {101, 21}}, 101)) // [20, 21]
}

// Time: O(n) | Space: O(n)
// LeetCode submission name: selectData
func SelectData(df [][]int, studentID int) []int {
	ages := []int{}
	for _, row := range df {
		if row[0] == studentID {
			ages = append(ages, row[1])
		}
	}
	return ages
}
```
