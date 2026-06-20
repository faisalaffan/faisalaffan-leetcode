# 0956 — Tallest Billboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func tallestBillboard(rods []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #956: Tallest Billboard
// https://leetcode.com/problems/tallest-billboard/
// Difficulty: Hard

import "fmt"

func tallestBillboard(rods []int) int {
	// dp[diff] = max total sum of both sides with this abs difference
	dp := map[int]int{0: 0}

	for _, r := range rods {
  // Membuat map (HashMap) — pencarian O(1)
		cur := make(map[int]int)
		for diff, total := range dp {
			// 1. skip this rod — use >= to propagate diff=0/total=0 (map zero-value)
			if total >= cur[diff] {
				cur[diff] = total
			}
			// 2. add to taller side: diff increases by r, total increases by r
			left := total + r
			if left > cur[diff+r] {
				cur[diff+r] = left
			}
			// 3. add to shorter side: abs diff changes
			right := total + r
			newDiff := diff - r
			if newDiff < 0 {
				newDiff = -newDiff
			}
			if right > cur[newDiff] {
				cur[newDiff] = right
			}
		}
		dp = cur
	}

	// dp[0] = max total sum when diff=0 (equal sides)
	// Each side height = dp[0] / 2
	return dp[0] / 2
}

func main() {
	fmt.Println("Example 1:")
	fmt.Println(tallestBillboard([]int{1, 2, 3, 6}))
	// Expected: 6

	fmt.Println("Example 2:")
	fmt.Println(tallestBillboard([]int{1, 2, 3, 4, 5, 6}))
	// Expected: 10

	fmt.Println("Example 3:")
	fmt.Println(tallestBillboard([]int{1, 2}))
	// Expected: 0
}
```
