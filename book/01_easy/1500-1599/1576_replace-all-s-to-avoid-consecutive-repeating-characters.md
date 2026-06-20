# 1576 — Replace All S To Avoid Consecutive Repeating Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func modifyString(s string) string

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1576: Replace All ?'s to Avoid Consecutive Repeating Characters
// https://leetcode.com/problems/replace-all-s-to-avoid-consecutive-repeating-characters/
// Difficulty: Easy
//
// LeetCode submission: func modifyString(s string) string

import "fmt"

func main() {
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("?zs")) // "azs"
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("ubv?w")) // "ubvaw"
	fmt.Println(ReplaceAllSToAvoidConsecutiveRepeatingCharacters("??yw?ipkj?")) // "abywcipkja"
}

// Time: O(n), Space: O(n)
func ReplaceAllSToAvoidConsecutiveRepeatingCharacters(s string) string {
	res := []byte(s)
	for i, ch := range res {
		if ch == '?' {
			for c := byte('a'); c <= 'z'; c++ {
				if (i == 0 || res[i-1] != c) && (i == len(res)-1 || res[i+1] != c) {
					res[i] = c
					break
				}
			}
		}
	}
	return string(res)
}
```
