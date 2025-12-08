// latihan
package main

import (
	"fmt"
	"time"
)

func cetak(n string) {
	fmt.Println("halo dari", n)
}

func main() {
	go cetak("goroutine 1")
	go cetak("goroutine 2")
	go cetak("goroutine 3")

	time.Sleep(1 * time.Second)
}
