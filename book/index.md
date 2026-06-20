# LeetCode Go Solutions

Kumpulan solusi **3963 LeetCode problems** diimplementasikan dalam **Go (Golang)**.

## Statistik

| Tingkat | Jumlah | Status |
|---------|--------|--------|
| Easy (Mudah) | 950 | ✅ Complete |
| Medium (Sedang) | 2069 | ✅ Complete |
| Hard (Sulit) | 944 | ✅ Complete |
| **Total** | **3963** | ✅ **All Build Passing** |

## Struktur

```
.
├── 01_easy/    # 950 solusi LeetCode Easy
├── 02_medium/  # 2069 solusi LeetCode Medium
├── 03_hard/    # 944 solusi LeetCode Hard
└── book/       # Dokumentasi (halaman ini)
```

## Pola Solusi

Setiap solusi mengikuti pola:

```go
package main

// LeetCode #X: Problem Name
// https://leetcode.com/problems/problem-name/
// Difficulty: Easy/Medium/Hard

import "fmt"

func main() {
    // Test cases dari LeetCode
    fmt.Println(solutionFunction(input))
}

// Time: O(...) | Space: O(...)
func solutionFunction(params) returnType {
    // Implementasi algoritma
}
```

## Cara Menjalankan

```bash
# Build semua solusi
go build ./...

# Jalankan satu problem
go run ./01_easy/0001_two-sum/

# Jalankan semua problem di satu tingkat
go run ./01_easy/...
```

## Kontributor

- **Muhammad Faisal Affan** — [GitHub](https://github.com/faisalaffan)

---

Dibuat dengan [mdBook](https://rust-lang.github.io/mdBook/). Deploy via GitHub Pages.
