# 0983 — Minimum Cost For Tickets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func mincostTickets(days []int, costs []int) int
```

> **💡 Hint:** DP (bottom-up) over travel days

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n) where n is the range of days (last travel day)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #983: Minimum Cost For Tickets
// https://leetcode.com/problems/minimum-cost-for-tickets/
// Difficulty: Medium
//
// Approach: DP (bottom-up) over travel days
// Time: O(n) where n is the range of days (last travel day)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(mincostTickets([]int{1, 4, 6, 7, 8, 20}, []int{2, 7, 15})) // 11
	fmt.Println(mincostTickets([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31}, []int{2, 7, 15})) // 17
	fmt.Println(mincostTickets([]int{1, 2, 3}, []int{2, 7, 15})) // 6
}

func mincostTickets(days []int, costs []int) int {
	lastDay := days[len(days)-1]
  // Alokasi slice integer
	dp := make([]int, lastDay+1)
  // Membuat map (HashMap) — pencarian O(1)
	travelSet := make(map[int]bool)
	for _, d := range days {
		travelSet[d] = true
	}

	for i := 1; i <= lastDay; i++ {
		if !travelSet[i] {
			dp[i] = dp[i-1]
			continue
		}
		one := dp[i-1] + costs[0]
		seven := dp[max(0, i-7)] + costs[1]
		thirty := dp[max(0, i-30)] + costs[2]
		dp[i] = min(one, min(seven, thirty))
	}

	return dp[lastDay]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
