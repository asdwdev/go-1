package main

import "fmt"

type Mahasiswa struct {
	Nama string
	Umur int
}

func main() {
	var m Mahasiswa

	if nama := "ajit"; nama == "ijat" {
		m.Nama = "ijat"
	} else {
		m.Nama = "warko"
	}

	fmt.Println(m.Nama)
}
