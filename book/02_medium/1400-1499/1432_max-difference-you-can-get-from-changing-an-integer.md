# 1432 — Max Difference You Can Get From Changing An Integer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxDiff(num int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = number of digits  
**Kompleksitas Ruang:** O(n) for string conversion

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1432: Max Difference You Can Get From Changing an Integer
// https://leetcode.com/problems/max-difference-you-can-get-from-changing-an-integer/
// Difficulty: Medium

import "fmt"
import "strconv"

func main() {
	// Test case 1
	fmt.Println(maxDiff(555)) // 888

	// Test case 2
	fmt.Println(maxDiff(9)) // 8

	// Test case 3
	fmt.Println(maxDiff(123456)) // 820000

	// Test case 4
	fmt.Println(maxDiff(10000)) // 20000

	// Test case 5
	fmt.Println(maxDiff(9288)) // 8700
}

// Time: O(n) where n = number of digits
// Space: O(n) for string conversion
func maxDiff(num int) int {
	s := strconv.Itoa(num)
	digits := []byte(s)

	// Find max: replace first non-'9' with '9'
	maxDigits := make([]byte, len(digits))
	copy(maxDigits, digits)
	targetMax := byte(0)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(maxDigits); i++ {
		if maxDigits[i] != '9' {
			targetMax = maxDigits[i]
			break
		}
	}
	if targetMax != 0 {
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(maxDigits); i++ {
			if maxDigits[i] == targetMax {
				maxDigits[i] = '9'
			}
		}
	}
	maxVal, _ := strconv.Atoi(string(maxDigits))

	// Find min: replace first non-'0'/'1' appropriately
	minDigits := make([]byte, len(digits))
	copy(minDigits, digits)

	// If first digit is not '1', replace it with '1'
	// Otherwise find first digit > '1' to replace with '0' (but not leading)
	targetMin := byte(0)
	replacementMin := byte(0)

	if minDigits[0] != '1' {
		targetMin = minDigits[0]
		replacementMin = '1'
	} else {
		for i := 1; i < len(minDigits); i++ {
			if minDigits[i] != '0' && minDigits[i] != '1' {
				targetMin = minDigits[i]
				replacementMin = '0'
				break
			}
		}
	}

	if targetMin != 0 {
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(minDigits); i++ {
			if minDigits[i] == targetMin {
				minDigits[i] = replacementMin
			}
		}
	}
	minVal, _ := strconv.Atoi(string(minDigits))

	return maxVal - minVal
}
```
