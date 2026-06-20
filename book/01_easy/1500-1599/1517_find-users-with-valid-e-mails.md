# 1517 — Find Users With Valid E Mails

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan string berisi tanda kurung: `()`, `[]`, `{}`. Tugasmu adalah memeriksa apakah string tersebut **valid** — setiap kurung buka harus ditutup oleh kurung yang sesuai dalam urutan benar.

Contoh valid: `()[]{}`, `({[]})`. Tidak valid: `(]`, `([)]`.

**Cara berpikir:** Gunakan Stack. Kurung buka → push. Kurung tutup → pop dan cek kecocokan. Di akhir, stack harus kosong.

**Fungsi Solusi:** `func FindUsersWithValidEMails() string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** N/A (SQL query), Space: N/A  |  **Ruang:** N/A


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
