# 1106 — Parsing A Boolean Expression

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func parseBoolExpr(expression string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1106: Parsing A Boolean Expression
// https://leetcode.com/problems/parsing-a-boolean-expression/
// Difficulty: Hard
//
// Recursive descent parser for boolean expressions:
//   't' → true
//   'f' → false
//   '!(expr)' → NOT
//   '&(expr,expr,...)' → AND
//   '|(expr,expr,...)' → OR

import "fmt"

func main() {
	fmt.Println(parseBoolExpr("&(|(f))"))
	fmt.Println(parseBoolExpr("|(f,f,f,t)"))
}

func parseBoolExpr(expression string) bool {
	idx := 0

	var parse func() bool
	parse = func() bool {
		ch := expression[idx]
		idx++

		if ch == 't' {
			return true
		}
		if ch == 'f' {
			return false
		}

		// ch is '!', '&', or '|'
		idx++ // skip '('

		var result bool
		if ch == '!' {
			result = !parse()
		} else {
			result = parse()
			for idx < len(expression) && expression[idx] == ',' {
				idx++ // skip ','
				val := parse()
				if ch == '&' {
					result = result && val
				} else { // '|'
					result = result || val
				}
			}
		}

		idx++ // skip ')'
		return result
	}

	return parse()
}
```
