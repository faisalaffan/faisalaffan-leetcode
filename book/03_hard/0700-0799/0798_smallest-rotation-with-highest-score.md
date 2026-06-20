# 0798 — Smallest Rotation With Highest Score

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func bestRotation(nums []int) int
```

> **💡 Hint:** difference array

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #798: Smallest Rotation with Highest Score
// https://leetcode.com/problems/smallest-rotation-with-highest-score/
// Difficulty: Hard
//
// For each rotation k (0 <= k < n), we move the first k elements to the end.
// An element at original index i with value A[i] gets a point at rotation k
// if its new position (i - k + n) % n >= A[i].
//
// Approach: difference array
//   - For each element A[i], the rotations where it does NOT score form a
//     contiguous (mod n) interval of size A[i].
//   - Mark these "bad" intervals in a difference array, then compute the
//     prefix sum to find the rotation with the fewest bad elements (== highest
//     score). Return the smallest such rotation.

import "fmt"

func main() {
	// Example: [2,3,1,4,0] → best k=3 (score 4)
	fmt.Println(bestRotation([]int{2, 3, 1, 4, 0}))
	// Already optimal at k=0
	fmt.Println(bestRotation([]int{1, 3, 0, 2, 4}))
	// Single element
	fmt.Println(bestRotation([]int{0}))
	// All zeros: every rotation scores n
	fmt.Println(bestRotation([]int{0, 0, 0}))
	// Descending
	fmt.Println(bestRotation([]int{3, 2, 1, 0}))
}

func bestRotation(nums []int) int {
	n := len(nums)
  // Alokasi slice integer
	diff := make([]int, n+1)

	for i, val := range nums {
		if val == 0 {
			continue // never bad (always scores at any position)
		}
		// Bad interval: k in [(i-val+1) mod n, i] (mod n)
		// This covers exactly `val` values of k.
		l := (i - val + 1 + n) % n
		r := i

		if l <= r {
			diff[l]++
			diff[r+1]--
		} else {
			// Wrapping interval: [l, n-1] and [0, r]
			diff[l]++
			diff[n]--
			diff[0]++
			diff[r+1]--
		}
	}

	minBad := n + 1
	bestK := 0
	bad := 0
	for k := 0; k < n; k++ {
		bad += diff[k]
		if bad < minBad {
			minBad = bad
			bestK = k
		}
	}

	return bestK
}
```
