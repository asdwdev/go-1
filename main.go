package main

import "fmt"

type Mahasiswa struct {
	Nama string
	Umur int
}

func main() {
	// 1. Karena ingin ngisi struktur itu belakangan, bukan saat deklarasi
	var m Mahasiswa

	fmt.Print("nama: ")
	fmt.Scan(&m.Nama)

	fmt.Print("umur: ")
	fmt.Scan(&m.Umur)
}
