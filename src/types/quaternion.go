/*

MIT License

Copyright (c) 2017 Peter Bjorklund

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

*/

// Package types ...
package types

import (
	"fmt"
	"math"
)

// Quaternion : Quaternion type
type Quaternion struct {
	X float32
	Y float32
	Z float32
	W float32
}

// Index :
func (v Quaternion) Index(i int) float32 {
	switch i {
	case 0:
		return v.X
	case 1:
		return v.Y
	case 2:
		return v.Z
	case 3:
		return v.W
	default:
		return -1
	}
}

// NewQuaternion : Creates a new vector
func NewQuaternion(x float32, y float32, z float32, w float32) Quaternion {
	return Quaternion{X: x, Y: y, Z: z, W: w}
}

func almostEqual(a float32, b float32) bool {
	return math.Abs(float64(a-b)) < 0.0005
}

func (v Quaternion) almostEqual(o Quaternion) bool {
	return almostEqual(v.X, o.X) &&
		almostEqual(v.Y, o.Y) &&
		almostEqual(v.Z, o.Z) &&
		almostEqual(v.W, o.W)
}

func (v Quaternion) invert() Quaternion {
	return Quaternion{X: -v.X, Y: -v.Y, Z: -v.Z, W: -v.W}
}

// SameRepresentation :
func (v Quaternion) SameRepresentation(o Quaternion) bool {
	return v.almostEqual(o) || v.almostEqual(o.invert())
}

func (v Quaternion) String() string {
	return fmt.Sprintf("[quaternion %0.2f, %0.2f, %0.2f, %0.2f]", v.X, v.Y, v.Z, v.W)
}
