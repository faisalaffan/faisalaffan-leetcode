# 2182 — Construct String With Repeat Limit

## Deskripsi

**Soal:** [2182. Construct String With Repeat Limit](https://leetcode.com/problems/construct-string-with-repeat-limit/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func repeatLimitedString(s string, repeatLimit int) string`

## Solusi Go

```go
package main

// LeetCode #2182: Construct String With Repeat Limit
// https://leetcode.com/problems/construct-string-with-repeat-limit/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func repeatLimitedString(s string, repeatLimit int) string {
  // Membuat slice untuk menyimpan hasil
	count := make([]int, 26)
	for _, ch := range s {
		count[ch-'a']++
	}

  // Membuat slice untuk menyimpan hasil
	result := make([]byte, 0, len(s))
	for i := 25; i >= 0; {
		if count[i] == 0 {
			i--
			continue
		}

		use := min(count[i], repeatLimit)
		for k := 0; k < use; k++ {
			result = append(result, byte('a'+i))
		}
		count[i] -= use

		if count[i] > 0 {
			j := i - 1
			for j >= 0 && count[j] == 0 {
				j--
			}
			if j < 0 {
				break
			}
			result = append(result, byte('a'+j))
			count[j]--
		} else {
			i--
		}
	}

	return string(result)
}

func main() {
	// Test case 1
	fmt.Println(repeatLimitedString("cczazcc", 3))
	// Expected: "zzcccac"

	// Test case 2
	fmt.Println(repeatLimitedString("aababab", 2))
	// Expected: "bbabaa"

	// Test case 3
	fmt.Println(repeatLimitedString("robnsdvpuxbapuqgopqvxdrchivlifeepy", 2))
	// Expected: "yxxvvuvusrrqqppopponliihgfeeddcba"
}
```
