# 0359 — Logger Rate Limiter

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() Logger
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #359: Logger Rate Limiter
// https://leetcode.com/problems/logger-rate-limiter/
// Difficulty: Easy [Paid]

import "fmt"

// Logger tracks messages and their last printed timestamp.
type Logger struct {
	lastPrinted map[string]int
}

// Constructor creates a new Logger.
func Constructor() Logger {
	return Logger{lastPrinted: make(map[string]int)}
}

// ShouldPrintMessage returns true if the message should be printed at the given timestamp.
// A message can be printed if it hasn't been printed in the last 10 seconds.
// Time: O(1), Space: O(n)
func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	if lastTs, ok := l.lastPrinted[message]; ok && timestamp-lastTs < 10 {
		return false
	}
	l.lastPrinted[message] = timestamp
	return true
}

func main() {
	obj := Constructor()
	fmt.Println(obj.ShouldPrintMessage(1, "foo"))
	fmt.Println(obj.ShouldPrintMessage(2, "bar"))
	fmt.Println(obj.ShouldPrintMessage(3, "foo"))
	fmt.Println(obj.ShouldPrintMessage(8, "bar"))
	fmt.Println(obj.ShouldPrintMessage(10, "foo"))
	fmt.Println(obj.ShouldPrintMessage(11, "foo"))
}
```
