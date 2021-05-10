package geometry

import (
	"errors"
)

func RectangleArea(a, b float64) (float64, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("enter only positive values")
	}
	return a * b, nil
}

func CircleArea(r float64) (float64, error) {
	if r <= 0 {
		return 0, errors.New("enter only positive values for radius")
	}
	return 3.1416 * r * r, nil
}
