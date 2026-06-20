# 2232 — Minimize Result By Adding Parentheses To Expression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func minimizeResult(expression string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n^2)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2232: Minimize Result by Adding Parentheses to Expression
// https://leetcode.com/problems/minimize-result-by-adding-parentheses-to-expression/
// Difficulty: Medium
// Time: O(n^2) | Space: O(1)

import (
	"fmt"
	"strconv"
)

func minimizeResult(expression string) string {
	plusIdx := 0
	for i, ch := range expression {
		if ch == '+' {
			plusIdx = i
			break
		}
	}
	left := expression[:plusIdx]
	right := expression[plusIdx+1:]

	minVal := int64(1 << 60)
	var bestLeft, bestRight int

  // Linear scan O(n)
	for i := 0; i < len(left); i++ {
		for j := 1; j <= len(right); j++ {
			leftPart1 := int64(1)
			if i > 0 {
				v, _ := strconv.ParseInt(left[:i], 10, 64)
				leftPart1 = v
			}
			middleLeft, _ := strconv.ParseInt(left[i:], 10, 64)
			middleRight, _ := strconv.ParseInt(right[:j], 10, 64)
			rightPart2 := int64(1)
			if j < len(right) {
				v, _ := strconv.ParseInt(right[j:], 10, 64)
				rightPart2 = v
			}

			val := leftPart1 * (middleLeft + middleRight) * rightPart2
			if val < minVal {
				minVal = val
				bestLeft = i
				bestRight = j
			}
		}
	}

	result := left[:bestLeft] + "(" + left[bestLeft:] + "+" + right[:bestRight] + ")" + right[bestRight:]
	return result
}

func main() {
	// Test case 1
	fmt.Println(minimizeResult("247+38"))
	// Expected: "2(47+38)" or "247(+3)8"

	// Test case 2
	fmt.Println(minimizeResult("12+34"))
	// Expected: "1(2+3)4"

	// Test case 3
	fmt.Println(minimizeResult("999+999"))
	// Expected: "(999+999)"
}
```
