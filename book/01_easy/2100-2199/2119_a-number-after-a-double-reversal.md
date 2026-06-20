# 2119 — A Number After A Double Reversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func ANumberAfterADoubleReversal(num int) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2119: A Number After a Double Reversal
// https://leetcode.com/problems/a-number-after-a-double-reversal/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ANumberAfterADoubleReversal(526))  // true
	fmt.Println(ANumberAfterADoubleReversal(1800)) // false
	fmt.Println(ANumberAfterADoubleReversal(0))    // true
}

// Time: O(1), Space: O(1)
func ANumberAfterADoubleReversal(num int) bool {
	// Reversing twice yields the same iff num has no trailing zeros
	return num == 0 || num%10 != 0
}
```
