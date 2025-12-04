// If, For, Switch (Kontrol Alur)
package main

import "fmt"

func main() {
	// 	2️⃣ FOR — Satu-satunya loop di Go

	// Go cuma punya 1 jenis loop yaitu for.
	// Tapi bisa dipakai dalam banyak gaya.

	// 2.2 For sebagai while loop
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}
}
