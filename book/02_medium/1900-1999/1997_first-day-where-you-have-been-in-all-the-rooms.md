# 1997 — First Day Where You Have Been In All The Rooms

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FirstDayWhereYouHaveBeenInAllTheRooms(nextVisit []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1997: First Day Where You Have Been in All the Rooms
// https://leetcode.com/problems/first-day-where-you-have-been-in-all-the-rooms/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(FirstDayWhereYouHaveBeenInAllTheRooms([]int{0, 0}))
	fmt.Println(FirstDayWhereYouHaveBeenInAllTheRooms([]int{0, 1, 2, 0}))
	fmt.Println(FirstDayWhereYouHaveBeenInAllTheRooms([]int{0, 0, 2}))
}

// Time: O(n), Space: O(n)
func FirstDayWhereYouHaveBeenInAllTheRooms(nextVisit []int) int {
	const mod = 1_000_000_007
	n := len(nextVisit)
  // Alokasi slice integer
	s := make([]int, n)

	for i := 0; i < n-1; i++ {
		j := nextVisit[i]
		s[i+1] = (s[i]*2 - s[j] + 2) % mod
		if s[i+1] < 0 {
			s[i+1] += mod
		}
	}

	return s[n-1]
}
```
