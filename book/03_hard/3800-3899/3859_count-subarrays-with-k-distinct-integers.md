# 3859 — Count Subarrays With K Distinct Integers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countSubarrays(nums []int, k int, m int) int64
```

> **💡 Hint:** Sliding window with frequency map. Track distinct count

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3859: Count Subarrays With K Distinct Integers
// https://leetcode.com/problems/count-subarrays-with-k-distinct-integers/
// Difficulty: Hard
//
// Count subarrays that contain exactly k distinct integers, where
// the maximum element in the subarray appears at least m times.
//
// Approach: Sliding window with frequency map. Track distinct count
// and max frequency. Expand right, shrink left when conditions
// violated.

import "fmt"

func main() {
	// Example 1
	fmt.Println(countSubarrays([]int{1, 2, 1, 2, 3}, 2, 1))
	// Example 2
	fmt.Println(countSubarrays([]int{1, 1, 2, 2, 3}, 2, 2))
	// Edge: single element
	fmt.Println(countSubarrays([]int{5}, 1, 1))
	// Edge: k = 1, m = 2
	fmt.Println(countSubarrays([]int{1, 1, 1, 2}, 1, 2))
}

func countSubarrays(nums []int, k int, m int) int64 {
	n := len(nums)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	var ans int64
	left := 0
	distinct := 0
	maxFreq := 0

	for right := 0; right < n; right++ {
		val := nums[right]
		freq[val]++
		if freq[val] == 1 {
			distinct++
		}
		if freq[val] > maxFreq {
			maxFreq = freq[val]
		}

		for distinct > k || (distinct == k && maxFreq < m) {
			freq[nums[left]]--
			if freq[nums[left]] == 0 {
				distinct--
			}
			left++
			maxFreq = 0
			for _, c := range freq {
				if c > maxFreq {
					maxFreq = c
				}
			}
		}

		if distinct == k && maxFreq >= m {
			ans += int64(right - left + 1)
		}
	}

	return ans
}
```
