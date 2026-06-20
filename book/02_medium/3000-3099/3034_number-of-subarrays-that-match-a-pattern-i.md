# 3034 — Number Of Subarrays That Match A Pattern I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countMatchingSubarrays(nums []int, pattern []int) (ans int)
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n*m)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3034: Number of Subarrays That Match a Pattern I
// https://leetcode.com/problems/number-of-subarrays-that-match-a-pattern-i/
// Difficulty: Medium
// Time: O(n*m) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(countMatchingSubarrays([]int{1, 2, 3, 4, 5, 6}, []int{1, 1}))
	fmt.Println(countMatchingSubarrays([]int{1, 4, 4, 1, 3, 5, 5, 3}, []int{1, 0, -1}))
}

func countMatchingSubarrays(nums []int, pattern []int) (ans int) {
	n, m := len(nums), len(pattern)
outer:
	for i := 0; i+m < n; i++ {
		for k := 0; k < m; k++ {
			diff := 0
			if nums[i+k+1] > nums[i+k] {
				diff = 1
			} else if nums[i+k+1] < nums[i+k] {
				diff = -1
			}
			if diff != pattern[k] {
				continue outer
			}
		}
		ans++
	}
	return
}
```
