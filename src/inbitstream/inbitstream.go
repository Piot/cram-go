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

// Package inbitstream ...
package inbitstream

import (
	brookinbitstream "github.com/piot/brook-go/src/inbitstream"
	"github.com/piot/cram-go/src/types"
)

// InBitStream : InBitStream type
type InBitStream struct {
	stream *brookinbitstream.InBitStream
}

func (s *InBitStream) readSignedScale(valueRange int, bits uint) float32 {
	valuesPossible := 2 << (bits - 2)
	sv, _ := s.stream.ReadSignedBits(bits)
	v := float32(sv) * float32(valueRange) / float32(valuesPossible)

	return float32(v)
}

// ReadVector3f : Reads a vector
func (s *InBitStream) ReadVector3f(valueRange int, bits uint) types.Vector3f {
	x := s.readSignedScale(valueRange, bits)
	y := s.readSignedScale(valueRange, bits)
	z := s.readSignedScale(valueRange, bits)

	return types.NewVector3f(x, y, z)
}

// ReadRotation : Reads a rotation
func (s *InBitStream) ReadRotation() types.Quaternion {
	s.stream.ReadBits(16)
	return types.NewQuaternion(0, 0, 0, 1)
}
