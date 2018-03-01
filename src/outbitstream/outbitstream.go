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

// Package outbitstream ...
package outbitstream

import (
	"fmt"

	brookoutbitstream "github.com/piot/brook-go/src/outbitstream"
	"github.com/piot/cram-go/src/types"
)

// OutBitStream : OutBitStream type
type OutBitStream struct {
	stream *brookoutbitstream.OutBitStream
}

func (s *OutBitStream) writeSignedScale(v float32, rangeValue int, bits uint) error {
	if bits < 2 {
		return fmt.Errorf("Must write at least two bits")
	}
	valuesPossible := 2 << (bits - 2)
	av := int(v * float32(valuesPossible) / float32(rangeValue))
	s.stream.WriteSignedBits(int32(av), bits)

	return nil
}

// WriteVector3f : Write vector to stream
func (s *OutBitStream) WriteVector3f(v types.Vector3f, rangeValue int, bits uint) {
	s.writeSignedScale(v.X, rangeValue, bits)
	s.writeSignedScale(v.Y, rangeValue, bits)
	s.writeSignedScale(v.Z, rangeValue, bits)
}

// WriteQuaternion : Write quaternion to stream
func (s *OutBitStream) WriteQuaternion(v types.Quaternion) {
	s.stream.WriteBits(0, 16)
}
