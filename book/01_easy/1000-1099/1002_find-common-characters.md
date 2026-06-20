# 1002 — Find Common Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func commonChars(words []string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * m) where n = len(words), m = avg word length. Space: O(1).  
**Kompleksitas Ruang:** O(1).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1002: Find Common Characters
// https://leetcode.com/problems/find-common-characters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(commonChars([]string{"bella", "label", "roller"})) // [e l l]
	fmt.Println(commonChars([]string{"cool", "lock", "cook"}))     // [c o]
}

// commonChars returns the common characters (with multiplicity) across all words.
// Time: O(n * m) where n = len(words), m = avg word length. Space: O(1).
func commonChars(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	// Initialize with counts from first word
	freq := [26]int{}
	for _, c := range words[0] {
		freq[c-'a']++
	}

	for i := 1; i < len(words); i++ {
		currFreq := [26]int{}
		for _, c := range words[i] {
			currFreq[c-'a']++
		}
		for i := 0; i < 26; i++ {
			if currFreq[i] < freq[i] {
				freq[i] = currFreq[i]
			}
		}
	}

	result := make([]string, 0)
	for i := 0; i < 26; i++ {
		for j := 0; j < freq[i]; j++ {
			result = append(result, string(rune('a'+i)))
		}
	}
	return result
}
```
