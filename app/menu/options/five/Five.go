package five

import (
	"fmt"
	"github.com/DumbNoxx/calculator/app/functions"
	"time"
)

type FiveOptionScheme struct {
	OptionOne, OptionTwo int
}

func (input FiveOptionScheme) Five() {
	functions.ClearScreen()
	fmt.Println("\tRESTO")
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
	rest := functions.RestScheme{NumberA: float64(input.OptionOne), NumberB: float64(input.OptionTwo)}
	str, result, error := rest.Rest("El resultado de la division es:")
	if error != nil {
		fmt.Println(error)
		fmt.Println("")
		return
	}
	time.Sleep(400 * time.Millisecond)
	fmt.Printf("%s %d\n", str, result)
	fmt.Println("")
}
