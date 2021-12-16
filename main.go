package explorepackage

import "errors"

func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func mult(a, b int) int {
	return a * b
}
func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero not valid")

	}
	return a / b, nil

}