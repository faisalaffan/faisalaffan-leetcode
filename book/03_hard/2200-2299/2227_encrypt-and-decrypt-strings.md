# 2227 — Encrypt And Decrypt Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(keys []byte, values []string, dictionary []string) Encrypter
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2227: Encrypt and Decrypt Strings
// https://leetcode.com/problems/encrypt-and-decrypt-strings/
// Difficulty: Hard
//
// You are given a character array keys and a string array values (same length),
// and a dictionary of allowed words (string array).
// The Encrypter class:
//   - Encrypt(s): replaces each character c with values[i] where keys[i]==c.
//   - Decrypt(s): returns the number of strings in the dictionary that could
//     result in s after encryption.

import (
	"fmt"
)

// Encrypter encrypts and decrypts strings.
type Encrypter struct {
	charToVal map[byte]string
	valToChar map[string]byte
	encrypted map[string]int // encrypted word -> frequency in dictionary
}

// Constructor initializes the Encrypter.
func Constructor(keys []byte, values []string, dictionary []string) Encrypter {
  // Membuat map (HashMap) — pencarian O(1)
	charToVal := make(map[byte]string)
  // Membuat map (HashMap) — pencarian O(1)
	valToChar := make(map[string]byte)

	for i, k := range keys {
		charToVal[k] = values[i]
		// allow multiple chars to map to same value, but we keep first
		if _, ok := valToChar[values[i]]; !ok {
			valToChar[values[i]] = k
		}
	}

  // Membuat map (HashMap) — pencarian O(1)
	encrypted := make(map[string]int)
	for _, word := range dictionary {
		encWord := encryptWord(word, charToVal)
		if encWord != "" {
			encrypted[encWord]++
		}
	}

	return Encrypter{
		charToVal: charToVal,
		valToChar: valToChar,
		encrypted: encrypted,
	}
}

// encryptWord encrypts a word using the char-to-value mapping.
// Returns empty string if the word contains a char not in keys.
func encryptWord(word string, charToVal map[byte]string) string {
	var result []byte
  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(word); i++ {
		val, ok := charToVal[word[i]]
		if !ok {
			return "" // can't encrypt
		}
		result = append(result, []byte(val)...)
	}
	return string(result)
}

// Encrypt encrypts the input string s.
func (e *Encrypter) Encrypt(s string) string {
	result := encryptWord(s, e.charToVal)
	return result
}

// Decrypt returns the number of dictionary words that could
// have produced the encrypted string.
func (e *Encrypter) Decrypt(s string) int {
	return e.encrypted[s]
}

func main() {
	keys := []byte{'a', 'b', 'c', 'd'}
	values := []string{"ei", "zf", "ei", "am"}
	dictionary := []string{"abcd", "acbd", "adbc", "badc", "dabc", "cabd"}

	encrypter := Constructor(keys, values, dictionary)

	// Test encrypt
	encrypted := encrypter.Encrypt("abcd")
	fmt.Println("Encrypted 'abcd':", encrypted) // "eizfeiam"

	// Test decrypt
	count := encrypter.Decrypt(encrypted)
	fmt.Println("Decrypt count:", count) // 2 ("abcd", "acbd" both map to same)
}
```
