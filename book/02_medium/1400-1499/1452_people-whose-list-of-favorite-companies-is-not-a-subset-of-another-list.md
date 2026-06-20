# 1452 — People Whose List Of Favorite Companies Is Not A Subset Of Another List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func peopleIndexes(favoriteCompanies [][]string) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n^2 * m) where n = number of people, m = avg companies per person  |  **Ruang:** O(n * m) for storing company sets

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1452: People Whose List of Favorite Companies Is Not a Subset of Another List
// https://leetcode.com/problems/people-whose-list-of-favorite-companies-is-not-a-subset-of-another-list/
// Difficulty: Medium

import "fmt"
import "sort"

func main() {
	// Test case 1
	fmt.Println(peopleIndexes([][]string{
		{"leetcode", "google", "facebook"},
		{"google", "microsoft"},
		{"google", "facebook"},
		{"google"},
		{"amazon"},
	}))
	// [0,1,4]

	// Test case 2
	fmt.Println(peopleIndexes([][]string{
		{"leetcode", "google", "facebook"},
		{"leetcode", "amazon"},
		{"facebook", "google"},
	}))
	// [0,1]

	// Test case 3
	fmt.Println(peopleIndexes([][]string{
		{"nxaqhyoprhlhvxojucbwdjggdhc", "jzobpva", "wcrrbiqiminbhyvbzrqffsxznfpy", "tvc", "inm", "qwe", "xdw"},
		{"nxaqhyoprhlhvxojucbwdjggdhc", "jzobpva"},
		{"wcrrbiqiminbhyvbzrqffsxznfpy", "jzobpva"},
		{"wcrrbiqiminbhyvbzrqffsxznfpy"},
		{"jzobpva"},
	}))
	// [0,2]
}

// Time: O(n^2 * m) where n = number of people, m = avg companies per person
// Space: O(n * m) for storing company sets
func peopleIndexes(favoriteCompanies [][]string) []int {
	// Convert each person's companies to a set
	sets := make([]map[string]bool, len(favoriteCompanies))
	for i, companies := range favoriteCompanies {
		sets[i] = make(map[string]bool)
		for _, c := range companies {
			sets[i][c] = true
		}
	}

	// Sort by company list size descending to check larger lists first
  // Alokasi slice
	indices := make([]int, len(favoriteCompanies))
  // Range loop
	for i := range indices {
		indices[i] = i
	}
  // Custom sort
	sort.Slice(indices, func(i, j int) bool {
		return len(favoriteCompanies[indices[i]]) > len(favoriteCompanies[indices[j]])
	})

	result := []int{}
	for _, i := range indices {
		isSubset := false
		for _, j := range indices {
			if i == j {
				continue
			}
			if isSubsetOf(sets[i], sets[j]) {
				isSubset = true
				break
			}
		}
		if !isSubset {
			result = append(result, i)
		}
	}

  // Sort O(n log n)
	sort.Ints(result)
	return result
}

func isSubsetOf(a, b map[string]bool) bool {
	if len(a) > len(b) {
		return false
	}
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}
```
