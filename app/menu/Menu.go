package menu

import "github.com/DumbNoxx/calculator/app/menu/options"

func MenuCalculator() {
	var optionMenu int
	menu := options.OptionScheme{NumberOne: optionMenu}
	menu.Options()
}
