# 0415 — Add Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func AddStrings(num1, num2 string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n+m), Space: O(max(n,m))  
**Kompleksitas Ruang:** O(max(n,m))

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #415: Add Strings
// https://leetcode.com/problems/add-strings/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(max(n,m))
func AddStrings(num1, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	carry := 0
	var result []byte
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		carry = sum / 10
		result = append([]byte{byte(sum%10 + '0')}, result...)
	}
	return string(result)
}

func main() {
	fmt.Println(AddStrings("11", "123"))
	fmt.Println(AddStrings("456", "77"))
	fmt.Println(AddStrings("0", "0"))
}
```
