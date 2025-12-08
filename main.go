// latihan
package main

import (
	"fmt"
	"sync"
)

func cetak(n string) {
	fmt.Println("halo dari", n)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		cetak("a")
		wg.Done()
	}()

	go func() {
		cetak("b")
		wg.Done()
	}()

	go func() {
		cetak("c")
		wg.Done()
	}()

	wg.Wait()
}
