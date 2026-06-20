# 3741 — Minimum Distance Between Three Equal Elements Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumDistanceBetweenThreeEqualElementsIi(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3741: Minimum Distance Between Three Equal Elements II
// https://leetcode.com/problems/minimum-distance-between-three-equal-elements-ii/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumDistanceBetweenThreeEqualElementsIi(nums []int) int {
	ans := -1
  // Membuat map (HashMap) — pencarian O(1)
	prev1 := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	prev2 := make(map[int]int)

	// prev1[v] = last index where v appeared
	// prev2[v] = second-last index where v appeared

	for i, v := range nums {
		if p2, ok := prev2[v]; ok {
			dist := 2 * (i - p2)
			if ans == -1 || dist < ans {
				ans = dist
			}
		}
		// Shift: prev2 gets prev1, prev1 gets current
		prev2[v] = prev1[v]
		prev1[v] = i
	}

	return ans
}

func main() {
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 3, 1, 1, 2, 1}))
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 2, 3, 4}))
	fmt.Println(minimumDistanceBetweenThreeEqualElementsIi([]int{1, 1, 1}))
}
```
