# 2023 — Number Of Pairs Of Strings With Concatenation Equal To Target

## Deskripsi

**Soal:** [2023. Number Of Pairs Of Strings With Concatenation Equal To Target](https://leetcode.com/problems/number-of-pairs-of-strings-with-concatenation-equal-to-target/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * L)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func numOfPairs(nums []string, target string) int`

## Solusi Go

```go
package main

// LeetCode #2023: Number of Pairs of Strings With Concatenation Equal to Target
// https://leetcode.com/problems/number-of-pairs-of-strings-with-concatenation-equal-to-target/
// Difficulty: Medium
// Time: O(n * L) | Space: O(n)

import "fmt"

func numOfPairs(nums []string, target string) int {
  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[string]int)
	count := 0

	for _, num := range nums {
		// Check if any prefix/suffix of target matches
		for i := 1; i < len(target); i++ {
			prefix := target[:i]
			suffix := target[i:]
			if num == prefix {
				count += freq[suffix]
			}
			if num == suffix {
				count += freq[prefix]
			}
		}
		freq[num]++
	}

	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numOfPairs([]string{"777", "7", "77", "77"}, "777"))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", numOfPairs([]string{"123", "4", "12", "34"}, "1234"))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", numOfPairs([]string{"1", "1", "1"}, "11"))
	// Expected: 6
}
```
