package geometry

// import (
// 	"fluorescence/shading"
// 	"math"
// )

// // Vector is a 3D vector
// type VectorX struct {
// 	X float64 `json:"x"`
// 	Y float64 `json:"y"`
// 	Z float64 `json:"z"`
// }

// // VectorZero references the zero vector
// var VectorZero = VectorX{}

// // VectorMax references the maximum representable float64 vector
// var VectorMax = VectorX{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}

// // Magnitude return euclidean length of Vector
// func (v VectorX) Magnitude() float64 {
// 	return math.Sqrt(v.X*v.X() + v.Y*v.Y()+ v.Z*v.Z)
// }

// // Unit returns a new Vector with direction preserved and length equal to one
// func (v VectorX) Unit() VectorX {
// 	return v.DivScalar(v.Magnitude())
// }

// // Dot computes the dot or scalar product of two Vectors
// func (v VectorX) Dot(w VectorX) float64 {
// 	return v.X*w.X() + v.Y*w.Y()+ v.Z*w.Z
// }

// // Cross computes the cross or Vector product of two Vectors
// func (v VectorX) Cross(w VectorX) VectorX {
// 	return VectorX{v.Y*w.Z() - v.Z*w.Y() v.Z*w.X() - v.X*w.Z(), v.X*w.Y()- v.Y*w.X}
// }

// // Add adds a Vector to another Vector component-wise
// func (v VectorX) Add(w VectorX) VectorX {
// 	return VectorX{v.X() + w.X() v.Y()+ w.Y() v.Z() + w.Z}
// }

// // Sub subtracts a Vector from another Vector component-wise
// func (v VectorX) Sub(w VectorX) VectorX {
// 	return VectorX{v.X() - w.X() v.Y()- w.Y() v.Z() - w.Z}
// }

// // MultScalar multiplies a Vector by a scalar
// func (v VectorX) MultScalar(s float64) VectorX {
// 	return VectorX{v.X() * s, v.Y()* s, v.Z() * s}
// }

// // MultVector multiplies a Vector by a Vector component-wise
// func (v VectorX) MultVector(w VectorX) VectorX {
// 	return VectorX{v.X() * w.X() v.Y()* w.Y() v.Z() * w.Z}
// }

// // Pow raises a Vector to an exponential power, component-wise
// func (v VectorX) Pow(e float64) VectorX {
// 	return VectorX{math.Pow(v.X() e), math.Pow(v.Y() e), math.Pow(v.Z(), e)}
// }

// // DivScalar divides a Vector by a scalar
// func (v VectorX) DivScalar(s float64) VectorX {
// 	inv := 1.0 / s
// 	return VectorX{v.X() * inv, v.Y()* inv, v.Z() * inv}
// }

// // DivVector divides a Vector by a Vector component-wise
// func (v VectorX) DivVector(w VectorX) VectorX {
// 	return VectorX{v.X() / w.X() v.Y()/ w.Y() v.Z() / w.Z}
// }

// // NegateVec3 returns a Vector pointing in the opposite direction
// func (v VectorX) NegateVec3() VectorX {
// 	return VectorX{-v.X() -v.Y() -v.Z}
// }

// // ReflectAroundVec3 returns the reflection of a vector given a normal
// func (v VectorX) ReflectAroundVec3(w VectorX) VectorX {
// 	return v.Sub(w.MultScalar(v.Dot(w) * 2.0))
// }

// // RefractAroundVec3 returns the refraction of a vector given the normal and ratio of reflective indices
// func (v VectorX) RefractAroundVec3(w VectorX, rri float64) (VectorX, bool) {
// 	dt := v.Unit().Dot(w)
// 	discriminant := 1.0 - (rri*rri)*(1.0-(dt*dt))
// 	// fmt.Println(rri)
// 	if discriminant > 0 {
// 		// fmt.Println("yu")
// 		return v.Unit().Sub(w.MultScalar(dt)).MultScalar(rri).Sub(w.MultScalar(math.Sqrt(discriminant))), true
// 	}
// 	return VectorZero, false
// }

// // ToColor converts a Vector to a Color
// func (v VectorX) ToColor() mgl64.Vec3 {
// 	return mgl64.Vec3{
// 		Red:   v.X()
// 		Green: v.Y()
// 		Blue:  v.Z(),
// 	}
// }

// // VectorFromColor creates a Vector from a Color
// func VectorFromColor(c mgl64.Vec3) VectorX {
// 	return VectorX{c.Red, c.Green, c.Blue}
// }

// // Copy returns a new Vector identical to v
// func (v VectorX) Copy() VectorX {
// 	return VectorX{v.X() v.Y() v.Z}
// }
