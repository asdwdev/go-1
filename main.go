// latihan
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i // kirim angka ke channel
		}
	}()

	for i := 0; i < 5; i++ {
		angka := <-ch // terima dari channel
		fmt.Println(angka)
	}
}
