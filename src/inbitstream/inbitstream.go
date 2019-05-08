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
	"github.com/piot/basal-go/src/types"
	brookinbitstream "github.com/piot/brook-go/src/inbitstream"
	"github.com/piot/cram-go/src/compression"
)

// InBitStream : InBitStream type
type InBitStream struct {
	stream brookinbitstream.InBitStream
}

// New : Create in bit stream
func New(stream brookinbitstream.InBitStream) *InBitStream {
	return &InBitStream{stream: stream}
}

func (s *InBitStream) ReadSignedScale(valueRange int, bits uint) (int32, error) {
	valuesPossible := 2 << (bits - 2)
	sv, readSignedErr := s.stream.ReadSignedBits(bits)
	if readSignedErr != nil {
		return 0.0, readSignedErr
	}
	v := int64(sv) * int64(valueRange) * types.FixedPointFactor / int64(valuesPossible)

	return int32(v), nil
}

// ReadVector3f : Reads a vector
func (s *InBitStream) ReadVector3f(valueRange int, bits uint) (types.Vector3f, error) {
	x, xErr := s.ReadSignedScale(valueRange, bits)
	if xErr != nil {
		return types.Vector3f{}, xErr
	}
	y, yErr := s.ReadSignedScale(valueRange, bits)
	if yErr != nil {
		return types.Vector3f{}, yErr
	}
	z, zErr := s.ReadSignedScale(valueRange, bits)
	if zErr != nil {
		return types.Vector3f{}, zErr
	}

	return types.NewVector3f(x, y, z), nil
}

// ReadRotation : Reads a rotation
func (s *InBitStream) ReadRotation() (types.Quaternion, error) {
	maxIndex, maxIndexErr := s.stream.ReadBits(3)
	if maxIndexErr != nil {
		return types.Quaternion{}, maxIndexErr
	}
	a, aErr := s.stream.ReadInt16()
	if aErr != nil {
		return types.Quaternion{}, aErr
	}
	b, bErr := s.stream.ReadInt16()
	if bErr != nil {
		return types.Quaternion{}, bErr
	}
	c, cErr := s.stream.ReadInt16()
	if maxIndexErr != cErr {
		return types.Quaternion{}, cErr
	}

	info := compression.QuaternionPackInfo{A: a, B: b, C: c, MaxIndex: byte(maxIndex)}
	q, qErr := compression.QuaternionUnPack(info)
	return q, qErr
}
