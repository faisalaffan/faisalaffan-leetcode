# 1796 — Second Largest Digit In A String

## Deskripsi

**Soal:** [1796. Second Largest Digit In A String](https://leetcode.com/problems/second-largest-digit-in-a-string/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func SecondHighest(s string) int`

## Solusi Go

```go
package main

// LeetCode #1796: Second Largest Digit in a String
// https://leetcode.com/problems/second-largest-digit-in-a-string/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func SecondHighest(s string) int {
	largest := -1
	second := -1
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			digit := int(s[i] - '0')
			if digit > largest {
				second = largest
				largest = digit
			} else if digit < largest && digit > second {
				second = digit
			}
		}
	}
	return second
}

func main() {
	fmt.Println(SecondHighest("dfa12321afd"))
	fmt.Println(SecondHighest("abc1111"))
	fmt.Println(SecondHighest("ck077"))
}
```
