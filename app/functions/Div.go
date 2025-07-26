package functions

import "errors"

type DivScheme struct {
	NumberA float64
	NumberB float64
}

func (input DivScheme) Div(message string) (string, float64, error) {
	if input.NumberA == 0 || input.NumberB == 0 {
		return "", 0, errors.New("Ha ocurrido un error, no puedes dividir entre 0")
	}
	result := input.NumberA / input.NumberB
	return message, result, nil
}
