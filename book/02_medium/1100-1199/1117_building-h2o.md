# 1117 — Building H2O

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func NewH2O() *H2O
```

> **💡 Hint:** Use mutex and condition variable to synchronize H and O threads.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n) where n = number of water molecules  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1117: Building H2O
// https://leetcode.com/problems/building-h2o/
// Difficulty: Medium
//
// Approach: Use mutex and condition variable to synchronize H and O threads.
//           Barrier ensures exactly 2 H + 1 O per molecule.
// Time: O(n) where n = number of water molecules
// Space: O(1)

import (
	"fmt"
	"sync"
)

type H2O struct {
	mu      sync.Mutex
	hCount  int
	oCount  int
	hNeed   int
	oNeed   int
	cond    *sync.Cond
}

func NewH2O() *H2O {
	h2o := &H2O{hNeed: 2, oNeed: 1}
	h2o.cond = sync.NewCond(&h2o.mu)
	return h2o
}

func (h2o *H2O) Hydrogen(releaseHydrogen func()) {
	h2o.mu.Lock()
	for h2o.hCount >= h2o.hNeed {
		h2o.cond.Wait()
	}
	h2o.hCount++
	releaseHydrogen()
	if h2o.hCount == h2o.hNeed && h2o.oCount == h2o.oNeed {
		h2o.hCount = 0
		h2o.oCount = 0
	}
	h2o.cond.Broadcast()
	h2o.mu.Unlock()
}

func (h2o *H2O) Oxygen(releaseOxygen func()) {
	h2o.mu.Lock()
	for h2o.oCount >= h2o.oNeed {
		h2o.cond.Wait()
	}
	h2o.oCount++
	releaseOxygen()
	if h2o.hCount == h2o.hNeed && h2o.oCount == h2o.oNeed {
		h2o.hCount = 0
		h2o.oCount = 0
	}
	h2o.cond.Broadcast()
	h2o.mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	h2o := NewH2O()
	output := make([]byte, 0, 12)

	// Form 2 water molecules: need 4 H and 2 O
	atoms := []byte{'H', 'H', 'O', 'H', 'H', 'O'}
	var mu sync.Mutex

	for _, atom := range atoms {
		wg.Add(1)
		go func(a byte) {
			defer wg.Done()
			if a == 'H' {
				h2o.Hydrogen(func() {
					mu.Lock()
					output = append(output, 'H')
					mu.Unlock()
				})
			} else {
				h2o.Oxygen(func() {
					mu.Lock()
					output = append(output, 'O')
					mu.Unlock()
				})
			}
		}(atom)
	}

	wg.Wait()
	fmt.Println(string(output))
}
```
