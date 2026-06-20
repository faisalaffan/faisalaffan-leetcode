# 3728 — Stable Subarrays With Equal Boundary And Interior Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func stableSubarraysWithEqualBoundaryAndInteriorSum(capacity []int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3728: Stable Subarrays With Equal Boundary and Interior Sum
// https://leetcode.com/problems/stable-subarrays-with-equal-boundary-and-interior-sum/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func stableSubarraysWithEqualBoundaryAndInteriorSum(capacity []int) int64 {
	n := len(capacity)
  // Alokasi slice integer
	s := make([]int64, n+1)
	for i := 0; i < n; i++ {
		s[i+1] = s[i] + int64(capacity[i])
	}

	type pair struct {
		val int
		sum int64
	}
  // Membuat map (HashMap) — pencarian O(1)
	cnt := make(map[pair]int)
	var ans int64

	for r := 0; r < n; r++ {
		// Query for valid left boundaries
		key := pair{val: capacity[r], sum: s[r] - int64(capacity[r])}
		ans += int64(cnt[key])

		// Insert position r-1 as future left boundary (delayed by 1 for length >= 3)
		if r >= 1 {
			ins := pair{val: capacity[r-1], sum: s[r]}
			cnt[ins]++
		}
	}

	return ans
}

func main() {
	fmt.Println(stableSubarraysWithEqualBoundaryAndInteriorSum([]int{9, 3, 3, 3, 9}))
	fmt.Println(stableSubarraysWithEqualBoundaryAndInteriorSum([]int{1, 2, 3, 4, 5}))
	fmt.Println(stableSubarraysWithEqualBoundaryAndInteriorSum([]int{5, 2, 1, 2, 5}))
}
```
