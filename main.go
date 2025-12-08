package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "halo dari goroutine"
	}()

	pesan := <-ch
	fmt.Println(pesan)
}
