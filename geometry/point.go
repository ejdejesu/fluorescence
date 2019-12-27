package geometry

import "math"

// Point in a small extention of a Vector, representing a point in 3D space
type Point Vector

// PointZero is the zero point, or the origin
var PointZero = Point{}

// PointMax is the maximum representable point
var PointMax = Point{math.MaxFloat64, math.MaxFloat64, math.MaxFloat64}

// MinComponents returns the Point construction of the minimums of two points component-wise
func MinComponents(p, q Point) Point {
	return Point{math.Min(p.X, q.X), math.Min(p.Y, q.Y), math.Min(p.Z, q.Z)}
}

// MaxComponents returns the Point construction of the maximums of two points component-wise
func MaxComponents(p, q Point) Point {
	return Point{math.Max(p.X, q.X), math.Max(p.Y, q.Y), math.Max(p.Z, q.Z)}
}

// To finds a Vector pointing from p to q
func (p Point) To(q Point) Vector {
	return q.asVector().Sub(p.asVector())
}

// From finds a Vector pointing from q to p
func (p Point) From(q Point) Vector {
	return p.asVector().Sub(q.asVector())
}

// AddVector adds a Vector c to a Point p
func (p Point) AddVector(v Vector) Point {
	return Point{p.X + v.X, p.Y + v.Y, p.Z + v.Z}
}

// SubPoint subtracts a Point q from a Point p
func (p Point) SubPoint(q Point) Vector {
	return p.asVector().Sub(q.asVector())
}

// SubVector subtracts a Vector v from a Point p
func (p Point) SubVector(v Vector) Point {
	return Point{p.X - v.X, p.Y - v.Y, p.Z - v.Z}
}

// Negate negates the components of a Point
func (p Point) Negate() Point {
	return Point{-p.X, -p.Y, -p.Z}
}

// asVector converts a Point to a Vector
func (p Point) asVector() Vector {
	return Vector(p)
}
