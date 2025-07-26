package functions

import "errors"

type SubScheme struct {
	NumberA int
	NumberB int
}

func (input SubScheme) Sub(message string) (string, int, error) {
	if input.NumberA == 0 && input.NumberB == 0 {
		return "", 0, errors.New("Ha ocurrido un error, la resta de 0 y 0 da como resultado 0")
	}
	result := input.NumberA - input.NumberB
	return message, result, nil
}
