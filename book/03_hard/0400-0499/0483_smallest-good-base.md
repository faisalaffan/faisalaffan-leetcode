# 0483 — Smallest Good Base

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func smallestGoodBase(n string) string
```

> **💡 Hint:** Binary search per length. For a number n represented as string,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #483: Smallest Good Base
// https://leetcode.com/problems/smallest-good-base/
// Difficulty: Hard
// Approach: Binary search per length. For a number n represented as string,
// find the smallest base k such that n = 1 + k + k^2 + ... + k^(m-1).
// The base k is smallest when m (number of digits) is largest.
// Max m is about log2(n)+1, min m is 2.

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println("483 - Smallest Good Base")

	// Test cases
	testCases := []string{"13", "4681", "1000000000000000000", "3", "15", "1", "2251799813685247"}
	for _, tc := range testCases {
		result := smallestGoodBase(tc)
		fmt.Printf("n=%s -> base=%s\n", tc, result)
	}
	// Expected: "13"->"3", "4681"->"8", "3"->"2", "15"->"2"
	// "1"->"" (or "1"? LeetCode says: For n=1, return "1" (since 1 in any base > 1 is just "1"))
}

func smallestGoodBase(n string) string {
	num, _ := new(big.Int).SetString(n, 10)

	// Compare with 1 (big.Int)
	one := big.NewInt(1)
	if num.Cmp(one) == 0 {
		return "1"
	}

	// n = 1 + k + k^2 + ... + k^(m-1)
	// The max possible m is when k=2: n >= 1 + 2 + ... + 2^(m-1) = 2^m - 1
	// so m <= log2(n+1) ~ 60 for n up to 10^18
	maxM := num.BitLen() + 1 // log2(n) + 1

	// Try each m from max down to 2 (larger m -> smaller base k)
	for m := maxM; m >= 2; m-- {
		// Binary search for k
		low := big.NewInt(2)
		high := new(big.Int).Sub(num, big.NewInt(1))

		for low.Cmp(high) <= 0 {
			mid := new(big.Int).Add(low, high)
			mid.Div(mid, big.NewInt(2))

			// Compute sum = 1 + mid + mid^2 + ... + mid^(m-1)
			sum := geometricSum(mid, m)

			cmp := sum.Cmp(num)
			if cmp == 0 {
				return mid.String()
			} else if cmp < 0 {
				low.Add(mid, big.NewInt(1))
			} else {
				high.Sub(mid, big.NewInt(1))
			}
		}
	}

	return strconv.FormatInt(num.Int64()-1, 10)
}

// geometricSum returns sum_{i=0}^{m-1} base^i
func geometricSum(base *big.Int, m int) *big.Int {
	// Use Horner's method: (((1*base + 1)*base + 1)*base + 1)
	// Actually: 1 + base*(1 + base*(1 + base*(...)))
	if m == 0 {
		return big.NewInt(0)
	}

	// Compute using the formula (base^m - 1) / (base - 1)
	pow := new(big.Int).Exp(base, big.NewInt(int64(m)), nil)
	pow.Sub(pow, big.NewInt(1))

	baseMinus1 := new(big.Int).Sub(base, big.NewInt(1))
	return pow.Div(pow, baseMinus1)
}
```
