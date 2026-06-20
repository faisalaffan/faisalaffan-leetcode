# 0900 — Rle Iterator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(encoding []int) RLEIterator
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #900: RLE Iterator
// https://leetcode.com/problems/rle-iterator/
// Difficulty: Medium

import "fmt"

type RLEIterator struct {
	encoding []int
	idx      int
}

func Constructor(encoding []int) RLEIterator {
	return RLEIterator{encoding: encoding, idx: 0}
}

func (this *RLEIterator) Next(n int) int {
	for this.idx < len(this.encoding) {
		if this.encoding[this.idx] >= n {
			this.encoding[this.idx] -= n
			return this.encoding[this.idx+1]
		}
		n -= this.encoding[this.idx]
		this.idx += 2
	}
	return -1
}

func main() {
	// Test: encoding = [3,8,0,9,2,5]
	obj := Constructor([]int{3, 8, 0, 9, 2, 5})
	fmt.Println(obj.Next(2))  // 8
	fmt.Println(obj.Next(1))  // 8
	fmt.Println(obj.Next(1))  // 5
	fmt.Println(obj.Next(2))  // 5

	fmt.Println("---")

	// Test: encoding = [2,1,3,2]
	obj2 := Constructor([]int{2, 1, 3, 2})
	fmt.Println(obj2.Next(3))  // 2
	fmt.Println(obj2.Next(2))  // -1
}
```
