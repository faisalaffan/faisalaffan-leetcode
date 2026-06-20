# 3114 — Latest Time You Can Obtain After Replacing Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func LatestTimeYouCanObtainAfterReplacingCharacters(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3114: Latest Time You Can Obtain After Replacing Characters
// https://leetcode.com/problems/latest-time-you-can-obtain-after-replacing-characters/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: findLatestTime
	fmt.Println(LatestTimeYouCanObtainAfterReplacingCharacters("1?:?4")) // 11:54
	fmt.Println(LatestTimeYouCanObtainAfterReplacingCharacters("0?:5?")) // 09:59
}

// Time: O(1) | Space: O(1)
// LeetCode submission name: findLatestTime
func LatestTimeYouCanObtainAfterReplacingCharacters(s string) string {
	time := []byte(s)

	// Replace hours
	if time[0] == '?' {
		if time[1] == '?' || time[1] <= '1' {
			time[0] = '1'
		} else {
			time[0] = '0'
		}
	}
	if time[1] == '?' {
		if time[0] == '1' {
			time[1] = '1'
		} else {
			time[1] = '9'
		}
	}

	// Replace minutes
	if time[3] == '?' {
		time[3] = '5'
	}
	if time[4] == '?' {
		time[4] = '9'
	}

	return string(time)
}
```
