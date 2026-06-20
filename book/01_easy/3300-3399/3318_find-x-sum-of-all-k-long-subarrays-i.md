# 3318 — Find X Sum Of All K Long Subarrays I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindXSumOfAllKLongSubarraysI(nums []int, k int, x int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n * k log k). Space: O(k).  
**Kompleksitas Ruang:** O(k).

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3318: Find X-Sum of All K-Long Subarrays I
// https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-i/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindXSumOfAllKLongSubarraysI([]int{1, 1, 2, 2, 3, 4, 2, 3}, 6, 2))
	fmt.Println(FindXSumOfAllKLongSubarraysI([]int{3, 8, 7, 8, 7, 5}, 2, 2))
}

// FindXSumOfAllKLongSubarraysI calculates x-sum for each k-length subarray.
// The x-sum is the sum of the top x most frequent elements (breaking ties by larger value).
// Time: O(n * k log k). Space: O(k).
func FindXSumOfAllKLongSubarraysI(nums []int, k int, x int) []int {
	n := len(nums)
  // Alokasi slice integer
	result := make([]int, n-k+1)
	for start := 0; start <= n-k; start++ {
  // Membuat map (HashMap) — pencarian O(1)
		freq := make(map[int]int)
		for i := start; i < start+k; i++ {
			freq[nums[i]]++
		}

		type pair struct {
			val int
			cnt int
		}
		pairs := make([]pair, 0, len(freq))
		for val, cnt := range freq {
			pairs = append(pairs, pair{val, cnt})
		}
  // Custom sort dengan comparator
		sort.Slice(pairs, func(i, j int) bool {
			if pairs[i].cnt != pairs[j].cnt {
				return pairs[i].cnt > pairs[j].cnt
			}
			return pairs[i].val > pairs[j].val
		})

		sum := 0
		for i := 0; i < x && i < len(pairs); i++ {
			sum += pairs[i].val * pairs[i].cnt
		}
		result[start] = sum
	}
	return result
}
```
