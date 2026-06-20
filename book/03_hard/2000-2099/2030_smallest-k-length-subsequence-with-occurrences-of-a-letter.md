# 2030 — Smallest K Length Subsequence With Occurrences Of A Letter

## Deskripsi

**Soal:** [2030. Smallest K Length Subsequence With Occurrences Of A Letter](https://leetcode.com/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Stack (tumpukan LIFO), Monotonic Stack (tumpukan monoton)

**Fungsi Solusi:** `func smallestKLengthSubsequence(s string, k int, letter byte, rep int) string`

> **Ide Kunci:** Monotonic Stack

## Solusi Go

```go
package main

// LeetCode #2030: Smallest K-Length Subsequence With Occurrences of a Letter
// https://leetcode.com/problems/smallest-k-length-subsequence-with-occurrences-of-a-letter/
// Difficulty: Hard
// Approach: Monotonic Stack

import "fmt"

func smallestKLengthSubsequence(s string, k int, letter byte, rep int) string {
	n := len(s)
	// Count total occurrences of letter in s
	totalLetter := 0
	for i := 0; i < n; i++ {
		if s[i] == letter {
			totalLetter++
		}
	}

  // Membuat slice untuk menyimpan hasil
	stack := make([]byte, 0, k)
	usedLetter := 0

	for i := 0; i < n; i++ {
		ch := s[i]
		// How many letters remain after current position
		remainingLetter := totalLetter
		if ch == letter {
			remainingLetter-- // decrement for current char
		}

		// While stack has elements, we can try to pop
		for len(stack) > 0 && stack[len(stack)-1] > ch {
			// If popping would make it impossible to reach k length, stop
			if n-i+len(stack)-1 < k {
				break
			}
			// If we're popping a letter, check if we'd still have enough letters
			if stack[len(stack)-1] == letter {
				if usedLetter-1+remainingLetter < rep {
					break
				}
				usedLetter--
			}
			stack = stack[:len(stack)-1]
		}

		// Add current character if we have room
		if len(stack) < k {
			stack = append(stack, ch)
			if ch == letter {
				usedLetter++
			}
		}

		// Update total remaining letter count
		if ch == letter {
			totalLetter--
		}
	}

	// If we have more than k characters, trim from end, but keep enough letters
	if len(stack) > k {
		extra := len(stack) - k
  // Membuat slice untuk menyimpan hasil
		newStack := make([]byte, 0, k)
		keptLetter := 0
  // Loop standar: indeks 0 sampai n-1
		for i := 0; i < len(stack)-extra; i++ {
			newStack = append(newStack, stack[i])
			if stack[i] == letter {
				keptLetter++
			}
		}
		// Need to add from the trimmed part if not enough letters
		if keptLetter < rep {
			for i := len(stack) - extra; i < len(stack) && keptLetter < rep; i++ {
				if stack[i] == letter {
					newStack = append(newStack, stack[i])
					keptLetter++
				}
			}
		}
		stack = newStack
	}

	return string(stack)
}

func main() {
	fmt.Println("2030. Smallest K-Length Subsequence With Occurrences of a Letter")

	// Example 1
	s1 := "leet"
	k1 := 3
	letter1 := byte('e')
	rep1 := 1
	fmt.Printf("s=%q k=%d letter=%c rep=%d → %q (expected \"eet\")\n",
		s1, k1, letter1, rep1, smallestKLengthSubsequence(s1, k1, letter1, rep1))

	// Example 2
	s2 := "leetcode"
	k2 := 4
	letter2 := byte('e')
	rep2 := 2
	fmt.Printf("s=%q k=%d letter=%c rep=%d → %q (expected \"ecde\")\n",
		s2, k2, letter2, rep2, smallestKLengthSubsequence(s2, k2, letter2, rep2))

	// Additional test
	s3 := "aaabbb"
	k3 := 3
	letter3 := byte('a')
	rep3 := 2
	fmt.Printf("s=%q k=%d letter=%c rep=%d → %q\n",
		s3, k3, letter3, rep3, smallestKLengthSubsequence(s3, k3, letter3, rep3))
}
```
