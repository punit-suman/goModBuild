package geometry

import "errors"

func RectangleArea(a, b float64) (float64, error) {
	if a <= 0 || b <= 0 {
		return 0, errors.New("enter only positive values")
	}
	return a * b, nil
}
