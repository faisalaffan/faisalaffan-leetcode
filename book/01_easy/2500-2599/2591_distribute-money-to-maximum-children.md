# 2591 — Distribute Money To Maximum Children

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func DistributeMoneyToMaximumChildren(money int, children int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2591: Distribute Money to Maximum Children
// https://leetcode.com/problems/distribute-money-to-maximum-children/
// Difficulty: Easy
// Time O(1) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DistributeMoneyToMaximumChildren(20, 3)) // 1
	fmt.Println(DistributeMoneyToMaximumChildren(16, 2)) // 2
}

func DistributeMoneyToMaximumChildren(money int, children int) int {
	if money < children {
		return -1
	}

	// Give each child 1 dollar first
	money -= children

	// Now we have 7-dollar increments (to make 8) for as many children as possible
	count := money / 7
	money %= 7

	// If count > children, we over-assigned
	if count > children {
		return children - 1
	}

	// If we have 3 children left and money = 3, we can't give it optimally
	remaining := children - count
	if remaining == 0 && money > 0 {
		count--
	} else if remaining == 1 && money == 3 {
		count--
	}

	return count
}
```
