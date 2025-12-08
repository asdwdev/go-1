// Struct adalah kumpulan beberapa field (variabel) dalam satu tipe.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

// 4️⃣ Methods (fungsi yang nempel pada struct)
// Seperti fungsi, tapi khusus untuk tipe tertentu.
func (p Person) Greet() {
	fmt.Println("halo, saya", p.Name)
}

func main() {
	p := Person{"andi", 21}
	p.Greet()
}
