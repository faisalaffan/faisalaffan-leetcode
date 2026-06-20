# 3488 — Closest Equal Element Queries

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func ClosestEqualElementQueries(nums []int, queries []int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3488: Closest Equal Element Queries
// https://leetcode.com/problems/closest-equal-element-queries/
// Difficulty: Medium
// Complexity: O(n + q) time, O(n) space

import "fmt"

func main() {
	// Test case 1
	nums := []int{1, 2, 3, 2, 1}
	queries := []int{0, 3, 4}
	fmt.Println("Test 1:", ClosestEqualElementQueries(nums, queries))

	// Test case 2
	nums2 := []int{1, 1, 1}
	queries2 := []int{0, 1}
	fmt.Println("Test 2:", ClosestEqualElementQueries(nums2, queries2))

	// Test case 3
	nums3 := []int{1, 2, 3}
	queries3 := []int{0}
	fmt.Println("Test 3:", ClosestEqualElementQueries(nums3, queries3))
}

func ClosestEqualElementQueries(nums []int, queries []int) []int {
  // Membuat map (HashMap) — pencarian O(1)
	pos := make(map[int][]int)
	for i, v := range nums {
		pos[v] = append(pos[v], i)
	}

  // Alokasi slice integer
	result := make([]int, len(queries))
	for idx, q := range queries {
		positions := pos[nums[q]]
		if len(positions) <= 1 {
			result[idx] = -1
			continue
		}
		// binary search for q in positions
		left, right := 0, len(positions)-1
		best := -1
		for left <= right {
			mid := left + (right-left)/2
			if positions[mid] == q {
				// check neighbors
				if mid > 0 {
					dist := q - positions[mid-1]
					if best == -1 || dist < best {
						best = dist
					}
				}
				if mid < len(positions)-1 {
					dist := positions[mid+1] - q
					if best == -1 || dist < best {
						best = dist
					}
				}
				break
			} else if positions[mid] < q {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		result[idx] = best
	}
	return result
}
```
