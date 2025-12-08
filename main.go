package main

import "fmt"

type Kendaraan interface {
	Info() string
}

type Mobil struct {
	Merk  string
	Tahun int
}

type Motor struct {
	Merk  string
	Tahun int
}

func (mobil Mobil) Info() string {
	return fmt.Sprintf("mobil %s (%d)", mobil.Merk, mobil.Tahun)
}

func (motor Motor) Info() string {
	return fmt.Sprintf("motor %s (%d)", motor.Merk, motor.Tahun)
}

func CetakInfo(k Kendaraan) {
	fmt.Println(k.Info())
}

func main() {
	mobil := Mobil{"honda", 2010}
	motor := Motor{"yamaha", 2015}
	CetakInfo(mobil)
	CetakInfo(motor)
}
