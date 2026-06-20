# 1241 — Number Of Comments Per Post

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan bilangan bulat dan diminta untuk menghitung, memanipulasi, atau menganalisis properti bilangan tersebut.

Soal tipe bilangan menguji pemahamanmu tentang operasi matematika, digit, atau properti bilangan (prima, palindrome, pembagi, dll). Kuncinya adalah menemukan pola matematika sebelum menulis kode.

**Konsep kunci:** modulo (%), pembagian integer, digit extraction, prime check, GCD/LCM.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1241: Number of Comments per Post
// https://leetcode.com/problems/number-of-comments-per-post/
// Difficulty: Easy [Paid] (SQL)
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println("SQL: SELECT DISTINCT p.sub_id AS post_id, (SELECT COUNT(DISTINCT c.sub_id) FROM Submissions c WHERE c.parent_id = p.sub_id) AS number_of_comments FROM Submissions p WHERE p.parent_id IS NULL ORDER BY p.sub_id")
}

// This is a SQL problem. The answer is the SQL query above.
```
