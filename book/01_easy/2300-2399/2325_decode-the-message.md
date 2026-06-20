# 2325 — Decode The Message

## Deskripsi

**Soal:** [2325. Decode The Message](https://leetcode.com/problems/decode-the-message/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2325: Decode the Message
// https://leetcode.com/problems/decode-the-message/
// Difficulty: Easy
// Time O(n + m) | Space O(1)

import "fmt"

func main() {
	fmt.Println(DecodeTheMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv")) // "this is a secret"
	fmt.Println(DecodeTheMessage("eljuxhpwnyrdgtqkviszcfmabo", "zwx hnfx lqantp mnoeius ycgk vcnjrdb")) // "the five boxing wizards jump quickly"
}

func DecodeTheMessage(key string, message string) string {
  // Membuat slice untuk menyimpan hasil
	mapping := make([]byte, 26)
	idx := byte(0)

  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(key); i++ {
		if key[i] != ' ' && mapping[key[i]-'a'] == 0 {
			mapping[key[i]-'a'] = 'a' + idx
			idx++
		}
	}

  // Membuat slice untuk menyimpan hasil
	res := make([]byte, len(message))
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(message); i++ {
		if message[i] == ' ' {
			res[i] = ' '
		} else {
			res[i] = mapping[message[i]-'a']
		}
	}
	return string(res)
}
```
