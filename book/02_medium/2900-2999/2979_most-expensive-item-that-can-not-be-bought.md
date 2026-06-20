# 2979 — Most Expensive Item That Can Not Be Bought

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func mostExpensiveItem(primeOne int, primeTwo int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2979: Most Expensive Item That Can Not Be Bought
// https://leetcode.com/problems/most-expensive-item-that-can-not-be-bought/
// Difficulty: Medium [Paid]
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(mostExpensiveItem(2, 5))
	fmt.Println(mostExpensiveItem(3, 7))
}

func mostExpensiveItem(primeOne int, primeTwo int) int {
	return primeOne*primeTwo - primeOne - primeTwo
}
```
