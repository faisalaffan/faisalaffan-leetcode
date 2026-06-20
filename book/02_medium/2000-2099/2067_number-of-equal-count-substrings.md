# 2067 — Number Of Equal Count Substrings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func equalCountSubstrings(s string, count int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * alphabet)  |  **Ruang:** O(alphabet)


## 💻 Solusi Go

```go
package main

// LeetCode #2067: Number of Equal Count Substrings
// https://leetcode.com/problems/number-of-equal-count-substrings/
// Difficulty: Medium [Paid]
// Time: O(n * alphabet) | Space: O(alphabet)

import "fmt"

func equalCountSubstrings(s string, count int) int {
	result := 0
	// Try different numbers of distinct characters
	for distinct := 1; distinct <= 26 && distinct*count <= len(s); distinct++ {
  // Alokasi slice
		freq := make([]int, 26)
		unique := 0
		exactCount := 0

  // Linear scan O(n)
		for i := 0; i < len(s); i++ {
			idx := int(s[i] - 'a')
			if freq[idx] == 0 {
				unique++
			}
			freq[idx]++
			if freq[idx] == count {
				exactCount++
			}

			// Remove leftmost when window too big
			if i >= distinct*count {
				left := int(s[i-distinct*count] - 'a')
				if freq[left] == count {
					exactCount--
				}
				freq[left]--
				if freq[left] == 0 {
					unique--
				}
			}

			if unique == distinct && exactCount == distinct {
				result++
			}
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", equalCountSubstrings("aaabc", 3))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", equalCountSubstrings("abcd", 1))
	// Expected: 10

	// Test case 3
	fmt.Println("Test 3:", equalCountSubstrings("aabbcc", 2))
	// Expected: 6
}
```
