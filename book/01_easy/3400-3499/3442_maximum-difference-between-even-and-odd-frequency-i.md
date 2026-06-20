# 3442 — Maximum Difference Between Even And Odd Frequency I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaximumDifferenceBetweenEvenAndOddFrequencyI(s string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3442: Maximum Difference Between Even and Odd Frequency I
// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MaximumDifferenceBetweenEvenAndOddFrequencyI("aaaaabbc"))
	fmt.Println(MaximumDifferenceBetweenEvenAndOddFrequencyI("abcabcab"))
}

// MaximumDifferenceBetweenEvenAndOddFrequencyI returns the max difference between max even-frequency and min odd-frequency in s.
// Time: O(n). Space: O(1).
func MaximumDifferenceBetweenEvenAndOddFrequencyI(s string) int {
  // Alokasi slice
	freq := make([]int, 26)
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
	}

	maxEven := 0
	minOdd := -1
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if f%2 == 0 && f > maxEven {
			maxEven = f
		} else if f%2 == 1 && (minOdd == -1 || f < minOdd) {
			minOdd = f
		}
	}

	if maxEven == 0 || minOdd == -1 {
		return 0
	}
	return maxEven - minOdd
}
```
