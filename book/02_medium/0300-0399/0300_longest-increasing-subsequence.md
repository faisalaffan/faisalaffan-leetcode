# 0300 — Longest Increasing Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func lengthOfLIS(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #300: Longest Increasing Subsequence
// https://leetcode.com/problems/longest-increasing-subsequence/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import "fmt"

func lengthOfLIS(nums []int) int {
	tails := []int{}

	for _, num := range nums {
		left, right := 0, len(tails)
  // Two-pointer: gerakkan kiri atau kanan
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] < num {
				left = mid + 1
			} else {
				right = mid
			}
		}

		if left == len(tails) {
			tails = append(tails, num)
		} else {
			tails[left] = num
		}
	}

	return len(tails)
}

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
}
```
