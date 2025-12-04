// Go punya dua cara utama membuat variabel:

// Menggunakan var

// Menggunakan short declaration :=
package main

import "fmt"

func main() {
	// print variabel
	// pakai fmt.Println():
	name := "budi"
	age := 30

	fmt.Println(name)
	fmt.Println(age)

	// atau dalam satu baris:
	fmt.Println("nama:", name, "umur:", age)
}
