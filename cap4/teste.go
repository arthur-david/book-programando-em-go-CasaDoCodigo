package main

import (
	"fmt"
)

func main() {
	a := map[int]string{}
	a[1] = "abc"
	a[2] = "def"
	a[3] = "ghi"

	fmt.Println(len(a))
	for i, v := range a {
		fmt.Println(i, "=", v)
	}
}
