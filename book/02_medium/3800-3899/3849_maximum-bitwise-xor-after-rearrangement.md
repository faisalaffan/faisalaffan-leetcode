# 3849 — Maximum Bitwise Xor After Rearrangement

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MaximumBitwiseXorAfterRearrangement(s string, t string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N)  |  **Ruang:** O(1)


## 💻 Solusi Go

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
