package functions

import "errors"

type Sum_scheme struct {
	Number_a int
	Number_b int
}

func (input Sum_scheme) Sum(message string) (string, int, error) {
	if input.Number_a == 0 && input.Number_b == 0 {
		return "", 0, errors.New("A ocurrido un error, la suma de 0 y 0 da como resultado 0")
	}
	result := input.Number_a + input.Number_b
	return message, result, nil
}
