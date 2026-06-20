# 1189 — Maximum Number Of Balloons

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func maxNumberOfBalloons(text string) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1189: Maximum Number of Balloons
// https://leetcode.com/problems/maximum-number-of-balloons/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))           // 1
	fmt.Println(maxNumberOfBalloons("loonbalxballpoon"))    // 2
	fmt.Println(maxNumberOfBalloons("leetcode"))            // 0
}

// LeetCode submission: maxNumberOfBalloons
func maxNumberOfBalloons(text string) int {
	count := [26]int{}
  // Linear scan O(n)
	for i := 0; i < len(text); i++ {
		count[text[i]-'a']++
	}
	ans := count[1]           // b
	ans = min(ans, count[0])  // a
	ans = min(ans, count[11]/2) // l (needs 2)
	ans = min(ans, count[14]/2) // o (needs 2)
	ans = min(ans, count[13]) // n
	return ans
}
```
