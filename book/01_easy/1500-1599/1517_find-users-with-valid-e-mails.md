# 1517 — Find Users With Valid E Mails

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan data terstruktur dan diminta untuk mencari elemen atau pola tertentu. Tugasmu adalah menemukan posisi, jumlah, atau keberadaan elemen dengan efisien.

Seperti mencari kata di kamus — kamu tidak membaca dari halaman 1, tapi langsung ke tengah (binary search), lalu maju/mundur. Teknik pencarian yang efisien sangat penting untuk interview.

**Konsep kunci:** linear search O(n), binary search O(log n), HashMap lookup O(1).

**Fungsi yang perlu kamu implementasikan:**
```go
func FindUsersWithValidEMails() string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** N/A (SQL query), Space: N/A  
**Kompleksitas Ruang:** N/A

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1517: Find Users With Valid E-Mails
// https://leetcode.com/problems/find-users-with-valid-e-mails/
// Difficulty: Easy
//
// This is a SQL problem. The solution is the SQL query below.
// Table: Users (user_id, name, mail)

import "fmt"

func main() {
	fmt.Println(FindUsersWithValidEMails())
}

// Time: N/A (SQL query), Space: N/A
func FindUsersWithValidEMails() string {
	return `SELECT user_id, name, mail
FROM Users
WHERE mail REGEXP '^[A-Za-z][A-Za-z0-9_.-]*@leetcode\\.com$';`
}
```
