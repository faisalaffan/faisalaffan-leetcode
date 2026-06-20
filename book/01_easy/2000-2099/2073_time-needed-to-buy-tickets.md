# 2073 — Time Needed To Buy Tickets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func TimeNeededToBuyTickets(tickets []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2073: Time Needed to Buy Tickets
// https://leetcode.com/problems/time-needed-to-buy-tickets/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(TimeNeededToBuyTickets([]int{2, 3, 2}, 2)) // 6
	fmt.Println(TimeNeededToBuyTickets([]int{5, 1, 1, 1}, 0)) // 8
}

// Time: O(n), Space: O(1)
func TimeNeededToBuyTickets(tickets []int, k int) int {
	time := 0
	for i, t := range tickets {
		if i <= k {
			if t <= tickets[k] {
				time += t
			} else {
				time += tickets[k]
			}
		} else {
			if t < tickets[k] {
				time += t
			} else {
				time += tickets[k] - 1
			}
		}
	}
	return time
}
```
