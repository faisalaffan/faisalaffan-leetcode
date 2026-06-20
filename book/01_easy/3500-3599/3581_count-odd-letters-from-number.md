# 3581 — Count Odd Letters From Number

## Deskripsi

**Soal:** [3581. Count Odd Letters From Number](https://leetcode.com/problems/count-odd-letters-from-number/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n) - number of digits of n  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3581: Count Odd Letters from Number
// https://leetcode.com/problems/count-odd-letters-from-number/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(CountOddLettersFromNumber(41))
	fmt.Println(CountOddLettersFromNumber(20))
	fmt.Println(CountOddLettersFromNumber(7))
}

// Time: O(log n) - number of digits of n
// Space: O(1)
func CountOddLettersFromNumber(n int) int {
	digitWords := [10]string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	mask := 0
	for n > 0 {
		d := n % 10
		word := digitWords[d]
		for _, ch := range word {
			mask ^= 1 << (ch - 'a')
		}
		n /= 10
	}

	ans := 0
	for mask > 0 {
		ans += mask & 1
		mask >>= 1
	}
	return ans
}
```
