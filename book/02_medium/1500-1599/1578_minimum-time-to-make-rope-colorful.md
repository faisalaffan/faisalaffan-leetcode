# 1578 — Minimum Time To Make Rope Colorful

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinCost(colors string, neededTime []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1578: Minimum Time to Make Rope Colorful
// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinCost("abaac", []int{1, 2, 3, 4, 5}))
	fmt.Println(MinCost("abc", []int{1, 2, 3}))
	fmt.Println(MinCost("aabaa", []int{1, 2, 3, 4, 1}))
}

func MinCost(colors string, neededTime []int) int {
	// Time: O(N), Space: O(1)
	n := len(colors)
	totalTime := 0

	i := 0
	for i < n {
		j := i
		maxTime := 0
		sum := 0
		for j < n && colors[j] == colors[i] {
			maxTime = maxInt(maxTime, neededTime[j])
			sum += neededTime[j]
			j++
		}
		// Keep the max time balloon, remove the rest
		totalTime += sum - maxTime
		i = j
	}

	return totalTime
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
