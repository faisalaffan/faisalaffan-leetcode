# 3849 — Maximum Bitwise Xor After Rearrangement

## Deskripsi

**Soal:** [3849. Maximum Bitwise Xor After Rearrangement](https://leetcode.com/problems/maximum-bitwise-xor-after-rearrangement/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MaximumBitwiseXorAfterRearrangement(s string, t string) string`

> **Ide Kunci:** Count 0s and 1s in t. Greedily match opposite bits for max XOR.

## Solusi Go

```go
package main

// LeetCode #3849: Maximum Bitwise XOR After Rearrangement
// https://leetcode.com/problems/maximum-bitwise-xor-after-rearrangement/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Count 0s and 1s in t. Greedily match opposite bits for max XOR.

import "fmt"

func MaximumBitwiseXorAfterRearrangement(s string, t string) string {
	ones, zeros := 0, 0
	for _, ch := range t {
		if ch == '1' {
			ones++
		} else {
			zeros++
		}
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]byte, len(s))
	for i, ch := range s {
		if ch == '1' {
			if zeros > 0 {
				ans[i] = '1'
				zeros--
			} else {
				ans[i] = '0'
				ones--
			}
		} else {
			if ones > 0 {
				ans[i] = '1'
				ones--
			} else {
				ans[i] = '0'
				zeros--
			}
		}
	}

	return string(ans)
}

func main() {
	// Example 1
	fmt.Println(MaximumBitwiseXorAfterRearrangement("101", "011")) // Expected: "110"

	// Example 2
	fmt.Println(MaximumBitwiseXorAfterRearrangement("0110", "1110")) // Expected: "1101"

	// Example 3
	fmt.Println(MaximumBitwiseXorAfterRearrangement("0101", "1001")) // Expected: "1111"
}
```
