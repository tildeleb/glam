// Copyright (c) 2013 Laurent Moussault. All rights reserved.
// Licensed under a simplified BSD license (see LICENSE file).

package glam

import "github.com/drakmaniso/glam/math"

//------------------------------------------------------------------------------

// `Vec4` is single-precision vector with 4 components.
type Vec4 struct {
	X float32
	Y float32
	Z float32
	W float32
}

//------------------------------------------------------------------------------

// `Dehomogenized` returns the dehomogenization of `a` (perspective divide).
// `a.W` must be non-zero.
func (a Vec4) Dehomogenized() Vec3 {
	return Vec3{a.X / a.W, a.Y / a.W, a.Z / a.W}
}

//------------------------------------------------------------------------------

// `Plus` returns the sum `a + b`.
//
// See also `Add`.
func (a Vec4) Plus(b Vec4) Vec4 {
	return Vec4{a.X + b.X, a.Y + b.Y, a.Z + b.Z, a.W + b.W}
}

// `Add` sets `a` to the sum `a + b`.
//
// More efficient than `Plus`.
func (a *Vec4) Add(b Vec4) {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
	a.W += b.W
}

//------------------------------------------------------------------------------

// `Minus` returns the difference `a - b`.
//
// See also `Subtract`.
func (a Vec4) Minus(b Vec4) Vec4 {
	return Vec4{a.X - b.X, a.Y - b.Y, a.Z - b.Z, a.W - b.W}
}

// `Subtract` sets `a` to the difference `a - b`.
// More efficient than `Minus`.
func (a *Vec4) Subtract(b Vec4) {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
	a.W -= b.W
}

//------------------------------------------------------------------------------

// `Inverse` return the inverse of `a`.
//
// See also `Invert`.
func (a Vec4) Inverse() Vec4 {
	return Vec4{-a.X, -a.Y, -a.Z, -a.W}
}

// `Invert` sets `a` to its inverse.
// More efficient than `Inverse`.
func (a *Vec4) Invert() {
	a.X = -a.X
	a.Y = -a.Y
	a.Z = -a.Z
	a.W = -a.W
}

//------------------------------------------------------------------------------

// `Times` returns the product of `a` with the scalar `s`.
//
// See also `Multiply`.
func (a Vec4) Times(s float32) Vec4 {
	return Vec4{a.X * s, a.Y * s, a.Z * s, a.W * s}
}

// `Multiply` sets `a` to the product of `a` with the scalar `s`.
// More efficient than `Times`.
func (a *Vec4) Multiply(s float32) {
	a.X *= s
	a.Y *= s
	a.Z *= s
	a.W *= s
}

//------------------------------------------------------------------------------

// `Slash` returns the division of `a` by the scalar `s`.
// `s` must be non-zero.
//
// See also `Divide`.
func (a Vec4) Slash(s float32) Vec4 {
	return Vec4{a.X / s, a.Y / s, a.Z / s, a.W / s}
}

// `Divide` sets `a` to the division of `a` by the scalar `s`.
// `s` must be non-zero.
//
// More efficient than `Slash`.
func (a *Vec4) Divide(s float32) {
	a.X /= s
	a.Y /= s
	a.Z /= s
	a.W /= s
}

// `Cross` returns the cross product of `a` and `b`.
func (a Vec4) Cross(b Vec4) Vec4 {
	return Vec4{
		a.Y*b.Z - a.Z*b.Y,
		a.Z*b.X - a.X*b.Z,
		a.X*b.Y - a.Y*b.X,
		a.W,
	}
}

//------------------------------------------------------------------------------

// `Dot` returns the dot product of `a` and `b`.
func (a Vec4) Dot(b Vec4) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z + a.W*b.W
}

func (a Vec4) Dot3(b Vec4) float32 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

//------------------------------------------------------------------------------

// `Length` returns `|a|` (the euclidian length of `a`).
func (a Vec4) Length() float32 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
}

// `Normalized` return `a/|a|` (i.e. the normalization of `a`).
// `a` must be non-zero.
//
// See also `Normalize`.
func (a Vec4) Normalized() Vec4 {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
	return Vec4{a.X / length, a.Y / length, a.Z / length, a.W / length}
}

// `Normalize` sets `a` to `a/|a|` (i.e. normalizes `a`).
// `a` must be non-zero.
//
// More efficitent than `Normalized`.
func (a *Vec4) Normalize() {
	length := math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z + a.W*a.W)
	a.X /= length
	a.Y /= length
	a.Z /= length
	a.W /= length
}

//------------------------------------------------------------------------------
