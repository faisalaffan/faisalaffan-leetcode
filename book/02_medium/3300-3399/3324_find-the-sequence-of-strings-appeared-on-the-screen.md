# 3324 — Find The Sequence Of Strings Appeared On The Screen

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func stringSequence(target string) []string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * 26) Space: O(n * 26) for output  
**Kompleksitas Ruang:** O(n * 26) for output

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3324: Find the Sequence of Strings Appeared on the Screen
// https://leetcode.com/problems/find-the-sequence-of-strings-appeared-on-the-screen/
// Difficulty: Medium
// Time: O(n * 26) Space: O(n * 26) for output

import "fmt"

func main() {
	fmt.Println(stringSequence("abc")) // [a aa ab aba abb abc]
	fmt.Println(stringSequence("ab"))  // [a aa ab]
	fmt.Println(stringSequence("z"))   // [a b c d e f g h i j k l m n o p q r s t u v w x y z]
}

func stringSequence(target string) []string {
	var result []string
	var cur []byte

	for _, ch := range target {
		cur = append(cur, 'a')
		result = append(result, string(cur))
		for cur[len(cur)-1] != byte(ch) {
			cur[len(cur)-1]++
			result = append(result, string(cur))
		}
	}

	return result
}
```
