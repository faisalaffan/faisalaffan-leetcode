# 3039 — Apply Operations To Make String Empty

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func lastNonEmptyString(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3039: Apply Operations to Make String Empty
// https://leetcode.com/problems/apply-operations-to-make-string-empty/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(lastNonEmptyString("aabcbbca"))
	fmt.Println(lastNonEmptyString("abcd"))
	fmt.Println(lastNonEmptyString("aaa"))
}

func lastNonEmptyString(s string) string {
	cnt := [26]int{}
	last := [26]int{}
	for i, ch := range s {
		idx := ch - 'a'
		cnt[idx]++
		last[idx] = i
	}
	maxFreq := 0
	for _, c := range cnt {
		if c > maxFreq {
			maxFreq = c
		}
	}
	type pair struct {
		pos int
		ch  byte
	}
	cands := []pair{}
	for i := 0; i < 26; i++ {
		if cnt[i] == maxFreq {
			cands = append(cands, pair{last[i], byte('a' + i)})
		}
	}
	// Sort by last occurrence position
  // Linear scan O(n)
	for i := 0; i < len(cands); i++ {
		for j := i + 1; j < len(cands); j++ {
			if cands[i].pos > cands[j].pos {
				cands[i], cands[j] = cands[j], cands[i]
			}
		}
	}
	ans := make([]byte, len(cands))
	for i, c := range cands {
		ans[i] = c.ch
	}
	return string(ans)
}
```
