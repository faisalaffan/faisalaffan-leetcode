# 2217 — Find Palindrome With Fixed Length

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu adalah memeriksa apakah string tersebut palindrome — dibaca sama dari depan dan belakang. Abaikan non-alfanumerik dan case.

**Cara berpikir:** Two Pointer — kiri dan kanan. Skip non-alfanumerik. Bandingkan.

**Fungsi Solusi:** `func kthPalindrome(queries []int, intLength int) []int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n * len)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2217: Find Palindrome With Fixed Length
// https://leetcode.com/problems/find-palindrome-with-fixed-length/
// Difficulty: Medium
// Time: O(n * len) | Space: O(1)

import "fmt"

func kthPalindrome(queries []int, intLength int) []int64 {
	halfLen := (intLength + 1) / 2
	start := 1
	for i := 1; i < halfLen; i++ {
		start *= 10
	}

  // Alokasi slice
	result := make([]int64, len(queries))
	for i, q := range queries {
		val := int64(start + q - 1)
		if val > int64(start*10-1) {
			result[i] = -1
			continue
		}
		s := fmt.Sprintf("%d", val)
		runes := []rune(s)
		// Mirror: for odd length, skip last char
		mirrorLen := halfLen - (intLength % 2)
		for j := mirrorLen - 1; j >= 0; j-- {
			s += string(runes[j])
		}
		fmt.Sscanf(s, "%d", &val)
		result[i] = val
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(kthPalindrome([]int{1, 2, 3, 4, 5, 90}, 3))
	// Expected: [101, 111, 121, 131, 141, 999]

	// Test case 2
	fmt.Println(kthPalindrome([]int{2, 4, 6}, 4))
	// Expected: [1111, 1331, 1551]
}
```
