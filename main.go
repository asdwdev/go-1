// Go punya dua cara utama membuat variabel:

// Menggunakan var

// Menggunakan short declaration :=
package main

import "fmt"

func main() {
	// multiple variabel declaration

	// contoh:
	var a, b, c int = 1, 2, 3
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// atau short:
	x, y := 10, 20
	fmt.Println(x)
	fmt.Println(y)

	// bisa beda tipe mamakai grup
	var (
		name string = "andi"
		age  int    = 25
		vip  bool   = true
	)
	fmt.Println(name, age, vip)
}
