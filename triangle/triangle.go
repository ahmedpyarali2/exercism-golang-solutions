// Package triangle provides functionality for determination of triangle kind.
package triangle

import "math"

// Kind represents the type of the triangle
type Kind int

const (
	// NaT represents Not A Triangle kind - 0
	NaT Kind = iota

	// Equ represents an Equilateral Triangle kind - 1
	Equ

	// Iso represents an Isosceles Triangle kind - 2
	Iso

	// Sca represents an Scalene Triangle kind - 3
	Sca
)

// KindFromSides determines the kind of the triangle based on the lengths of the
// three given sides of the triangle.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case isValidTriangle(a, b, c) == false:
		k = NaT
	case isEquilateralTriangle(a, b, c):
		k = Equ
	case isIsoscelesTriangle(a, b, c):
		k = Iso
	default:
		k = Sca
	}
	return k
}

// isValidTriangle determines if the given lengths of the three sides of a
// triangle form a valid triangle or not.
func isValidTriangle(a, b, c float64) bool {
	// A triangle is considered to be a valid triangle if:
	//	1. all sides must have a length > 0
	//	2. sum of lengths of any two sides must be greater than or equals to the third side
	inf := math.Inf(1)
	return (a > 0 && b > 0 && c > 0) && (a < inf && b < inf && c < inf) && (a+b >= c && b+c >= a && a+c >= b)
}

// isEquilateralTriangle determines if the given triangle sides form an
// equilateral triangle or not.
func isEquilateralTriangle(a, b, c float64) bool {
	// An equilateral triangle has all three sides of the same length
	// so we simply check the equality of all three sides by equivalence property.
	return a == b && b == c
}

// isIsoscelesTriangle determines if the given triangle sides form an
// isosceles triangle or not.
func isIsoscelesTriangle(a, b, c float64) bool {
	// An isosceles triangle has at least two sides of the same length
	// so we check if either two of the sides are same or not.
	// We do this by checking if side a is equals to side b or side b is equals
	// to side c or side a equals to side c.
	return a == b || b == c || a == c
}
