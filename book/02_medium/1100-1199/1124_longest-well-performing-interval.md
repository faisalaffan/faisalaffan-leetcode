# 1124 — Longest Well Performing Interval

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func longestWPI(hours []int) int
```

> **💡 Hint:** Prefix sum. Map first occurrence of each prefix sum.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Prefix Sum

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1124: Longest Well-Performing Interval
// https://leetcode.com/problems/longest-well-performing-interval/
// Difficulty: Medium
//
// Approach: Prefix sum. Map first occurrence of each prefix sum.
//           hours[i] > 8 -> +1, else -1. Find longest subarray with sum > 0.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(longestWPI([]int{9, 9, 6, 0, 6, 6, 9})) // 3
	fmt.Println(longestWPI([]int{6, 6, 6}))              // 0
}

func longestWPI(hours []int) int {
	prefix := 0
  // Membuat map (HashMap) — pencarian O(1)
	firstSeen := make(map[int]int)
	result := 0

	for i, h := range hours {
		if h > 8 {
			prefix++
		} else {
			prefix--
		}

		if prefix > 0 {
			result = i + 1
		} else {
			if _, ok := firstSeen[prefix]; !ok {
				firstSeen[prefix] = i
			}
			if j, ok := firstSeen[prefix-1]; ok {
				if i-j > result {
					result = i - j
				}
			}
		}
	}

	return result
}
```
