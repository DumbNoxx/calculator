package functions

import "errors"

type MultiScheme struct {
	NumberA int
	NumberB int
}

func (input MultiScheme) Multi(message string) (string, int, error) {
	if input.NumberA == 0 && input.NumberB == 0 {
		return "", 0, errors.New("Ha ocurrido un error, la multiplicacion de 0 y 0 da como resultado 0")
	}
	result := input.NumberA * input.NumberB
	return message, result, nil
}
