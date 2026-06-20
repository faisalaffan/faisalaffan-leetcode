# 1105 — Filling Bookcase Shelves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minHeightShelves(books [][]int, shelfWidth int) int
```

> **💡 Hint:** DP. dp[i] = min height to place first i books.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Bitmask

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1105: Filling Bookcase Shelves
// https://leetcode.com/problems/filling-bookcase-shelves/
// Difficulty: Medium
//
// Approach: DP. dp[i] = min height to place first i books.
//           Try placing books i-1..j on the same shelf.
// Time: O(n^2)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4)) // 6
	fmt.Println(minHeightShelves([][]int{{1, 3}, {2, 4}, {3, 2}}, 6))                               // 4
}

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
  // Alokasi slice integer
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = 1<<31 - 1
	}

	for i := 1; i <= n; i++ {
		width := 0
		height := 0
		for j := i; j > 0; j-- {
			width += books[j-1][0]
			if width > shelfWidth {
				break
			}
			if books[j-1][1] > height {
				height = books[j-1][1]
			}
			if dp[j-1]+height < dp[i] {
				dp[i] = dp[j-1] + height
			}
		}
	}

	return dp[n]
}
```
