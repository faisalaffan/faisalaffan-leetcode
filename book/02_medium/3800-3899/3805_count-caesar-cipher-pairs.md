# 3805 — Count Caesar Cipher Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountCaesarCipherPairs(words []string) int
```

> **💡 Hint:** Normalize each string by shifting so that its first character

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N * M)  
**Kompleksitas Ruang:** O(N * M)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3805: Count Caesar Cipher Pairs
// https://leetcode.com/problems/count-caesar-cipher-pairs/
// Difficulty: Medium
// Time: O(N * M) | Space: O(N * M)
// Approach: Normalize each string by shifting so that its first character
// becomes 'a'. Strings in the same Caesar-shift equivalence class will have
// the same normalized form. Count pairs using hash map.

import "fmt"

func CountCaesarCipherPairs(words []string) int {
	normalize := func(s string) string {
		shift := int(s[0] - 'a')
		res := make([]byte, len(s))
  // Loop linear O(n): iterasi setiap elemen
		for i := 0; i < len(s); i++ {
			res[i] = byte((int(s[i]-'a')-shift+26)%26 + 'a')
		}
		return string(res)
	}

  // Membuat map (HashMap) — pencarian O(1)
	count := make(map[string]int)
	ans := 0

	for _, w := range words {
		norm := normalize(w)
		ans += count[norm]
		count[norm]++
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(CountCaesarCipherPairs([]string{"fusion", "layout"})) // Expected: 1

	// Example 2
	fmt.Println(CountCaesarCipherPairs([]string{"ab", "aa", "za", "aa"})) // Expected: 2

	// Example 3
	fmt.Println(CountCaesarCipherPairs([]string{"abc", "bcd", "cde", "xyz"})) // Expected: 3
}
```
