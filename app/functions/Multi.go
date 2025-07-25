package functions

import "errors"

type Multi_Scheme struct {
	Number_a int
	Number_b int
}

func (input Multi_Scheme) Multi(message string) (string, int, error) {
	if input.Number_a == 0 && input.Number_b == 0 {
		return "", 0, errors.New("La multiplicacion de 0 y 0 da como resultado 0")
	}
	result := input.Number_a * input.Number_b
	return message, result, nil
}
