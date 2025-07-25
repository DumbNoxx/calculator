package menu

import (
	"fmt"
	"github.com/DumbNoxx/calculator/app/functions"
	"time"
)

func MenuCalculator() {
	var option int
	for {
		fmt.Println("\tMENU CALCULATOR")
		fmt.Println("1. Sumar")
		fmt.Println("2. Restar")
		fmt.Println("3. Multiplicar")
		fmt.Println("4. Dividir")
		fmt.Println("5. Resto")
		fmt.Println("6. Salir")
		fmt.Print(" R: ")
		fmt.Scan(&option)
		switch option {
		case 1:
			var n1, n2 int
			fmt.Print("Escribe el primero numero: ")
			_, errorOne := fmt.Scan(&n1)
			if errorOne != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			fmt.Print("Escribe el segundo numero: ")
			_, errorTwo := fmt.Scan(&n2)
			if errorTwo != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			sum := functions.Sum_scheme{Number_a: n1, Number_b: n2}
			str, result, error := sum.Sum("El resultado de la suma es:")
			if error != nil {
				fmt.Println(error)
				continue
			}
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("%s %d\n", str, result)
		case 2:
			var n1, n2 int
			fmt.Print("Escribe el primer numero: ")
			_, errorOne := fmt.Scan(&n1)
			if errorOne != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			fmt.Print("Escribe el segundo numero: ")
			_, errorTwo := fmt.Scan(&n2)
			if errorTwo != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			sub := functions.Sub_scheme{Number_a: n1, Number_b: n2}
			str, result, error := sub.Sub("El resultado de la resta es:")
			if error != nil {
				fmt.Println(error)
				continue
			}
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("%s %d\n", str, result)
		case 3:
			var n1, n2 int
			fmt.Print("Escribe el primer numero: ")
			_, errorOne := fmt.Scan(&n1)
			if errorOne != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			fmt.Print("Escribe el segundo numero: ")
			_, errorTwo := fmt.Scan(&n2)
			if errorTwo != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			multi := functions.Multi_Scheme{Number_a: n1, Number_b: n2}
			str, result, error := multi.Multi("El resultado de la multiplicacion es:")
			if error != nil {
				fmt.Println(error)
				continue
			}
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("%s %d\n", str, result)
		case 4:
			var n1, n2 int
			fmt.Print("Escribe el primero numero: ")
			_, errorOne := fmt.Scan(&n1)
			if errorOne != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			fmt.Print("Escribe el segundo numero: ")
			_, errorTwo := fmt.Scan(&n2)
			if errorTwo != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			div := functions.Div_scheme{Number_a: float64(n1), Number_b: float64(n2)}
			str, result, error := div.Div("El resultado de la division es:")
			if error != nil {
				fmt.Println(error)
				continue
			}
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("%s %.2f\n", str, result)
		case 5:
			var n1, n2 int
			fmt.Print("Escribe el primero numero: ")
			_, errorOne := fmt.Scan(&n1)
			if errorOne != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			fmt.Print("Escribe el segundo numero: ")
			_, errorTwo := fmt.Scan(&n2)
			if errorTwo != nil {
				fmt.Println("Tienes que escribir un numero")
				continue
			}
			rest := functions.Rest_scheme{Number_a: float64(n1), Number_b: float64(n2)}
			str, result, error := rest.Rest("El resultado de la division es:")
			if error != nil {
				fmt.Println(error)
				continue
			}
			time.Sleep(400 * time.Millisecond)
			fmt.Printf("%s %d\n", str, result)
		case 6:
			fmt.Println("Saliendo...")
			time.Sleep(600 * time.Millisecond)
			return
		default:
			fmt.Println("Escoge una de las opciones.")
		}
	}
}
