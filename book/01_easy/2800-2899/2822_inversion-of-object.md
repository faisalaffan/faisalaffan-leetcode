# 2822 — Inversion Of Object

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func InversionOfObject(obj map[string]int) map[int]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2822: Inversion of Object
// https://leetcode.com/problems/inversion-of-object/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(n)
// Note: JS problem, adapted to Go. Swaps keys and values of a map.

import "fmt"

func main() {
	input := map[string]int{"a": 1, "b": 2, "c": 1}
	fmt.Println(InversionOfObject(input))
}

func InversionOfObject(obj map[string]int) map[int]string {
  // Membuat map (HashMap) — pencarian O(1)
	result := make(map[int]string, len(obj))
	for k, v := range obj {
		result[v] = k
	}
	return result
}
```
