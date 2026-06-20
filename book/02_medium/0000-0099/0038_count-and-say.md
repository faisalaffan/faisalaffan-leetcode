# 0038 — Count And Say

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countAndSay(n int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(2^n)  |  **Ruang:** O(2^n)


## 💻 Solusi Go

```go
package main

// LeetCode #38: Count and Say
// https://leetcode.com/problems/count-and-say/
// Difficulty: Medium

import "fmt"

func countAndSay(n int) string {
	curr := "1"

	for i := 2; i <= n; i++ {
		var next []byte
		count := 1
		for j := 1; j < len(curr); j++ {
			if curr[j] == curr[j-1] {
				count++
			} else {
				next = append(next, byte('0'+count), curr[j-1])
				count = 1
			}
		}
		next = append(next, byte('0'+count), curr[len(curr)-1])
		curr = string(next)
	}

	return curr
}

func main() {
	// Test case 1
	fmt.Println(countAndSay(4)) // "1211"
	fmt.Println(countAndSay(1)) // "1"
	fmt.Println(countAndSay(5)) // "111221"
}

// Time: O(2^n) | Space: O(2^n)
```
