// Struct adalah kumpulan beberapa field (variabel) dalam satu tipe.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// 2️⃣ Struct literal (cara cepat membuat struct)
	p := Person{"andi", 20}
	fmt.Println(p)

	p.Age = 25
	fmt.Println(p)
}
