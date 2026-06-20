# 3081 — Replace Question Marks In String To Minimize Its Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minimizeStringValue(s string) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3081: Replace Question Marks in String to Minimize Its Value
// https://leetcode.com/problems/replace-question-marks-in-string-to-minimize-its-value/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import (
	"fmt"
)

func main() {
	fmt.Println(minimizeStringValue("abc?def"))
	fmt.Println(minimizeStringValue("???"))
	fmt.Println(minimizeStringValue("a?b?c?"))
}

func minimizeStringValue(s string) string {
	cnt := [26]int{}
	qpos := []int{}
	for i, ch := range s {
		if ch == '?' {
			qpos = append(qpos, i)
		} else {
			cnt[ch-'a']++
		}
	}
	type pair struct {
		ch   byte
		pos  int
	}
	toFill := make([]byte, len(qpos))
	for i, pos := range qpos {
		bestCh := byte('a')
		bestCost := cnt[0]
		for c := 1; c < 26; c++ {
			if cnt[c] < bestCost {
				bestCost = cnt[c]
				bestCh = byte('a' + c)
			}
		}
		toFill[i] = bestCh
		cnt[bestCh-'a']++
		_ = pos
	}
	ans := []byte(s)
	for i, pos := range qpos {
		ans[pos] = toFill[i]
	}
	return string(ans)
}
```
