# 0916 — Word Subsets

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func WordSubsets(words1 []string, words2 []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O((n + m) * L)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #916: Word Subsets
// https://leetcode.com/problems/word-subsets/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "o"}))
	fmt.Println(WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"l", "e"}))
	fmt.Println(WordSubsets([]string{"amazon", "apple", "facebook", "google", "leetcode"}, []string{"e", "oo"}))
}

// Time: O((n + m) * L) | Space: O(1)
func WordSubsets(words1 []string, words2 []string) []string {
	maxCnt := [26]int{}
	for _, word := range words2 {
		var cnt [26]int
		for _, ch := range word {
			cnt[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if cnt[i] > maxCnt[i] {
				maxCnt[i] = cnt[i]
			}
		}
	}

	var ans []string
	for _, word := range words1 {
		var cnt [26]int
		for _, ch := range word {
			cnt[ch-'a']++
		}
		ok := true
		for i := 0; i < 26; i++ {
			if cnt[i] < maxCnt[i] {
				ok = false
				break
			}
		}
		if ok {
			ans = append(ans, word)
		}
	}

	return ans
}
```
