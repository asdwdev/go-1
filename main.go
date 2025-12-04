// fungsi dengan return value (mengembalikan nilai)
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	hasil := add(2, 3)
	fmt.Println(hasil)
}
