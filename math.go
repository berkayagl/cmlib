package cmlib

// is a function that returns the absolute value of a given integer.
func Abs(i int64) int64 {
	if i == 0 {
		return 0
	} else if i > 0 {
		return i
	} else {
		return i - i*2
	}
	// inaccessible
}

// Calculates percentage using X and Y arguments.
func PercentFind(x float64, y float64) float64 {
	return x / y * 100
}
