package three

import (
	"fmt"
	"github.com/DumbNoxx/calculator/app/functions"
	"time"
)

type ThreeOptionScheme struct {
	OptionOne, OptionTwo int
}

func (input ThreeOptionScheme) Three() {
	functions.ClearScreen()
	fmt.Println("\tMULTIPLICAR")
	fmt.Println("")
	fmt.Print("Escribe el primer numero: ")
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
	multi := functions.MultiScheme{NumberA: input.OptionOne, NumberB: input.OptionTwo}
	str, result, error := multi.Multi("El resultado de la multiplicacion es:")
	if error != nil {
		fmt.Println(error)
		fmt.Println("")
		return
	}
	time.Sleep(600 * time.Millisecond)
	fmt.Printf("%s %d\n\n", str, result)
}
