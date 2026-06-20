# 1111 — Maximum Nesting Depth Of Two Valid Parentheses Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func maxDepthAfterSplit(seq string) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #1111: Maximum Nesting Depth of Two Valid Parentheses Strings
// https://leetcode.com/problems/maximum-nesting-depth-of-two-valid-parentheses-strings/
// Difficulty: Medium
//
// Approach: Assign '(' to group A or B based on even/odd depth.
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(maxDepthAfterSplit("(()())")) // [0,1,1,1,1,0] or similar
	fmt.Println(maxDepthAfterSplit("()(())()")) // [0,0,0,1,1,0,0,0]
}

func maxDepthAfterSplit(seq string) []int {
  // Alokasi slice
	result := make([]int, len(seq))
	depth := 0

	for i, c := range seq {
		if c == '(' {
			depth++
			result[i] = depth % 2
		} else {
			result[i] = depth % 2
			depth--
		}
	}

	return result
}
```
