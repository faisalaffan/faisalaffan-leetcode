# 3335 — Total Characters In String After Transformations I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func lengthAfterTransformations(s string, t int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + 26t) Space: O(26)  
**Kompleksitas Ruang:** O(26)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3335: Total Characters in String After Transformations I
// https://leetcode.com/problems/total-characters-in-string-after-transformations-i/
// Difficulty: Medium
// Time: O(n + 26t) Space: O(26)

import "fmt"

func main() {
	fmt.Println(lengthAfterTransformations("ab", 1)) // 2
	fmt.Println(lengthAfterTransformations("z", 1))  // 2
	fmt.Println(lengthAfterTransformations("az", 2)) // 5
}

func lengthAfterTransformations(s string, t int) int {
	const mod = 1_000_000_007
	freq := [26]int64{}
	for _, c := range s {
		freq[c-'a']++
	}

	for ; t > 0; t-- {
		next := [26]int64{}
		for i, cnt := range freq {
			if cnt == 0 {
				continue
			}
			if i == 25 { // 'z' -> "ab"
				next[0] = (next[0] + cnt) % mod
				next[1] = (next[1] + cnt) % mod
			} else {
				next[i+1] = (next[i+1] + cnt) % mod
			}
		}
		freq = next
	}

	var ans int64
	for _, cnt := range freq {
		ans = (ans + cnt) % mod
	}
	return int(ans)
}
```
