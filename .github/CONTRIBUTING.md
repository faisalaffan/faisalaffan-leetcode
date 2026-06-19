# Contributing

Terima kasih sudah berkontribusi! Repo ini adalah kumpulan solusi LeetCode dalam Go.

## Cara Kontribusi

### Menambah Solusi

1. Pilih soal yang belum ada solusinya
2. Cari file stub di `easy/`, `medium/`, atau `hard/`
3. Implementasi function yang sudah disediakan
4. Pastikan kode bisa kompilasi: `go build ./...`
5. Buka PR

### Konvensi

- **File**: `{4digit_id}_{slug}.go` — jangan rename
- **Function**: PascalCase dari slug, sudah disediakan di stub
- **Solusi**: se-simpel dan se-efisien mungkin
- **Comment**: tambahkan kompleksitas waktu/ruang di atas function

### Format PR

```markdown
## Problem
- LeetCode #123: Problem Title
- Difficulty: Easy/Medium/Hard

## Solution
- Approach: (two-pointer, DP, BFS, dll)
- Time: O(n)
- Space: O(1)
```
