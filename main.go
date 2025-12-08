// Channel (Komunikasi Antar Goroutine)
// Channel adalah cara goroutine saling mengirim data.
// Ini yang membuat concurrency di Go beda dari bahasa lain.

// Channel = “pipa” tempat goroutine bisa saling mengirim nilai.
package main

import "fmt"

func main() {
	// 1️⃣ Membuat channel
	ch := make(chan int)

	// 2️⃣ Mengirim data ke channel
	ch <- 10

	// 3️⃣ Menerima data dari channel
	v := <-ch

	fmt.Println(v)

	// 	4️⃣ Sifat channel: blocking

	// Ini penting:

	// Jika goroutine mengirim, dia berhenti sampai ada yang menerima.

	// Jika goroutine menerima, dia berhenti sampai ada yang mengirim.

	// Karena itu, channel otomatis menyinkronkan goroutine.
}
