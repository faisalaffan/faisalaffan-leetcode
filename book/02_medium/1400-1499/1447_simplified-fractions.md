# 1447 — Simplified Fractions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func simplifiedFractions(n int) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** GCD / Matematika

**Kompleksitas Waktu:** O(n^2 * log(min(i,j))) for generating all fractions  
**Kompleksitas Ruang:** O(n^2) for result

> **Untuk fresh graduate:** Kuasai dulu teknik **GCD / Matematika** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1447: Simplified Fractions
// https://leetcode.com/problems/simplified-fractions/
// Difficulty: Medium

import "fmt"
import "strconv"

func main() {
	// Test case 1
	fmt.Println(simplifiedFractions(2)) // ["1/2"]

	// Test case 2
	fmt.Println(simplifiedFractions(3)) // ["1/2","1/3","2/3"]

	// Test case 3
	fmt.Println(simplifiedFractions(4)) // ["1/2","1/3","1/4","2/3","3/4"]

	// Test case 4
	fmt.Println(simplifiedFractions(1)) // []
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Time: O(n^2 * log(min(i,j))) for generating all fractions
// Space: O(n^2) for result
func simplifiedFractions(n int) []string {
	result := make([]string, 0)

	for denominator := 2; denominator <= n; denominator++ {
		for numerator := 1; numerator < denominator; numerator++ {
			if gcd(numerator, denominator) == 1 {
				result = append(result, strconv.Itoa(numerator)+"/"+strconv.Itoa(denominator))
			}
		}
	}

	return result
}
```
