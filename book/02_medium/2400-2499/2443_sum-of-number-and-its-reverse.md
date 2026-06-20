# 2443 — Sum Of Number And Its Reverse

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func sumOfNumberAndReverse(num int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(num)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2443: Sum of Number and Its Reverse
// https://leetcode.com/problems/sum-of-number-and-its-reverse/
// Difficulty: Medium
// Time: O(num) | Space: O(1)
// Iterate from 0 to num, check if i + reverse(i) == num.

import "fmt"

func main() {
	fmt.Println(sumOfNumberAndReverse(443)) // true (241+142=443)
	fmt.Println(sumOfNumberAndReverse(63))  // false
	fmt.Println(sumOfNumberAndReverse(181)) // true (90+9=99? wait) (140+41=181)
}

func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		if i+reverse(i) == num {
			return true
		}
	}
	return false
}

func reverse(n int) int {
	r := 0
	for n > 0 {
		r = r*10 + n%10
		n /= 10
	}
	return r
}
```
