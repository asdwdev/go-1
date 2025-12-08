// Struct adalah kumpulan beberapa field (variabel) dalam satu tipe.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Println("halo, saya", p.Name)
}

// 5️⃣ Pointer receiver (mengubah nilai struct)
// Kalau kamu mau method bisa mengubah data struct, buat begini:
func (p *Person) Birthday() {
	p.Age++
}

func main() {
	p := Person{"andi", 21}
	p.Birthday()
	fmt.Println(p)
}
