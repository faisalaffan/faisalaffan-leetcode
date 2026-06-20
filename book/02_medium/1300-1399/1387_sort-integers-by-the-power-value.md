# 1387 — Sort Integers By The Power Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func getKth(lo int, hi int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Dynamic Programming

**Kompleksitas Waktu:** O(n log n) for sorting  
**Kompleksitas Ruang:** O(n) for memoization and sorted array

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1387: Sort Integers by The Power Value
// https://leetcode.com/problems/sort-integers-by-the-power-value/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(getKth(12, 15, 2)) // 13

	// Test case 2
	fmt.Println(getKth(1, 1, 1)) // 1

	// Test case 3
	fmt.Println(getKth(7, 11, 4)) // 7

	// Test case 4
	fmt.Println(getKth(10, 20, 5)) // 13
}

// Time: O(n log n) for sorting
// Space: O(n) for memoization and sorted array
func getKth(lo int, hi int, k int) int {
  // Membuat map (HashMap) — pencarian O(1)
	memo := make(map[int]int)
	memo[1] = 0

	var power func(int) int
	power = func(x int) int {
		if val, ok := memo[x]; ok {
			return val
		}
		if x%2 == 0 {
			memo[x] = 1 + power(x/2)
		} else {
			memo[x] = 1 + power(3*x+1)
		}
		return memo[x]
	}

	type pair struct {
		val, power int
	}
	pairs := make([]pair, 0, hi-lo+1)
	for i := lo; i <= hi; i++ {
		pairs = append(pairs, pair{i, power(i)})
	}

  // Custom sort dengan comparator
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].power != pairs[j].power {
			return pairs[i].power < pairs[j].power
		}
		return pairs[i].val < pairs[j].val
	})

	return pairs[k-1].val
}
```
