# 3347 — Maximum Frequency Of An Element After Performing Operations Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxFrequency(nums []int, k int, numOperations int) int
```

> **💡 Hint:** For each unique value, consider it as the final value.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer, Sliding Window, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3347: Maximum Frequency of an Element After Performing Operations II
// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-ii/
// Difficulty: Hard
//
// In one operation, add k to any element in nums. Perform at most
// numOperations operations. Maximize the frequency of any single
// value in the resulting array.
//
// Approach: For each unique value, consider it as the final value.
// Count how many existing elements can reach it within allowed
// operations. Use prefix sums over sorted values.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maxFrequency([]int{1, 2, 4}, 2, 2))
	// Example 2
	fmt.Println(maxFrequency([]int{5, 5, 5, 10}, 5, 1))
	// Edge: single element
	fmt.Println(maxFrequency([]int{7}, 3, 0))
}

func maxFrequency(nums []int, k int, numOperations int) int {
  // Urutkan secara ascending — O(n log n)
	sort.Ints(nums)
	n := len(nums)

	// Count frequency of each value
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}

	// Collect unique values
  // Alokasi slice integer
	unique := make([]int, 0, len(freq))
	for v := range freq {
		unique = append(unique, v)
	}
  // Urutkan secara ascending — O(n log n)
	sort.Ints(unique)

	// Sliding window: count elements in range [val - k, val + k]
	// But only numOperations of them can be changed to val
	ans := 0
	left := 0
	for right := 0; right < n; right++ {
		// Shrink window to [target - k, target + k]
		for nums[right]-nums[left] > 2*k {
			left++
		}
		total := right - left + 1
		originalCnt := freq[nums[right]]
		// We can change at most numOperations elements to this value
		possible := originalCnt + min(numOperations, total-originalCnt)
		if possible > ans {
			ans = possible
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
