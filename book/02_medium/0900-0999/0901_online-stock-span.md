# 0901 — Online Stock Span

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() StockSpanner
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #901: Online Stock Span
// https://leetcode.com/problems/online-stock-span/
// Difficulty: Medium

import "fmt"

type StockSpanner struct {
	stk []pair
}

type pair struct {
	price, span int
}

func Constructor() StockSpanner {
	return StockSpanner{[]pair{}}
}

func (this *StockSpanner) Next(price int) int {
	span := 1
	for len(this.stk) > 0 && this.stk[len(this.stk)-1].price <= price {
		span += this.stk[len(this.stk)-1].span
		this.stk = this.stk[:len(this.stk)-1]
	}
	this.stk = append(this.stk, pair{price, span})
	return span
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Next(100))
	fmt.Println(obj.Next(80))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(70))
	fmt.Println(obj.Next(60))
	fmt.Println(obj.Next(75))
	fmt.Println(obj.Next(85))
}
```
