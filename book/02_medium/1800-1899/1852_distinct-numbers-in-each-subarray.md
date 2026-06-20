# 1852 — Distinct Numbers In Each Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DistinctNumbers(nums []int, k int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Sliding Window

**Kompleksitas Waktu:** O(n), Space: O(k) where k = distinct elements in window  
**Kompleksitas Ruang:** O(k) where k = distinct elements in window

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1852: Distinct Numbers in Each Subarray
// https://leetcode.com/problems/distinct-numbers-in-each-subarray/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(DistinctNumbers([]int{1, 2, 3, 2, 2, 1, 3}, 3))
	fmt.Println(DistinctNumbers([]int{1, 1, 1, 1, 1}, 2))
	fmt.Println(DistinctNumbers([]int{1, 2, 3, 4}, 1))
}

// Time: O(n), Space: O(k) where k = distinct elements in window
func DistinctNumbers(nums []int, k int) []int {
	n := len(nums)
	if k > n {
		return nil
	}
  // Alokasi slice integer
	result := make([]int, n-k+1)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	distinct := 0

	for i := 0; i < n; i++ {
		// Add right element
		freq[nums[i]]++
		if freq[nums[i]] == 1 {
			distinct++
		}

		// Remove left element when window exceeds k
		if i >= k {
			freq[nums[i-k]]--
			if freq[nums[i-k]] == 0 {
				distinct--
			}
		}

		// Record result for completed windows
		if i >= k-1 {
			result[i-k+1] = distinct
		}
	}
	return result
}
```
