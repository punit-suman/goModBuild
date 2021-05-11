package geometry

import (
	"errors"
	"math"
)

func RectangleArea(a, b float64) (float64, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("enter only positive values")
	}
	return a * b, nil
}

func RectanglePerimeter(a, b float64) (float64, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("enter only positive values")
	}
	return 2 * (a + b), nil
}

func CircleArea(r float64) (float64, error) {
	if r <= 0 {
		return 0, errors.New("enter only positive values for radius")
	}
	return math.Pi * r * r, nil
}

func CircleCircumference(r float64) (float64, error) {
	if r <= 0 {
		return 0, errors.New("enter only positive values for radius")
	}
	return 2 * math.Pi * r, nil
}
