# 0544 — Output Contest Matches

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func FindContestMatch(n int) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #544: Output Contest Matches
// https://leetcode.com/problems/output-contest-matches/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(FindContestMatch(2))
	fmt.Println(FindContestMatch(4))
	fmt.Println(FindContestMatch(8))
}

func FindContestMatch(n int) string {
	teams := make([]string, n)
	for i := 0; i < n; i++ {
		teams[i] = strconv.Itoa(i + 1)
	}

	for n > 1 {
		for i := 0; i < n/2; i++ {
			teams[i] = "(" + teams[i] + "," + teams[n-1-i] + ")"
		}
		n /= 2
	}

	return teams[0]
}
```
