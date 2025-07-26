package functions

import "errors"

type SumScheme struct {
	NumberA int
	NumberB int
}

func (input SumScheme) Sum(message string) (string, int, error) {
	if input.NumberA == 0 && input.NumberB == 0 {
		return "", 0, errors.New("A ocurrido un error, la suma de 0 y 0 da como resultado 0")
	}
	result := input.NumberA + input.NumberB
	return message, result, nil
}
