// latihan bagian 4
package main

import "fmt"

func main() {
	var x int
	fmt.Print("masukkan angka: ")
	fmt.Scanln(&x)

	if x%2 == 0 {
		fmt.Println("angka genap")
	} else {
		fmt.Println("angka ganjil")
	}

	i := 1
	for i <= x {
		fmt.Println(i)
		i++
	}

	switch {
	case x > 10:
		fmt.Println("angka besar")
	}
}
