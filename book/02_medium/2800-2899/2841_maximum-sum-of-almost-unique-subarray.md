# 2841 — Maximum Sum Of Almost Unique Subarray

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumSumOfAlmostUniqueSubarray(nums []int, m int, k int) int64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2841: Maximum Sum of Almost Unique Subarray
// https://leetcode.com/problems/maximum-sum-of-almost-unique-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func MaximumSumOfAlmostUniqueSubarray(nums []int, m int, k int) int64 {
	n := len(nums)
	if n < k {
		return 0
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	var sum int64
	var best int64

	for i := 0; i < k; i++ {
		freq[nums[i]]++
		sum += int64(nums[i])
	}

	if len(freq) >= m {
		best = sum
	}

	for i := k; i < n; i++ {
		// Remove leftmost
		left := nums[i-k]
		freq[left]--
		if freq[left] == 0 {
			delete(freq, left)
		}
		sum -= int64(left)

		// Add rightmost
		right := nums[i]
		freq[right]++
		sum += int64(right)

		if len(freq) >= m && sum > best {
			best = sum
		}
	}

	return best
}

func main() {
	fmt.Println(MaximumSumOfAlmostUniqueSubarray([]int{2, 6, 7, 3, 1, 7}, 3, 4))
	fmt.Println(MaximumSumOfAlmostUniqueSubarray([]int{1, 1, 1, 3}, 2, 2))
}
```
