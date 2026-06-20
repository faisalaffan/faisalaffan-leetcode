# 1476 — Subrectangle Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewSubrectangleQueries(rectangle [][]int) SubrectangleQueries
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1476: Subrectangle Queries
// https://leetcode.com/problems/subrectangle-queries/
// Difficulty: Medium

import "fmt"

type SubrectangleQueries struct {
	rectangle   [][]int
	updates     [][5]int // row1, col1, row2, col2, newValue (lazy)
}

func main() {
	sq := NewSubrectangleQueries([][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}})
	fmt.Println(sq.GetValue(0, 2)) // 1
	sq.UpdateSubrectangle(0, 0, 3, 2, 5)
	fmt.Println(sq.GetValue(0, 2)) // 5
	fmt.Println(sq.GetValue(3, 1)) // 5
	sq.UpdateSubrectangle(3, 0, 3, 2, 10)
	fmt.Println(sq.GetValue(3, 1)) // 10
	fmt.Println(sq.GetValue(0, 2)) // 5

	sq2 := NewSubrectangleQueries([][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}})
	fmt.Println(sq2.GetValue(0, 0)) // 1
	sq2.UpdateSubrectangle(0, 0, 2, 2, 100)
	fmt.Println(sq2.GetValue(0, 0)) // 100
	fmt.Println(sq2.GetValue(2, 2)) // 100
	sq2.UpdateSubrectangle(1, 1, 2, 2, 200)
	fmt.Println(sq2.GetValue(0, 0)) // 100
	fmt.Println(sq2.GetValue(1, 1)) // 200
}

func NewSubrectangleQueries(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

// Time: O(1)
func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	this.updates = append(this.updates, [5]int{row1, col1, row2, col2, newValue})
}

// Time: O(k) where k = number of updates
func (this *SubrectangleQueries) GetValue(row int, col int) int {
	// Check most recent update first
	for i := len(this.updates) - 1; i >= 0; i-- {
		u := this.updates[i]
		if row >= u[0] && row <= u[2] && col >= u[1] && col <= u[3] {
			return u[4]
		}
	}
	return this.rectangle[row][col]
}
```
