package geometry

import "testing"

var resultV float64
var resultVV Vector

func BenchmarkVectorMultiplicationScalar(b *testing.B) {
	v := Vector{
		X: 1.5,
		Y: 3.5,
		Z: 5.5,
	}
	var rV Vector
	floatN := float64(b.N)
	b.ResetTimer()
	for i := 0.0; i < floatN; i++ {
		rV = v.MultScalar(i)
	}
	resultVV = rV
}

func BenchmarkVectorDivision(b *testing.B) {
	v := Vector{
		X: 1.5,
		Y: 3.5,
		Z: 5.5,
	}
	var rV Vector
	floatN := float64(b.N)
	b.ResetTimer()
	for i := 0.0; i < floatN; i++ {
		rV = v.DivScalar(i)
	}
	resultVV = rV
}

func BenchmarkVectorUnit(b *testing.B) {
	v := Vector{
		X: 1.5,
		Y: 3.5,
		Z: 5.5,
	}
	var rV Vector
	floatN := float64(b.N)
	b.ResetTimer()
	for i := 0.0; i < floatN; i++ {
		rV = v.Unit()
	}
	resultVV = rV
}
