# 1880 — Check If Word Equals Summation Of Two Words

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array integer dan sebuah target. Tugasmu adalah mencari **dua angka** yang jika dijumlahkan menghasilkan target. Kembalikan **indeks** (posisi) kedua angka.

Contoh: `nums=[2,7,11,15], target=9` → `2+7=9` → `[0,1]`.

**Cara berpikir:** Gunakan HashMap. Untuk setiap angka, cek apakah `target-angka` sudah ada di map. Kalau sudah → ketemu pasangan. Kalau belum → simpan angka ke map.

**Fungsi Solusi:** `func IsSumEqual(firstWord string, secondWord string, targetWord string) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #1880: Check if Word Equals Summation of Two Words
// https://leetcode.com/problems/check-if-word-equals-summation-of-two-words/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func IsSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return wordValue(firstWord)+wordValue(secondWord) == wordValue(targetWord)
}

func wordValue(s string) int {
	val := 0
  // Linear scan O(n)
	for i := 0; i < len(s); i++ {
		val = val*10 + int(s[i]-'a')
	}
	return val
}

func main() {
	fmt.Println(IsSumEqual("acb", "cba", "cdb"))
	fmt.Println(IsSumEqual("aaa", "a", "aab"))
	fmt.Println(IsSumEqual("aaa", "a", "aaaa"))
}
```
