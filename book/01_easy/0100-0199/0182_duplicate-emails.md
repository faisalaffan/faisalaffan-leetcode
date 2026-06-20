# 0182 — Duplicate Emails

## Deskripsi

**Soal:** [0182. Duplicate Emails](https://leetcode.com/problems/duplicate-emails/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func DuplicateEmails() string`

## Solusi Go

```go
package main

// LeetCode #182: Duplicate Emails
// https://leetcode.com/problems/duplicate-emails/
// Difficulty: Easy

import "fmt"

func DuplicateEmails() string {
	return "SELECT email FROM Person GROUP BY email HAVING COUNT(email) > 1"
}

func main() {
	fmt.Println(DuplicateEmails())
}
```
