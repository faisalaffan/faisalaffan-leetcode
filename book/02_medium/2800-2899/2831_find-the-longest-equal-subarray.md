# 2831 — Find The Longest Equal Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindTheLongestEqualSubarray(nums []int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2831: Find the Longest Equal Subarray
// https://leetcode.com/problems/find-the-longest-equal-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func FindTheLongestEqualSubarray(nums []int, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	best := 0
	for _, indices := range pos {
		left := 0
		for right := 0; right < len(indices); right++ {
			// Elements between indices[left] and indices[right] that need to be removed
			for indices[right]-indices[left]-(right-left) > k {
				left++
			}
			if right-left+1 > best {
				best = right - left + 1
			}
		}
	}

	return best
}

func main() {
	fmt.Println(FindTheLongestEqualSubarray([]int{1, 3, 2, 3, 1, 3}, 3))
	fmt.Println(FindTheLongestEqualSubarray([]int{1, 1, 2, 2, 1, 1}, 2))
}
```
