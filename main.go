// Go punya dua cara utama membuat variabel:

// Menggunakan var

// Menggunakan short declaration :=
package main

func main() {
	// a. menggunakan var
	var x int = 10
	var y string = "halo"
	var z bool = true

	// bisa juga tanpa tipe
	var xy = 20 // otomatis menebak tipenya (int)

	// kalau tanpa nilai sama sekali
	var b int // maka nilainya otomatis jadi zero value
}
