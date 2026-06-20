# 3878 — Count Good Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countGoodSubarrays(nums []int) int64
```

> **💡 Hint:** Sliding window with frequency map. For each right index,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3878: Count Good Subarrays
// https://leetcode.com/problems/count-good-subarrays/
// Difficulty: Hard
//
// Count subarrays where all elements are distinct (no duplicates).
//
// Approach: Sliding window with frequency map. For each right index,
// maintain window with all distinct elements. Count subarrays ending
// at right with all distinct elements.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countGoodSubarrays([]int{1, 2, 3}))
	// Example 2
	fmt.Println(countGoodSubarrays([]int{1, 2, 1, 3}))
	// Edge: all same
	fmt.Println(countGoodSubarrays([]int{1, 1, 1}))
	// Edge: empty
	fmt.Println(countGoodSubarrays([]int{}))
}

func countGoodSubarrays(nums []int) int64 {
	n := len(nums)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 0
	}

  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	var ans int64
	left := 0

	for right := 0; right < n; right++ {
		freq[nums[right]]++

		for freq[nums[right]] > 1 {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				delete(freq, nums[left])
			}
			left++
		}

		ans += int64(right - left + 1)
	}

	return ans
}
```
