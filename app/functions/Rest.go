package functions

import "errors"

type RestScheme struct {
	NumberA, NumberB float64
}

func (input RestScheme) Rest(message string) (string, int, error) {
	if input.NumberA == 0 || input.NumberB == 0 {
		return "", 0, errors.New("Ha ocurrido un error, no puedes dividir entre 0")
	}
	result := int(input.NumberA) % int(input.NumberB)
	return message, result, nil
}
