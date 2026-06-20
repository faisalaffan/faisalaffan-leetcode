# 2347 — Best Poker Hand

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func BestPokerHand(ranks []int, suits []byte) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2347: Best Poker Hand
// https://leetcode.com/problems/best-poker-hand/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(BestPokerHand([]int{13, 2, 3, 1, 9}, []byte{'a', 'a', 'a', 'a', 'a'})) // "Flush"
	fmt.Println(BestPokerHand([]int{4, 4, 2, 4, 4}, []byte{'d', 'a', 'a', 'b', 'c'})) // "Three of a Kind"
	fmt.Println(BestPokerHand([]int{10, 10, 2, 12, 9}, []byte{'a', 'b', 'c', 'a', 'd'})) // "Pair"
}

func BestPokerHand(ranks []int, suits []byte) string {
	// Check flush
	if suits[0] == suits[1] && suits[1] == suits[2] && suits[2] == suits[3] && suits[3] == suits[4] {
		return "Flush"
	}

	// Check three of a kind or pair
	rankCount := [14]int{}
	for _, r := range ranks {
		rankCount[r]++
		if rankCount[r] == 3 {
			return "Three of a Kind"
		}
	}
	for _, c := range rankCount {
		if c == 2 {
			return "Pair"
		}
	}
	return "High Card"
}
```
