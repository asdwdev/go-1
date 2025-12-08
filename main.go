// Struct adalah kumpulan beberapa field (variabel) dalam satu tipe.
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// membuat object-nya:
	p := Person{
		Name: "budi",
		Age:  21,
	}
	// akses field:
	fmt.Println(p.Name)
	fmt.Println(p.Age)
}
