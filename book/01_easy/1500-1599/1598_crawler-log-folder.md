# 1598 — Crawler Log Folder

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(logs []string) int

import "fmt"

func main()
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1598: Crawler Log Folder
// https://leetcode.com/problems/crawler-log-folder/
// Difficulty: Easy
//
// LeetCode submission: func minOperations(logs []string) int

import "fmt"

func main() {
	fmt.Println(CrawlerLogFolder([]string{"d1/", "d2/", "../", "d21/", "./"}))            // 2
	fmt.Println(CrawlerLogFolder([]string{"d1/", "../", "../", "../"}))                   // 0
	fmt.Println(CrawlerLogFolder([]string{"./", "../", "./"}))                            // 0
}

// Time: O(n), Space: O(1)
func CrawlerLogFolder(logs []string) int {
	depth := 0
	for _, log := range logs {
		switch log {
		case "../":
			if depth > 0 {
				depth--
			}
		case "./":
			// do nothing
		default:
			depth++
		}
	}
	return depth
}
```
