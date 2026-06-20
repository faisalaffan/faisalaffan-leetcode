# 3436 — Find Valid Emails

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindValidEmails(users []User) []User
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3436: Find Valid Emails
// https://leetcode.com/problems/find-valid-emails/
// Difficulty: Easy

import (
	"fmt"
	"regexp"
)

func main() {
	users := []User{
		{UserID: 1, Email: "alice@example.com"},
		{UserID: 2, Email: "bob@example"},
		{UserID: 3, Email: "@example.com"},
		{UserID: 4, Email: "charlie@example.com"},
	}
	result := FindValidEmails(users)
	for _, r := range result {
		fmt.Printf("%d: %s\n", r.UserID, r.Email)
	}
}

// User represents a user with an email.
type User struct {
	UserID int
	Email  string
}

// FindValidEmails returns users with valid email addresses (alphanumeric prefix, letter-only domain, .com suffix).
// Time: O(n). Space: O(n).
func FindValidEmails(users []User) []User {
	re := regexp.MustCompile(`^[A-Za-z0-9_]+@[A-Za-z]+\.com$`)
	result := []User{}
	for _, u := range users {
		if re.MatchString(u.Email) {
			result = append(result, u)
		}
	}
	return result
}
```
