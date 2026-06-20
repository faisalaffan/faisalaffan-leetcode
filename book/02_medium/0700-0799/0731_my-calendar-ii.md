# 0731 — My Calendar Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ConstructorCalendar2() MyCalendarTwo
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2) per booking  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #731: My Calendar II
// https://leetcode.com/problems/my-calendar-ii/
// Difficulty: Medium
// Time: O(n^2) per booking
// Space: O(n)

import "fmt"

func main() {
	cal := ConstructorCalendar2()
	fmt.Println(cal.Book(10, 20))
	fmt.Println(cal.Book(50, 60))
	fmt.Println(cal.Book(10, 40))
	fmt.Println(cal.Book(5, 15))
	fmt.Println(cal.Book(5, 10))
	fmt.Println(cal.Book(25, 55))
}

type MyCalendarTwo struct {
	books    [][2]int
	overlaps [][2]int
}

func ConstructorCalendar2() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (c *MyCalendarTwo) Book(start int, end int) bool {
	for _, o := range c.overlaps {
		if max(o[0], start) < min(o[1], end) {
			return false
		}
	}
	for _, b := range c.books {
		if max(b[0], start) < min(b[1], end) {
			c.overlaps = append(c.overlaps, [2]int{max(b[0], start), min(b[1], end)})
		}
	}
	c.books = append(c.books, [2]int{start, end})
	return true
}
```
