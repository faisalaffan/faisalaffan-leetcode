# 3348 — Smallest Divisible Digit Product Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func gcd64(a, b int64) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3348: Smallest Divisible Digit Product II
// https://leetcode.com/problems/smallest-divisible-digit-product-ii/
// Difficulty: Hard
//
// Given num (string) and t (int64), find smallest zero-free number >= num
// whose digit product is divisible by t.

import "fmt"

func main() {
	fmt.Println(SmallestDivisibleDigitProductIi("123", 12))
}

func gcd64(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func SmallestDivisibleDigitProductIi(num string, t int64) string {
	if t == 1 {
		b := []byte(num)
  // Range loop
		for i := range b {
			if b[i] == '0' {
				b[i] = '1'
			}
		}
		return string(b)
	}

	tmp := t
	for f := int64(9); f > 1; f-- {
		for tmp%f == 0 {
			tmp /= f
		}
	}
	if tmp > 1 {
		return "-1"
	}

	n := len(num)
  // Alokasi slice
	leftT := make([]int64, n+1)
	leftT[0] = t
	firstZero := n - 1
	for i := 0; i < n; i++ {
		if num[i] == '0' {
			firstZero = i
			break
		}
		d := int64(num[i] - '0')
		g := gcd64(leftT[i], d)
		leftT[i+1] = leftT[i] / g
	}
	if leftT[n] == 1 {
		return num
	}

	s := []byte(num)
	for i := firstZero; i >= 0; i-- {
		for d := s[i] + 1; d <= '9'; d++ {
			g := gcd64(leftT[i], int64(d-'0'))
			tt := leftT[i] / g
			for j := n - 1; j > i; j-- {
				if tt == 1 {
					s[j] = '1'
					continue
				}
				for k := 9; k >= 2; k-- {
					if tt%int64(k) == 0 {
						s[j] = byte('0' + k)
						tt /= int64(k)
						break
					}
				}
			}
			if tt == 1 {
				s[i] = d
				return string(s)
			}
		}
	}

	var factors []byte
	tt := t
	for f := int64(9); f >= 2; f-- {
		for tt%f == 0 {
			factors = append(factors, byte('0'+f))
			tt /= f
		}
	}
	for i, j := 0, len(factors)-1; i < j; i, j = i+1, j-1 {
		factors[i], factors[j] = factors[j], factors[i]
	}
	padLen := n + 1 - len(factors)
	if padLen < 0 {
		padLen = 0
	}
	result := make([]byte, 0, padLen+len(factors))
	for i := 0; i < padLen; i++ {
		result = append(result, '1')
	}
	result = append(result, factors...)
	return string(result)
}
```
