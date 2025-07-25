package functions

import "errors"

type Rest_scheme struct {
	Number_a float64
	Number_b float64
}

func (input Rest_scheme) Rest(message string) (string, int, error) {
	if input.Number_a == 0 || input.Number_b == 0 {
		return "", 0, errors.New("No puedes dividir por 0")
	}
	result := int(int(input.Number_a) % int(input.Number_b))
	return message, result, nil
}
