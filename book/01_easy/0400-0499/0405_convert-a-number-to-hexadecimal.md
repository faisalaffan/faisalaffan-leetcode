# 0405 — Convert A Number To Hexadecimal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ConvertANumberToHexadecimal(num int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #405: Convert a Number to Hexadecimal
// https://leetcode.com/problems/convert-a-number-to-hexadecimal/
// Difficulty: Easy

import "fmt"

// Time: O(1), Space: O(1)
func ConvertANumberToHexadecimal(num int) string {
	if num == 0 {
		return "0"
	}
	hex := "0123456789abcdef"
	var result []byte
	// Use uint32 to handle negative numbers via two's complement.
	n := uint32(num)
	for n > 0 {
		result = append([]byte{hex[n&0xf]}, result...)
		n >>= 4
	}
	return string(result)
}

func main() {
	fmt.Println(ConvertANumberToHexadecimal(26))
	fmt.Println(ConvertANumberToHexadecimal(-1))
	fmt.Println(ConvertANumberToHexadecimal(0))
}
```
