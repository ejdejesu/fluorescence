package geometry

// import "math"

// // Point in a small extention of a Vector, representing a point in 3D space
// type PointX VectorX

// // PointZero is the zero point, or the origin
// var PointZero = PointX{}

// // PointMax is the maximum representable point
// var PointMax = PointX{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}

// // MinComponents returns the Point construction of the minimums of two points component-wise
// func MinComponents(p, q PointX) PointX {
// 	return PointX{math.Min(p.X() q.X), math.Min(p.Y(),q.Y), math.Min(p.Z(), q.Z)}
// }

// // MaxComponents returns the Point construction of the maximums of two points component-wise
// func MaxComponents(p, q PointX) PointX {
// 	return PointX{math.Max(p.X() q.X), math.Max(p.Y(),q.Y), math.Max(p.Z(), q.Z)}
// }

// // To finds a Vector pointing from p to q
// func (p PointX) To(q PointX) VectorX {
// 	return q.asVector().Sub(p.asVector())
// }

// // From finds a Vector pointing from q to p
// func (p PointX) From(q PointX) VectorX {
// 	return p.asVector().Sub(q.asVector())
// }

// // AddVector adds a Vector c to a Point p
// func (p PointX) AddVector(v VectorX) PointX {
// 	return PointX{p.X() + v.X() p.Y()+ v.Y() p.Z() + v.Z}
// }

// // SubPoint subtracts a Point q from a Point p
// func (p PointX) SubPoint(q PointX) VectorX {
// 	return p.asVector().Sub(q.asVector())
// }

// // SubVector subtracts a Vector v from a Point p
// func (p PointX) SubVector(v VectorX) PointX {
// 	return PointX{p.X() - v.X() p.Y()- v.Y() p.Z() - v.Z}
// }

// // NegateVec3 negates the components of a Point
// func (p PointX) NegateVec3() PointX {
// 	return PointX{-p.X() -p.Y() -p.Z}
// }

// // asVector converts a Point to a Vector
// func (p PointX) asVector() VectorX {
// 	return VectorX(p)
// }
