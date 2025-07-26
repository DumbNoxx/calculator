package options

import (
	"fmt"
	"github.com/DumbNoxx/calculator/app/menu/options/five"
	"github.com/DumbNoxx/calculator/app/menu/options/four"
	"github.com/DumbNoxx/calculator/app/menu/options/one"
	"github.com/DumbNoxx/calculator/app/menu/options/six"
	"github.com/DumbNoxx/calculator/app/menu/options/three"
	"github.com/DumbNoxx/calculator/app/menu/options/two"
)

type OptionScheme struct {
	NumberOne int
}

func (input OptionScheme) Options() {
	for {
		fmt.Println("\tMENU CALCULATOR")
		fmt.Println("1. Sumar")
		fmt.Println("2. Restar")
		fmt.Println("3. Multiplicar")
		fmt.Println("4. Dividir")
		fmt.Println("5. Resto")
		fmt.Println("6. Salir")
		fmt.Print(" R: ")
		fmt.Scan(&input.NumberOne)
		switch input.NumberOne {
		case 1:
			var n1, n2 int
			OptionOne := one.OneOptionsScheme{OptionOne: n1, OptionTwo: n2}
			OptionOne.One()
		case 2:
			var n1, n2 int
			OptionTwo := two.OptionTwoScheme{OptionOne: n1, OptionTwo: n2}
			OptionTwo.Two()
		case 3:
			var n1, n2 int
			OptionThree := three.ThreeOptionScheme{OptionOne: n1, OptionTwo: n2}
			OptionThree.Three()
		case 4:
			var n1, n2 int
			OptionFour := four.FourOptionScheme{OptionOne: n1, OptionTwo: n2}
			OptionFour.Four()
		case 5:
			var n1, n2 int
			OptionFive := five.FiveOptionScheme{OptionOne: n1, OptionTwo: n2}
			OptionFive.Five()
		case 6:
			six.Six()
			return
		default:
			fmt.Println("Escoge una de las opciones.")
		}
	}
}
