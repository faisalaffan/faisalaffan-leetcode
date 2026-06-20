# 3442 — Maximum Difference Between Even And Odd Frequency I

## Deskripsi

**Soal:** [3442. Maximum Difference Between Even And Odd Frequency I](https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	freq := make([]int, 26)
  // Loop standar: indeks 0 sampai n-1
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
