# 2423 — Remove Letter To Equalize Frequency

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func RemoveLetterToEqualizeFrequency(word string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2423: Remove Letter To Equalize Frequency
// https://leetcode.com/problems/remove-letter-to-equalize-frequency/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(RemoveLetterToEqualizeFrequency("abcc")) // true
	fmt.Println(RemoveLetterToEqualizeFrequency("aazz")) // false
	fmt.Println(RemoveLetterToEqualizeFrequency("bac"))  // true
}

func RemoveLetterToEqualizeFrequency(word string) bool {
  // Alokasi slice integer
	freq := make([]int, 26)
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
		freq[word[i]-'a']++
	}

	// Try removing one occurrence of each letter
	for i := 0; i < 26; i++ {
		if freq[i] == 0 {
			continue
		}
		freq[i]--
		if allSameFreq(freq) {
			return true
		}
		freq[i]++
	}
	return false
}

func allSameFreq(freq []int) bool {
	target := 0
	for _, f := range freq {
		if f == 0 {
			continue
		}
		if target == 0 {
			target = f
		} else if f != target {
			return false
		}
	}
	return true
}
```
