package functions

import "errors"

type Div_scheme struct {
	Number_a float64
	Number_b float64
}

func (input Div_scheme) Div(message string) (string, float64, error) {
	if input.Number_a == 0 || input.Number_b == 0 {
		return "", 0, errors.New("No puedes dividir entre 0")
	}
	result := input.Number_a / input.Number_b
	return message, result, nil
}
