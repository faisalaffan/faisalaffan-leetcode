# 1231 — Divide Chocolate

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximizeSweetness(sweetness []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1231: Divide Chocolate
// https://leetcode.com/problems/divide-chocolate/
// Difficulty: Hard [Paid]

import "fmt"

func main() {
	fmt.Println("1231. Divide Chocolate")
	fmt.Println("[1,2,3,4,5,6,7,8,9], k=4:", maximizeSweetness([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 4), "(expected 7)")
	fmt.Println("[5,6,7,8,9,1,2,3,4], k=4:", maximizeSweetness([]int{5, 6, 7, 8, 9, 1, 2, 3, 4}, 4), "(expected 7)")
	fmt.Println("[1,2,2,1,2,2,1,2,2], k=2:", maximizeSweetness([]int{1, 2, 2, 1, 2, 2, 1, 2, 2}, 2), "(expected 5)")
}

func maximizeSweetness(sweetness []int, k int) int {
	// We split into k+1 pieces (k friends + ourselves).
	total := 0
	for _, s := range sweetness {
		total += s
	}

	left, right := 1, total/(k+1)
  // Two-pointer: gerakkan kiri atau kanan
	for left < right {
		mid := (left + right + 1) / 2
		if canSplit(sweetness, k+1, mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

// canSplit checks if sweetness can be split into at least 'pieces' contiguous
// subarrays each with sum >= target.
func canSplit(sweetness []int, pieces int, target int) bool {
	count := 0
	sum := 0
	for _, s := range sweetness {
		sum += s
		if sum >= target {
			count++
			sum = 0
		}
	}
	return count >= pieces
}
```
