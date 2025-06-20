package calculator

import (
	"errors"
	"strconv"
)

// ErrDivisionByZero is returned when attempting to divide by zero
var ErrDivisionByZero = errors.New("division by zero")

// Add adds two float64 numbers
func Add(a, b float64) float64 {
<<<<<<< HEAD
	// TODO: Implement this function
	return 0
=======
	return a + b
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}

// Subtract subtracts b from a
func Subtract(a, b float64) float64 {
<<<<<<< HEAD
	// TODO: Implement this function
	return 0
=======
	return a - b
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}

// Multiply multiplies two float64 numbers
func Multiply(a, b float64) float64 {
<<<<<<< HEAD
	// TODO: Implement this function
	return 0
=======
	return a * b
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}

// Divide divides a by b, returns an error if b is zero
func Divide(a, b float64) (float64, error) {
<<<<<<< HEAD
	// TODO: Implement this function
	return 0, nil
=======
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}

// StringToFloat converts a string to float64
func StringToFloat(s string) (float64, error) {
<<<<<<< HEAD
	// TODO: Implement this function
	return 0, nil
=======
    f, err := strconv.ParseFloat(s, 64)
    if err != nil {
        return 0, err  
    }
    return f, nil
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}

// FloatToString converts a float64 to string with specified precision
func FloatToString(f float64, precision int) string {
<<<<<<< HEAD
	// TODO: Implement this function
	return ""
=======
	return strconv.FormatFloat(f, 'f', precision, 64)
>>>>>>> a0b7266 (lab01: реализованы базовые задачи Go и Flutter)
}
