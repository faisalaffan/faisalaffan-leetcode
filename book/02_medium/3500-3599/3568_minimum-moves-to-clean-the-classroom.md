# 3568 — Minimum Moves To Clean The Classroom

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumMovesToCleanTheClassroom(room [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3568: Minimum Moves to Clean the Classroom
// https://leetcode.com/problems/minimum-moves-to-clean-the-classroom/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	room := [][]int{{0, 0}, {0, 0}}
	fmt.Println("Test 1:", MinimumMovesToCleanTheClassroom(room))
	// Test case 2
	room2 := [][]int{{1, 0}, {0, 0}}
	fmt.Println("Test 2:", MinimumMovesToCleanTheClassroom(room2))
	// Test case 3
	room3 := [][]int{{1, 1, 1}, {1, 0, 1}}
	fmt.Println("Test 3:", MinimumMovesToCleanTheClassroom(room3))
}

func MinimumMovesToCleanTheClassroom(room [][]int) int {
	m, n := len(room), len(room[0])
	moves := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if room[i][j] == 1 {
				moves++
			}
		}
	}
	return moves
}
```
