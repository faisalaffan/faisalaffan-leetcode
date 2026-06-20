# 3860 — Unique Email Groups

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func UniqueEmailGroups(emails []string) int
```

> **💡 Hint:** Normalize each email by processing local part (ignore dots,

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N * M)  
**Kompleksitas Ruang:** O(N * M)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3860: Unique Email Groups
// https://leetcode.com/problems/unique-email-groups/
// Difficulty: Medium [Paid]
// Time: O(N * M) | Space: O(N * M)
// Approach: Normalize each email by processing local part (ignore dots,
// ignore after +) and lowercase everything. Count unique normalized forms.

import (
	"fmt"
	"strings"
)

func UniqueEmailGroups(emails []string) int {
  // Membuat map (HashMap) — pencarian O(1)
	seen := make(map[string]bool)

	for _, email := range emails {
		parts := strings.SplitN(email, "@", 2)
		local, domain := parts[0], parts[1]

		var normLocal strings.Builder
		for _, ch := range local {
			if ch == '+' {
				break
			}
			if ch != '.' {
				normLocal.WriteRune(ch)
			}
		}

		normalized := strings.ToLower(normLocal.String()) + "@" + strings.ToLower(domain)
		seen[normalized] = true
	}

	return len(seen)
}

func main() {
	// Example 1
	emails1 := []string{
		"test.email+alex@leetcode.com",
		"test.e.mail+bob.cathy@leetcode.com",
		"testemail+david@lee.tcode.com",
	}
	fmt.Println(UniqueEmailGroups(emails1)) // Expected: 2

	// Example 2
	emails2 := []string{"A@B.com", "a@b.com", "ab+xy@b.com", "a.b@b.com"}
	fmt.Println(UniqueEmailGroups(emails2)) // Expected: 2

	// Example 3
	emails3 := []string{
		"a.b+c.d+e@DoMain.com",
		"ab+xyz@domain.com",
		"ab@domain.com",
	}
	fmt.Println(UniqueEmailGroups(emails3)) // Expected: 1
}
```
