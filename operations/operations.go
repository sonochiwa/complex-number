package operations

import (
	"fmt"
	"math"
)

// ComplexNumber представляет комплексное число
type ComplexNumber struct {
	Real      float64
	Imaginary float64
}

// NewComplexNumber создает новое комплексное число
func NewComplexNumber(real, imaginary float64) ComplexNumber {
	return ComplexNumber{Real: real, Imaginary: imaginary}
}

// String возвращает строковое представление комплексного числа
func (c ComplexNumber) String() string {
	if c.Imaginary >= 0 {
		return fmt.Sprintf("%.2f + %.2fi", c.Real, c.Imaginary)
	}
	return fmt.Sprintf("%.2f - %.2fi", c.Real, -c.Imaginary)
}

// Add складывает два комплексных числа
func (c ComplexNumber) Add(other ComplexNumber) ComplexNumber {
	return NewComplexNumber(c.Real+other.Real, c.Imaginary+other.Imaginary)
}

// Subtract вычитает два комплексных числа
func (c ComplexNumber) Subtract(other ComplexNumber) ComplexNumber {
	return NewComplexNumber(c.Real-other.Real, c.Imaginary-other.Imaginary)
}

// Multiply умножает два комплексных числа
func (c ComplexNumber) Multiply(other ComplexNumber) ComplexNumber {
	return NewComplexNumber(
		c.Real*other.Real-c.Imaginary*other.Imaginary,
		c.Real*other.Imaginary+c.Imaginary*other.Real,
	)
}

// Divide делит два комплексных числа
func (c ComplexNumber) Divide(other ComplexNumber) ComplexNumber {
	denominator := other.Real*other.Real + other.Imaginary*other.Imaginary
	return NewComplexNumber(
		(c.Real*other.Real+c.Imaginary*other.Imaginary)/denominator,
		(c.Imaginary*other.Real-c.Real*other.Imaginary)/denominator,
	)
}

// Modulus возвращает модуль комплексного числа
func (c ComplexNumber) Modulus() float64 {
	return math.Sqrt(c.Real*c.Real + c.Imaginary*c.Imaginary)
}

// Angle возвращает угол комплексного числа (в радианах)
func (c ComplexNumber) Angle() float64 {
	return math.Atan2(c.Imaginary, c.Real)
}

// ToTrigonometricForm возвращает тригонометрическую форму комплексного числа
func (c ComplexNumber) ToTrigonometricForm() string {
	mag := c.Modulus()
	angle := c.Angle()
	return fmt.Sprintf("%.2f * (cos(%.2f) + i*sin(%.2f))", mag, angle, angle)
}

// ToExponentialForm возвращает экспоненциальную форму комплексного числа
func (c ComplexNumber) ToExponentialForm() string {
	mag := c.Modulus()
	angle := c.Angle()
	return fmt.Sprintf("%.2f * e^(i*%.2f)", mag, angle)
}
