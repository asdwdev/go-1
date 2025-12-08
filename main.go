// sync.WaitGroup (Cara Benar Menunggu Goroutine)

// time.Sleep cuma cocok buat contoh sederhana.
// Di dunia nyata, kita pakai WaitGroup.

// WaitGroup = alat untuk:

// Menambah jumlah goroutine yang perlu ditunggu

// Menandai goroutine selesai

// Menunggu semuanya selesai
package main

import (
	"fmt"
	"sync"
)

func cetak(n string) {
	fmt.Println("halo dari", n)
}

func main() {
	var wg sync.WaitGroup

	// 3️⃣ Tambah jumlah goroutine
	wg.Add(3)

	go func() {
		cetak("goroutine 1")
		wg.Done()
	}()

	go func() {
		cetak("goroutine 2")
		wg.Done()
	}()

	go func() {
		cetak("goroutine 3")
		wg.Done()
	}()

	wg.Wait()
}
