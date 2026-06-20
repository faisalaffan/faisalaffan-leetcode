# 3692 — Majority Frequency Characters

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MajorityFrequencyCharacters(s string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3692: Majority Frequency Characters
// https://leetcode.com/problems/majority-frequency-characters/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MajorityFrequencyCharacters("aaabbbccdddde"))
	fmt.Println(MajorityFrequencyCharacters("abcd"))
	fmt.Println(MajorityFrequencyCharacters("pfpfgi"))
}

// Time: O(n)
// Space: O(1)
func MajorityFrequencyCharacters(s string) string {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}

  // Alokasi slice integer
	freq := make([]int, len(s)+1)
	for i := 0; i < 26; i++ {
		if cnt[i] > 0 {
			freq[cnt[i]]++
		}
	}

	bestSize := 0
	bestFreq := 0
	for k := 1; k <= len(s); k++ {
		if freq[k] > bestSize || (freq[k] == bestSize && k > bestFreq) {
			bestSize = freq[k]
			bestFreq = k
		}
	}

	res := make([]byte, 0)
	for i := 0; i < 26; i++ {
		if cnt[i] == bestFreq {
			res = append(res, byte('a'+i))
		}
	}
	return string(res)
}
```
