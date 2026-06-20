# 3941 — Password Strength

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func PasswordStrength(password string) int
```

> **💡 Hint:** Track distinct characters per category. Sum points:

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3941: Password Strength
// https://leetcode.com/problems/password-strength/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Track distinct characters per category. Sum points:
// lowercase=1, uppercase=2, digit=3, special=5.

import "fmt"

func PasswordStrength(password string) int {
  // Membuat map (HashMap) — pencarian O(1)
	lower := make(map[byte]bool)
  // Membuat map (HashMap) — pencarian O(1)
	upper := make(map[byte]bool)
  // Membuat map (HashMap) — pencarian O(1)
	digit := make(map[byte]bool)
  // Membuat map (HashMap) — pencarian O(1)
	special := make(map[byte]bool)

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(password); i++ {
		ch := password[i]
		if ch >= 'a' && ch <= 'z' {
			lower[ch] = true
		} else if ch >= 'A' && ch <= 'Z' {
			upper[ch] = true
		} else if ch >= '0' && ch <= '9' {
			digit[ch] = true
		} else {
			// Special: ! @ # $
			special[ch] = true
		}
	}

	return len(lower)*1 + len(upper)*2 + len(digit)*3 + len(special)*5
}

func main() {
	// Example 1
	fmt.Println(PasswordStrength("aA1!")) // Expected: 11

	// Example 2
	fmt.Println(PasswordStrength("bbB11#")) // Expected: 11
}
```
