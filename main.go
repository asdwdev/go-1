// Buffered Channel & Close()
// (dasar penting untuk produksi)

// 1️⃣ Unbuffered vs Buffered Channel
// Unbuffered
// ch := make(chan int)

// Setiap pengiriman harus ditunggu sampai ada penerima.

// Ini membuat sinkronisasi otomatis.

// Buffered
// ch := make(chan int, 3)

// Channel bisa menampung 3 data tanpa harus langsung diterima.

// Pengiriman hanya berhenti kalau buffer penuh.

// Ilustrasi:

// • Isi buffer: [ _ _ _ ] → kosong
// ch <- 10 → tidak block
// ch <- 20 → tidak block
// ch <- 30 → tidak block
// ch <- 40 → block (buffer penuh)

// 2️⃣ Menutup channel: close(ch)

// Kenapa perlu?

// Supaya penerima tahu tidak ada data lagi.

// Kalau channel ditutup, range channel akan berhenti.

// Contoh:

// close(ch)

// Setelah ditutup, mengirim data → PANIC
// Menerima data → aman (akan mendapat nilai zero sampai habis)

// ✔️ Contoh Idiomatik
package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	go func() {
		ch <- "pesan 1"
		ch <- "pesan 2"
		ch <- "pesan 3"
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
