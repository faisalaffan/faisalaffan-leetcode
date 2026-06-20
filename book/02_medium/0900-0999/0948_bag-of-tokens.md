# 0948 — Bag Of Tokens

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func bagOfTokensScore(tokens []int, power int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #948: Bag of Tokens
// https://leetcode.com/problems/bag-of-tokens/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1)
func bagOfTokensScore(tokens []int, power int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(tokens)
	left, right := 0, len(tokens)-1
	score, maxScore := 0, 0

	for left <= right {
		if power >= tokens[left] {
			power -= tokens[left]
			left++
			score++
			if score > maxScore {
				maxScore = score
			}
		} else if score > 0 {
			power += tokens[right]
			right--
			score--
		} else {
			break
		}
	}
	return maxScore
}

func main() {
	fmt.Println(bagOfTokensScore([]int{100}, 50))
	fmt.Println(bagOfTokensScore([]int{200, 100}, 150))
	fmt.Println(bagOfTokensScore([]int{100, 200, 300, 400}, 200))
}
```
