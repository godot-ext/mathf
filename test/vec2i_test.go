package test

import (
	"godot-ext/mathf"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVector2i(t *testing.T) {
	// Test construction
	v1 := mathf.NewVector2i(1, 2)
	v2 := mathf.NewVector2i(3, 4)

	// Test addition
	sum := v1.Add(v2)
	assert.Equal(t, mathf.NewVector2i(4, 6), sum, "Addition failed")

	// Test subtraction
	diff := v2.Sub(v1)
	assert.Equal(t, mathf.NewVector2i(2, 2), diff, "Subtraction failed")

	// Test multiplication
	prod := v1.Mul(v2)
	assert.Equal(t, mathf.NewVector2i(3, 8), prod, "Multiplication failed")

	// Test division
	quot := v2.Div(v1)
	assert.Equal(t, mathf.NewVector2i(3, 2), quot, "Division failed")

	// Test scalar operations
	scalar := 2
	scalarAdd := v1.Addi(scalar)
	assert.Equal(t, mathf.NewVector2i(3, 4), scalarAdd, "Scalar addition failed")

	scalarSub := v1.Subi(scalar)
	assert.Equal(t, mathf.NewVector2i(-1, 0), scalarSub, "Scalar subtraction failed")

	scalarMul := v1.Muli(scalar)
	assert.Equal(t, mathf.NewVector2i(2, 4), scalarMul, "Scalar multiplication failed")

	scalarDiv := v2.Divi(scalar)
	assert.Equal(t, mathf.NewVector2i(1, 2), scalarDiv, "Scalar division failed")

	// Test dot product
	dot := v1.Dot(v2)
	assert.InDelta(t, float64(11), dot, 0.0001, "Dot product failed")

	// Test length calculations
	length := v1.Length()
	assert.InDelta(t, 2.236068, length, 0.0001, "Length calculation failed")

	lengthSquared := v1.LengthSquared()
	assert.Equal(t, 5, lengthSquared, "Length squared calculation failed")

	// Test distance calculations
	distance := v1.DistanceTo(v2)
	assert.InDelta(t, 2.828427, distance, 0.0001, "Distance calculation failed")

	distanceSquared := v1.DistanceSquaredTo(v2)
	assert.Equal(t, 8, distanceSquared, "Distance squared calculation failed")

	// Test abs
	negVector := mathf.NewVector2i(-1, -2)
	absVector := negVector.Abs()
	assert.Equal(t, v1, absVector, "Abs failed")

	// Test sign
	signVector := negVector.Sign()
	assert.Equal(t, mathf.NewVector2i(-1, -1), signVector, "Sign failed")

	// Test clamp
	min := mathf.NewVector2i(0, 0)
	max := mathf.NewVector2i(5, 5)
	clamped := v1.Clamp(min, max)
	assert.Equal(t, v1, clamped, "Clamp failed")

	// Test negation
	neg := v1.Neg()
	assert.Equal(t, mathf.NewVector2i(-1, -2), neg, "Negation failed")
}