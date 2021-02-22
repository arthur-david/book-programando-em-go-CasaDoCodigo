package main

import "fmt"

func main() {
	c := make(chan int)
	go produzir(c)
	go teste(c)

	valor := <-c
	fmt.Println(valor)
}

func produzir(c chan int) {
	c <- 33
}

func teste(c chan int) {
	c <- 22
}
