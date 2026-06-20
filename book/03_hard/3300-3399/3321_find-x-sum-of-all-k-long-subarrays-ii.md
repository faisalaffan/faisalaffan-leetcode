# 3321 — Find X Sum Of All K Long Subarrays Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func findXSum(nums []int, k int, x int) []int64
```

> **💡 Hint:** Sliding window with two ordered sets (balanced BST)

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3321: Find X-Sum of All K-Long Subarrays II
// https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-ii/
// Difficulty: Hard
//
// For each k-length subarray, compute the x-sum: sum of top x most
// frequent elements (by frequency, then by value). Return an array
// of x-sums for each subarray.
//
// Approach: Sliding window with two ordered sets (balanced BST)
// implemented via sorted slices. Maintain top x elements and the
// remaining elements.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(findXSum([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2))
	// Example 2
	fmt.Println(findXSum([]int{3, 3, 3, 3}, 3, 1))
	// Edge: k = 1
	fmt.Println(findXSum([]int{5, 5, 5}, 1, 1))
}

type pair struct {
	val int
	cnt int
}

func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums)
  // Alokasi slice integer
	ans := make([]int64, n-k+1)
  // Membuat map (HashMap) — pencarian O(1)
	freq := make(map[int]int)

	for i := 0; i < n; i++ {
		freq[nums[i]]++
		if i >= k {
			freq[nums[i-k]]--
			if freq[nums[i-k]] == 0 {
				delete(freq, nums[i-k])
			}
		}
		if i >= k-1 {
			// Build list and compute x-sum
			var list []pair
			for val, cnt := range freq {
				list = append(list, pair{val, cnt})
			}
  // Custom sort dengan comparator
			sort.Slice(list, func(i, j int) bool {
				if list[i].cnt != list[j].cnt {
					return list[i].cnt > list[j].cnt
				}
				return list[i].val > list[j].val
			})
			var sum int64
			for j := 0; j < x && j < len(list); j++ {
				sum += int64(list[j].val) * int64(list[j].cnt)
			}
			ans[i-k+1] = sum
		}
	}
	return ans
}
```
