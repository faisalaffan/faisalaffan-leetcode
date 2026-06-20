# 0241 — Different Ways To Add Parentheses

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func diffWaysToCompute(expression string) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(2^n), Space: O(2^n) for result storage  
**Kompleksitas Ruang:** O(2^n) for result storage

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #241: Different Ways to Add Parentheses
// https://leetcode.com/problems/different-ways-to-add-parentheses/
// Difficulty: Medium
// Time: O(2^n), Space: O(2^n) for result storage

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(expression string) []int {
	if isNumber(expression) {
		num, _ := strconv.Atoi(expression)
		return []int{num}
	}

	result := []int{}
	for i, ch := range expression {
		if ch == '+' || ch == '-' || ch == '*' {
			left := diffWaysToCompute(expression[:i])
			right := diffWaysToCompute(expression[i+1:])

			for _, l := range left {
				for _, r := range right {
					switch ch {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}

	return result
}

func isNumber(s string) bool {
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
	fmt.Println(diffWaysToCompute("3"))
}
```
