package two

import (
	"fmt"
	"github.com/DumbNoxx/calculator/app/functions"
	"time"
)

type OptionTwoScheme struct {
	OptionOne, OptionTwo int
}

func (input OptionTwoScheme) Two() {
	functions.ClearScreen()
	fmt.Println("\tRESTAR")
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
	sub := functions.SubScheme{NumberA: input.OptionOne, NumberB: input.OptionTwo}
	str, result, error := sub.Sub("El resultado de la resta es:")
	if error != nil {
		fmt.Println(error)
		fmt.Println("")
		return
	}
	time.Sleep(600 * time.Millisecond)
	fmt.Printf("%s %d\n", str, result)
	fmt.Println("")
}
