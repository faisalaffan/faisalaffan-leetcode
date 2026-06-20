# 2081 — Sum Of K Mirror Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func kMirror(k int, n int) int64
```

> **💡 Hint:** Generate palindromes in base k, check decimal palindrome

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2081: Sum of k-Mirror Numbers
// https://leetcode.com/problems/sum-of-k-mirror-numbers/
// Difficulty: Hard
// Approach: Generate palindromes in base k, check decimal palindrome

import (
	"fmt"
	"strconv"
)

func kMirror(k int, n int) int64 {
	var sum int64
	count := 0

	// Generate palindromes in base k
	length := 1
	for count < n {
		// Generate palindromes of given length in base k
		// For odd length, the first half determines the palindrome
		// For even length, the first half determines the palindrome

		halfLen := (length + 1) / 2
		start := pow(k, halfLen-1)
		end := pow(k, halfLen)

		for base := start; base < end && count < n; base++ {
			// Build the full palindrome in base k
			pal := buildPalindrome(base, length%2 == 1, k)
			if pal == 0 {
				continue
			}

			// Convert to decimal and check if it's a palindrome in decimal
			dec := toDecimal(pal, k)
			if isDecimalPalindrome(dec) {
				sum += dec
				count++
				if count >= n {
					break
				}
			}
		}
		length++
	}

	return sum
}

// pow computes k^e
func pow(k, e int) int {
	res := 1
	for i := 0; i < e; i++ {
		res *= k
	}
	return res
}

// buildPalindrome builds a number whose decimal representation is the palindrome in base k.
// The palindrome is formed by mirroring the half around an optional middle digit.
// Example: half=7("10") in base 7, even length → "1001" → returns 1001.
func buildPalindrome(half int, oddLen bool, k int) int {
	if half == 0 {
		return 0
	}
	// Extract digits of half in base k (least significant first)
	tmp := half
  // Alokasi slice integer
	digits := make([]int, 0)
	for tmp > 0 {
		digits = append(digits, tmp%k)
		tmp /= k
	}
	// Reverse to get most-significant first
	m := len(digits)
	for i, j := 0, m-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	// Build the palindrome as a decimal number whose digits are the base-k digits
	result := 0
	for _, d := range digits {
		result = result*10 + d
	}
	start := m - 1
	if oddLen {
		start = m - 2
	}
	for i := start; i >= 0; i-- {
		result = result*10 + digits[i]
	}
	return result
}

// toDecimal converts a number (whose decimal digits represent base-k digits) to a decimal value.
// Example: toDecimal(1001, 7) where 1001 decimal represents "1001" in base 7 = 1*343+0*49+0*7+1 = 344.
func toDecimal(num, k int) int64 {
	var dec int64 = 0
	mult := int64(1)
	for num > 0 {
		dec += int64(num%10) * mult
		num /= 10
		mult *= int64(k)
	}
	return dec
}

// isDecimalPalindrome checks if a decimal number is a palindrome
func isDecimalPalindrome(num int64) bool {
	s := strconv.FormatInt(num, 10)
	i, j := 0, len(s)-1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println("2081. Sum of k-Mirror Numbers")

	// Example 1
	k1 := 2
	n1 := 5
	fmt.Printf("k=%d n=%d → %d (expected 25)\n", k1, n1, kMirror(k1, n1))

	// Example 2
	k2 := 3
	n2 := 7
	fmt.Printf("k=%d n=%d → %d (expected 499)\n", k2, n2, kMirror(k2, n2))

	// Example 3 (n=17, not 7)
	k3 := 7
	n3 := 17
	fmt.Printf("k=%d n=%d → %d (expected 20379000)\n", k3, n3, kMirror(k3, n3))
}
```
