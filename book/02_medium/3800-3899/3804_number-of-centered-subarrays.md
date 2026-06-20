# 3804 — Number Of Centered Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NumberOfCenteredSubarrays(nums []int) int
```

> **💡 Hint:** For each start index, expand subarrays and track running sum

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3804: Number of Centered Subarrays
// https://leetcode.com/problems/number-of-centered-subarrays/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)
// Approach: For each start index, expand subarrays and track running sum
// with a set of seen elements. A subarray is centered if its sum equals
// at least one element within it.

import "fmt"

func NumberOfCenteredSubarrays(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		sum := 0
  // Membuat map (HashMap) — pencarian O(1)
		seen := make(map[int]bool)
		for j := i; j < n; j++ {
			seen[nums[j]] = true
			sum += nums[j]
			if seen[sum] {
				ans++
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(NumberOfCenteredSubarrays([]int{-1, 1, 0})) // Expected: 5

	// Example 2
	fmt.Println(NumberOfCenteredSubarrays([]int{2, -3})) // Expected: 2

	// Example 3
	fmt.Println(NumberOfCenteredSubarrays([]int{1, 2, 3})) // Expected: 3 (all single elements are centered)
}
```
