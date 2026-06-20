# 0838 — Push Dominoes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func PushDominoes(dominoes string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #838: Push Dominoes
// https://leetcode.com/problems/push-dominoes/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(PushDominoes("RR.L"))
	fmt.Println(PushDominoes(".L.R...LR..L.."))
	fmt.Println(PushDominoes("L.R"))
}

// Time: O(n) | Space: O(n)
func PushDominoes(dominoes string) string {
	n := len(dominoes)
	res := []byte(dominoes)

	for i := 0; i < n; i++ {
		if res[i] == 'R' {
			// Find the next non-dot character
			j := i + 1
			for j < n && res[j] == '.' {
				j++
			}
			if j == n || res[j] == 'R' {
				// All dots between i and j become 'R'
				for k := i + 1; k < j; k++ {
					res[k] = 'R'
				}
			} else if res[j] == 'L' {
				// Collision: left and right meet in the middle
				left, right := i+1, j-1
  // Two-pointer: gerakkan kiri atau kanan
				for left < right {
					res[left] = 'R'
					res[right] = 'L'
					left++
					right--
				}
			}
			i = j
		} else if res[i] == 'L' {
			// Propagate L leftwards
			j := i - 1
			for j >= 0 && res[j] == '.' {
				res[j] = 'L'
				j--
			}
		}
	}

	// Handle initial dots before first 'L'
	for i := 0; i < n && res[i] == '.'; i++ {
		if i+1 < n && res[i+1] == 'L' {
			res[i] = 'L'
		}
	}

	return string(res)
}
```
