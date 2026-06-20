# 2863 — Maximum Length Of Semi Decreasing Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumLengthOfSemiDecreasingSubarrays(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2863: Maximum Length of Semi-Decreasing Subarrays
// https://leetcode.com/problems/maximum-length-of-semi-decreasing-subarrays/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func MaximumLengthOfSemiDecreasingSubarrays(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// Store indices where nums[i] > nums[i+1] (start of decreasing)
  // Alokasi slice integer
	starts := make([]int, 0)
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			starts = append(starts, i)
		}
	}

	if len(starts) == 0 {
		return 1
	}

	best := 1
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(starts); i++ {
		end := starts[i] + 1
		// If not the last start, limit by next start
		limit := n - 1
		if i+1 < len(starts) {
			limit = starts[i+1]
		}
		// Extend to the right while keeping non-increasing property
		for end < limit && nums[end] >= nums[end+1] {
			end++
		}
		length := end - starts[i] + 1
		if length > best {
			best = length
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumLengthOfSemiDecreasingSubarrays([]int{1, 2, 3, 4}))
	fmt.Println(MaximumLengthOfSemiDecreasingSubarrays([]int{7, 6, 5, 4, 3, 2, 1, 6, 10, 11}))
}
```
