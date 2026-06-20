# 2244 — Minimum Rounds To Complete All Tasks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumRounds(tasks []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2244: Minimum Rounds to Complete All Tasks
// https://leetcode.com/problems/minimum-rounds-to-complete-all-tasks/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func minimumRounds(tasks []int) int {
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, t := range tasks {
		freq[t]++
	}

	rounds := 0
	for _, c := range freq {
		if c == 1 {
			return -1
		}
		// 3*c + 2*(c/3 remainder)
		rounds += c / 3
		if c%3 != 0 {
			rounds++
		}
	}
	return rounds
}

func main() {
	// Test case 1
	fmt.Println(minimumRounds([]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4}))
	// Expected: 4

	// Test case 2
	fmt.Println(minimumRounds([]int{2, 3, 3}))
	// Expected: -1

	// Test case 3
	fmt.Println(minimumRounds([]int{5, 5, 5, 5}))
	// Expected: 2
}
```
