# 3729 — Count Distinct Subarrays Divisible By K In Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numGoodSubarrays(nums []int, k int) int64
```

> **💡 Hint:** Prefix sum + hash set of subarray content for dedup.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3729: Count Distinct Subarrays Divisible by K in Sorted Array
// https://leetcode.com/problems/count-distinct-subarrays-divisible-by-k-in-sorted-array/
// Difficulty: Hard
//
// Count distinct subarrays (by value sequence) whose sum is divisible by k.
// Array is sorted, so equal values are contiguous. Subarrays are distinct
// when their sequences of values differ.
//
// Approach: Prefix sum + hash set of subarray content for dedup.

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Example 1
	fmt.Println(numGoodSubarrays([]int{1, 2, 3}, 3))
	// Example 2
	fmt.Println(numGoodSubarrays([]int{1, 1, 2}, 2))
	// Edge: single element
	fmt.Println(numGoodSubarrays([]int{5}, 5))
	// Edge: all same
	fmt.Println(numGoodSubarrays([]int{2, 2, 2}, 2))
}

func numGoodSubarrays(nums []int, k int) int64 {
	n := len(nums)
  // Alokasi slice integer
	pref := make([]int, n+1)
	for i, v := range nums {
		pref[i+1] = pref[i] + v
	}

  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]bool)
	var ans int64

	for i := 0; i < n; i++ {
		var sb strings.Builder
		for j := i; j < n; j++ {
			if j > i {
				sb.WriteByte(',')
			}
			sb.WriteString(strconv.Itoa(nums[j]))
			if (pref[j+1]-pref[i])%k == 0 {
				key := sb.String()
				if !seen[key] {
					seen[key] = true
					ans++
				}
			}
		}
	}

	return ans
}
```
