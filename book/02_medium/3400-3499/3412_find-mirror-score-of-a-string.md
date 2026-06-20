# 3412 — Find Mirror Score Of A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func calculateScore(s string) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Stack

**Waktu:** O(n) Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Stack** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3412: Find Mirror Score of a String
// https://leetcode.com/problems/find-mirror-score-of-a-string/
// Difficulty: Medium
// Time: O(n) Space: O(n)

import "fmt"

func calculateScore(s string) int64 {
  // Matriks 2D
	stacks := make([][]int, 26)
	var ans int64
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		ch := int(s[i] - 'a')
		mirror := 25 - ch
		if len(stacks[mirror]) > 0 {
			j := stacks[mirror][len(stacks[mirror])-1]
			stacks[mirror] = stacks[mirror][:len(stacks[mirror])-1]
			ans += int64(i - j)
		} else {
			stacks[ch] = append(stacks[ch], i)
		}
	}
	return ans
}

func main() {
	fmt.Println(calculateScore("aczzx")) // 5
	fmt.Println(calculateScore("abcdef")) // 0
}
```
