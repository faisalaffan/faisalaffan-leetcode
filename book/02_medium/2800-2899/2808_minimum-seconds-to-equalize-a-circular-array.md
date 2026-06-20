# 2808 — Minimum Seconds To Equalize A Circular Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumSecondsToEqualizeACircularArray(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2808: Minimum Seconds to Equalize a Circular Array
// https://leetcode.com/problems/minimum-seconds-to-equalize-a-circular-array/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MinimumSecondsToEqualizeACircularArray(nums []int) int {
	n := len(nums)
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

	best := n / 2
	for _, positions := range pos {
		if len(positions) == 0 {
			continue
		}
		maxGap := 0
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(positions); i++ {
			curr := positions[i]
			var prev int
			if i > 0 {
				prev = positions[i-1]
			} else {
				prev = positions[len(positions)-1] - n
			}
			gap := curr - prev
			if gap > maxGap {
				maxGap = gap
			}
		}
		seconds := maxGap / 2
		if seconds < best {
			best = seconds
		}
	}

	return best
}

func main() {
	fmt.Println(MinimumSecondsToEqualizeACircularArray([]int{1, 2, 1, 2}))
	fmt.Println(MinimumSecondsToEqualizeACircularArray([]int{2, 1, 3, 3, 2}))
}
```
