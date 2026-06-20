# 0855 — Exam Room

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(n int) ExamRoom
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #855: Exam Room
// https://leetcode.com/problems/exam-room/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

type ExamRoom struct {
	seats []int
	n     int
}

func Constructor(n int) ExamRoom {
	return ExamRoom{seats: []int{}, n: n}
}

func (this *ExamRoom) Seat() int {
	if len(this.seats) == 0 {
		this.seats = append(this.seats, 0)
		return 0
	}

	maxDist := this.seats[0]
	pos := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(this.seats)-1; i++ {
		dist := (this.seats[i+1] - this.seats[i]) / 2
		if dist > maxDist {
			maxDist = dist
			pos = this.seats[i] + dist
		}
	}

	dist := this.n - 1 - this.seats[len(this.seats)-1]
	if dist > maxDist {
		pos = this.n - 1
	}

	idx := sort.SearchInts(this.seats, pos)
	this.seats = append(this.seats, 0)
	copy(this.seats[idx+1:], this.seats[idx:])
	this.seats[idx] = pos

	return pos
}

func (this *ExamRoom) Leave(p int) {
	idx := sort.SearchInts(this.seats, p)
	this.seats = append(this.seats[:idx], this.seats[idx+1:]...)
}

func main() {
	// Test case 1
	obj1 := Constructor(10)
	fmt.Println(obj1.Seat())
	fmt.Println(obj1.Seat())
	fmt.Println(obj1.Seat())
	fmt.Println(obj1.Seat())
	obj1.Leave(4)
	fmt.Println(obj1.Seat())

	fmt.Println("---")

	// Test case 2
	obj2 := Constructor(5)
	fmt.Println(obj2.Seat())
	fmt.Println(obj2.Seat())
	fmt.Println(obj2.Seat())
	fmt.Println(obj2.Seat())
	fmt.Println(obj2.Seat())
}
```
