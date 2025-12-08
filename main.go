package main

import "fmt"

type Mahasiswa struct {
	Nama    string
	Umur    int
	Jurusan string
}

func (m Mahasiswa) Perkenalan() {
	fmt.Println("halo, nama saya", m.Nama, "umur", m.Umur, "jurusan", m.Jurusan)
}

func (m *Mahasiswa) UlangTahun() {
	m.Umur++
}

func main() {
	m := Mahasiswa{"nil", 21, "informatika"}
	m.Perkenalan()
	m.UlangTahun()
	m.Perkenalan()
}
