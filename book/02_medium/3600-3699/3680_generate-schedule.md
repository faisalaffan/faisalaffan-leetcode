# 3680 — Generate Schedule

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func generateSchedule(n int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3680: Generate Schedule
// https://leetcode.com/problems/generate-schedule/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n^2)

import "fmt"

func generateSchedule(n int) [][]int {
	if n < 5 {
		return [][]int{}
	}

	totalMatches := n * (n - 1)
  // Membuat matriks/slice 2D untuk DP
	schedule := make([][]int, 0, totalMatches)

	// Phase 1: offset 2 to n-2
	for offset := 2; offset < n-1; offset++ {
		for team := 0; team < n; team++ {
			schedule = append(schedule, []int{team, (team + offset) % n})
		}
	}

	// Phase 2: wrap-around pairs
	for team := 0; team < n; team++ {
		schedule = append(schedule, []int{team, (team + 1) % n})
		schedule = append(schedule, []int{(team + 4) % n, (team + 3) % n})
	}

	return schedule
}

func main() {
	fmt.Println(len(generateSchedule(5)))
	fmt.Println(len(generateSchedule(3)))
	fmt.Println(len(generateSchedule(6)))
}
```
