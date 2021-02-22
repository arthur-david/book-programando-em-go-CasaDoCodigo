package main

import (
	"fmt"
	"time"
)

// func imprimir(n int) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("%d ", n)
// 		time.Sleep(200 * time.Millisecond)
// 	}
// }

func dormir() {
	fmt.Println("Goroutine dormindo por 5 segundos...")
	time.Sleep(5 * time.Second)
	fmt.Println("Goroutine finalizada.")
}

func main() {
	go dormir()
	fmt.Println("Main dormindo por 3 segundos...")
	time.Sleep(3 * time.Second)
	fmt.Println("Main finalizada.")
	// go imprimir(2)
	// imprimir(3)
}
