# 3541 — Find Most Frequent Vowel And Consonant

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func isVowel(b byte) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3541: Find Most Frequent Vowel and Consonant
// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindMostFrequentVowelAndConsonant("hello world"))
	fmt.Println(FindMostFrequentVowelAndConsonant("aabbccddee"))
}

// isVowel returns true if the byte is a lowercase vowel.
func isVowel(b byte) bool {
	return b == 'a' || b == 'e' || b == 'i' || b == 'o' || b == 'u'
}

// FindMostFrequentVowelAndConsonant returns the sum of max vowel frequency and max consonant frequency.
// Time: O(n). Space: O(1).
func FindMostFrequentVowelAndConsonant(s string) int {
  // Alokasi slice integer
	vowelFreq := make([]int, 26)
  // Alokasi slice integer
	consonantFreq := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= 'a' && ch <= 'z' {
			if isVowel(ch) {
				vowelFreq[ch-'a']++
			} else {
				consonantFreq[ch-'a']++
			}
		}
	}
	maxVowel := 0
	for _, v := range vowelFreq {
		if v > maxVowel {
			maxVowel = v
		}
	}
	maxConsonant := 0
	for _, v := range consonantFreq {
		if v > maxConsonant {
			maxConsonant = v
		}
	}
	return maxVowel + maxConsonant
}
```
