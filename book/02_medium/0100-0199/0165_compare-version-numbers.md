# 0165 — Compare Version Numbers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func compareVersion(version1 string, version2 string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n+m), Space: O(n+m)  |  **Ruang:** O(n+m)


## 💻 Solusi Go

```go
package main

// LeetCode #165: Compare Version Numbers
// https://leetcode.com/problems/compare-version-numbers/
// Difficulty: Medium
// Time: O(n+m), Space: O(n+m)

import (
	"fmt"
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	n := len(v1)
	if len(v2) > n {
		n = len(v2)
	}

	for i := 0; i < n; i++ {
		num1, num2 := 0, 0
		if i < len(v1) {
			num1, _ = strconv.Atoi(v1[i])
		}
		if i < len(v2) {
			num2, _ = strconv.Atoi(v2[i])
		}

		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
	}

	return 0
}

func main() {
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
	fmt.Println(compareVersion("0.1", "1.1"))
}
```
