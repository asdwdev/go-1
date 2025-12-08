package main

import "fmt"

// definisi struct
type Mahasiswa struct {
	Nama string
	Umur int
}

// fungsi yang mengisi data mahasiswa
func isiMahasiswa(m *Mahasiswa) {
	fmt.Println("masukkan data mahasiswa")

	fmt.Print("nama: ")
	fmt.Scan(&m.Nama)

	fmt.Print("umur: ")
	fmt.Scan(&m.Umur)
}

func main() {
	// bikin struct kosong
	var m Mahasiswa

	// oper pointer ke fungsi supaya struct bisa diisi
	isiMahasiswa(&m)

	// tampilkan hasilnya
	fmt.Println("hasil:")
	fmt.Println("nama:", m.Nama)
	fmt.Println("umur:", m.Umur)

}
