# 0866 — Prime Palindrome

## Deskripsi

**Soal:** [0866. Prime Palindrome](https://leetcode.com/problems/prime-palindrome/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N * sqrt(N))  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #866: Prime Palindrome
// https://leetcode.com/problems/prime-palindrome/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PrimePalindrome(6))
	fmt.Println(PrimePalindrome(8))
	fmt.Println(PrimePalindrome(13))
}

// Time: O(N * sqrt(N)) | Space: O(1)
func PrimePalindrome(N int) int {
	if N <= 2 {
		return 2
	}

	for {
		// Even-length palindromes > 11 are divisible by 11, skip them
		if N >= 10000000 && N < 100000000 {
			N = 100000000
		}

		if isPalindrome(N) && isPrime(N) {
			return N
		}
		N++
	}
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	orig, rev := x, 0
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return orig == rev
}

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}
	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}
```
