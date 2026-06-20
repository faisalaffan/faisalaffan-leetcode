# 0729 — My Calendar I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConstructorCalendar() MyCalendar
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(log n) per booking  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #729: My Calendar I
// https://leetcode.com/problems/my-calendar-i/
// Difficulty: Medium
// Time: O(log n) per booking
// Space: O(n)

import "fmt"

func main() {
	cal := ConstructorCalendar()
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(15, 25))
	fmt.Println(cal.Book(20, 30))
}

type MyCalendar struct {
	books [][2]int
}

func ConstructorCalendar() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	for _, b := range c.books {
		if max(b[0], start) < min(b[1], end) {
			return false
		}
	}
	c.books = append(c.books, [2]int{start, end})
	return true
}
```
