# 1153 — String Transforms Into Another String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func canTransform(str1 string, str2 string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1153: String Transforms Into Another String
// https://leetcode.com/problems/string-transforms-into-another-string/
// Difficulty: Hard [Paid]
//
// Given two strings str1 and str2, determine if str1 can be transformed into
// str2 by repeatedly replacing ALL occurrences of one character in the current
// string with another character. Each operation replaces every occurrence of
// a chosen character X with character Y (X and Y are lowercase letters).

import "fmt"

func main() {
	// Example: "aabcc" -> "ccdee" => true
	// a->c, b->d, c->e (since ALL c's become e after the first replacement)
	fmt.Println(canTransform("aabcc", "ccdee")) // true

	// Example: "leetcode" -> "codeleet" => false
	fmt.Println(canTransform("leetcode", "codeleet")) // false

	// Identity
	fmt.Println(canTransform("abc", "abc")) // true

	// Cycle: a->b, b->a (needs temp char)
	fmt.Println(canTransform("ab", "ba")) // true

	// All 26 chars used in str2 (no temp char)
	str1 := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(canTransform(str1, str1)) // true (identical)

	// Inconsistent mapping
	fmt.Println(canTransform("aa", "bc")) // false (a->b and a->c)

	// Different lengths
	fmt.Println(canTransform("a", "bc")) // false
}

// canTransform returns true if str1 can be converted to str2 using the allowed
// operation (replace ALL occurrences of a character in one move).
//
// Key observations:
// 1. The mapping from str1[i] to str2[i] must be consistent — each character
//    in str1 can map to at most one character in str2.
// 2. Since ALL occurrences of a character are replaced at once, the mapping
//    must be a function (each source char maps to exactly one target char).
// 3. If str2 uses all 26 lowercase letters, no character is available as a
//    temporary placeholder to break cycles. In that case, success is only
//    possible if str1 already equals str2.
// 4. Otherwise (at most 25 distinct chars in str2), cycles can be broken
//    using an unused character as intermediate, so any consistent mapping
//    is achievable.
func canTransform(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

  // Membuat map (HashMap) — pencarian O(1)
	mapping := make(map[byte]byte) // str1[i] -> str2[i]
  // Membuat map (HashMap) — pencarian O(1)
	usedInStr2 := make(map[byte]bool)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(str1); i++ {
		c1 := str1[i]
		c2 := str2[i]

		if prev, ok := mapping[c1]; ok {
			if prev != c2 {
				return false // inconsistent mapping
			}
		} else {
			mapping[c1] = c2
		}
		usedInStr2[c2] = true
	}

	// If str2 uses all 26 lowercase letters, there's no temporary character
	// to break cycles. Transformation is only possible if str1 already equals
	// str2 (mapping is identity).
	if len(usedInStr2) == 26 {
		// Check if str1 == str2
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(str1); i++ {
			if str1[i] != str2[i] {
				return false
			}
		}
	}

	return true
}
```
