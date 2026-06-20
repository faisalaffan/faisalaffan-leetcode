# 3761 — Minimum Absolute Distance Between Mirror Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumAbsoluteDistanceBetweenMirrorPairs(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3761: Minimum Absolute Distance Between Mirror Pairs
// https://leetcode.com/problems/minimum-absolute-distance-between-mirror-pairs/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumAbsoluteDistanceBetweenMirrorPairs(nums []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	prev := make(map[int]int)
	ans := -1

	for j, v := range nums {
		if pos, ok := prev[v]; ok {
			dist := j - pos
			if ans == -1 || dist < ans {
				ans = dist
			}
		}
		// Store reversed number
		rev := 0
		for x := v; x > 0; x /= 10 {
			rev = rev*10 + x%10
		}
		prev[rev] = j
	}

	return ans
}

func main() {
	fmt.Println(minimumAbsoluteDistanceBetweenMirrorPairs([]int{12, 21, 45, 33, 54}))
	fmt.Println(minimumAbsoluteDistanceBetweenMirrorPairs([]int{1, 2, 3, 4}))
	fmt.Println(minimumAbsoluteDistanceBetweenMirrorPairs([]int{11, 22, 11}))
}
```
