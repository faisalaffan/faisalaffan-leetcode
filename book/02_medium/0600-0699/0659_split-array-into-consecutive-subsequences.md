# 0659 — Split Array Into Consecutive Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isPossible(nums []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #659: Split Array into Consecutive Subsequences
// https://leetcode.com/problems/split-array-into-consecutive-subsequences/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 3, 4, 4, 5, 5}))
	fmt.Println(isPossible([]int{1, 2, 3, 4, 4, 5}))
}

func isPossible(nums []int) bool {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
  // Membuat map (HashMap) — pencarian O(1)
	tail := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	for _, num := range nums {
		if freq[num] == 0 {
			continue
		}

		if tail[num] > 0 {
			tail[num]--
			freq[num]--
			tail[num+1]++
		} else if freq[num+1] > 0 && freq[num+2] > 0 {
			freq[num]--
			freq[num+1]--
			freq[num+2]--
			tail[num+3]++
		} else {
			return false
		}
	}

	return true
}
```
