# 0592 — Fraction Addition And Subtraction

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func FractionAddition(expression string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(n) where n = length of expression  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #592: Fraction Addition and Subtraction
// https://leetcode.com/problems/fraction-addition-and-subtraction/
// Difficulty: Medium
// Time: O(n) where n = length of expression
// Space: O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FractionAddition("-1/2+1/2"))
	fmt.Println(FractionAddition("-1/2+1/2+1/3"))
	fmt.Println(FractionAddition("1/3-1/2"))
}

func FractionAddition(expression string) string {
	num := 0
	den := 1
	i := 0
	n := len(expression)

	for i < n {
		sign := 1
		if expression[i] == '-' {
			sign = -1
			i++
		} else if expression[i] == '+' {
			sign = 1
			i++
		}

		// Parse numerator
		j := i
		for j < n && expression[j] >= '0' && expression[j] <= '9' {
			j++
		}
		currNum, _ := strconv.Atoi(expression[i:j])
		i = j + 1 // skip '/'

		// Parse denominator
		j = i
		for j < n && expression[j] >= '0' && expression[j] <= '9' {
			j++
		}
		currDen, _ := strconv.Atoi(expression[i:j])
		i = j

		currNum *= sign
		num = num*currDen + currNum*den
		den = den * currDen

		g := gcd(abs(num), abs(den))
		num /= g
		den /= g
	}

	if den < 0 {
		num = -num
		den = -den
	}

	return fmt.Sprintf("%d/%d", num, den)
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
