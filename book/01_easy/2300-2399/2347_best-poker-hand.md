# 2347 — Best Poker Hand

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func BestPokerHand(ranks []int, suits []byte) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


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
