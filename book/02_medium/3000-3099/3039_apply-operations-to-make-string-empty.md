# 3039 — Apply Operations To Make String Empty

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func lastNonEmptyString(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

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
  // Loop linear O(n): iterasi setiap elemen
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
