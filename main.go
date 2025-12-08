// 1️⃣ Apa itu Goroutine?

// Goroutine adalah cara Go menjalankan fungsi secara concurrent (bersamaan), tapi lebih ringan dari thread.

package main

import (
	"fmt"
	"time"
)

func main() {
	// go func() {
	// 	fmt.Println("halo dari goroutine")
	// }()

	go fmt.Println("halo")

	// 	Sering tidak muncul apa-apa.

	// Kenapa?
	// Karena main selesai duluan, goroutine belum sempat jalan.

	// Makanya kita butuh synchronization.

	// 3️⃣ Cara sinkronisasi: time.Sleep (cara paling basic)
	time.Sleep(1 * time.Second)

	// 	Ini hanya untuk belajar dasar.
	// Nanti kita pakai WaitGroup yang benar.
}
