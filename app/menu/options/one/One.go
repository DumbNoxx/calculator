package one

import (
	"fmt"
	"github.com/DumbNoxx/calculator/app/functions"
	"time"
)

type OneOptionsScheme struct {
	OptionOne, OptionTwo int
}

func (input OneOptionsScheme) One() {
	functions.ClearScreen()
	fmt.Println("\tSUMAR")
	fmt.Println("")
	fmt.Print("Escribe el primero numero: ")
	_, errorOne := fmt.Scan(&input.OptionOne)
	if errorOne != nil {
		fmt.Println("Tienes que escribir un numero")
	}
	fmt.Print("Escribe el segundo numero: ")
	_, errorTwo := fmt.Scan(&input.OptionTwo)
	fmt.Println("")
	if errorTwo != nil {
		fmt.Println("Tienes que escribir un numero")
	}
	sum := functions.SumScheme{NumberA: input.OptionOne, NumberB: input.OptionTwo}
	str, result, error := sum.Sum("El resultado de la suma es:")
	if error != nil {
		fmt.Println(error)
		fmt.Println("")
		return
	}
	time.Sleep(600 * time.Millisecond)
	fmt.Printf("%s %d\n\n", str, result)
}
