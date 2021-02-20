package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	expr := regexp.MustCompile("\\b\\w")
	transformadora := func(s string) string {
		return strings.ToLower(s)
	}
	texto := "ANTONIO CARLOS JOBIM"
	fmt.Println(transformadora(texto))
	fmt.Println(expr.ReplaceAllStringFunc(texto, transformadora))
	texto2 := "ARTHUR DAVID MOREIRA RODRIGUES"
	fmt.Println(transformadora(texto2))
	fmt.Println(expr.ReplaceAllStringFunc(texto2, transformadora))
}
