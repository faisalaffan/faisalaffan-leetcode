# 2459 — Sort Array By Moving Items To Empty Space

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func sortArray(nums []int) int
```

> **💡 Hint:** Cycle decomposition. For each cycle, if it contains 0,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2459: Sort Array by Moving Items to Empty Space
// https://leetcode.com/problems/sort-array-by-moving-items-to-empty-space/
// Difficulty: Hard [Paid]
//
// Given a permutation of 0..n-1 where 0 represents empty space,
// find the minimum number of moves to sort the array.
// A move consists of moving any element to the empty space position.
//
// Approach: Cycle decomposition. For each cycle, if it contains 0,
// cycleLen-1 moves are needed. Otherwise, cycleLen+1 moves needed
// (to bring 0 in and back out).

import "fmt"

func main() {
	// Example 1
	fmt.Println(sortArray([]int{4, 2, 0, 3, 1}))
	// Example 2
	fmt.Println(sortArray([]int{1, 0, 2, 3}))
	// Example 3
	fmt.Println(sortArray([]int{0, 1, 2, 3}))
	// Edge: already sorted
	fmt.Println(sortArray([]int{0, 1, 2, 3, 4}))
}

func sortArray(nums []int) int {
	n := len(nums)
	visited := make([]bool, n)
	ans := 0
	zeroPos := 0
	for i, v := range nums {
		if v == 0 {
			zeroPos = i
			break
		}
	}

	for i := 0; i < n; i++ {
		if visited[i] || nums[i] == i {
			visited[i] = true
			continue
		}
		// Find cycle
		cycleLen := 0
		hasZero := false
		j := i
		for !visited[j] {
			visited[j] = true
			cycleLen++
			if nums[j] == 0 {
				hasZero = true
			}
			j = nums[j]
		}
		if cycleLen > 0 {
			if hasZero {
				ans += cycleLen - 1
			} else {
				ans += cycleLen + 1
			}
		}
	}
	_ = zeroPos
	return ans
}
```
