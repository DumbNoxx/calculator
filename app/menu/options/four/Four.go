package four

import (
	"fmt"
	"github.com/DumbNoxx/calculator/app/functions"
	"time"
)

type FourOptionScheme struct {
	OptionOne, OptionTwo int
}

func (input FourOptionScheme) Four() {
	functions.ClearScreen()
	fmt.Println("\tDIVIDIR")
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
	div := functions.DivScheme{NumberA: float64(input.OptionOne), NumberB: float64(input.OptionTwo)}
	str, result, error := div.Div("El resultado de la division es:")
	if error != nil {
		fmt.Println(error)
		fmt.Println("")
		return
	}
	time.Sleep(600 * time.Millisecond)
	fmt.Printf("%s %.2f\n\n", str, result)
}
