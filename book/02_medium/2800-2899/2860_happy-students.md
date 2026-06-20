# 2860 — Happy Students

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func HappyStudents(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2860: Happy Students
// https://leetcode.com/problems/happy-students/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func HappyStudents(nums []int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)
	count := 0

	// Check for empty selection (all students are happy)
	if nums[0] != 0 {
		count++
	}

	for i := 0; i < n; i++ {
		selected := i + 1
		// nums[i] < selected: the student at position i needs fewer than selected
		if nums[i] < selected {
			// Check if next student (if exists) needs more than selected
			if i+1 >= n || nums[i+1] > selected {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(HappyStudents([]int{1, 1}))
	fmt.Println(HappyStudents([]int{6, 0, 3, 3, 6, 7, 2, 7}))
}
```
