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

package compression

import (
	"testing"

	"github.com/piot/cram-go/src/types"
)

const EPSILON float32 = 0.0005

func floatAlmostEqual(a, b float32) bool {
	return (a-b) < EPSILON && (b-a) < EPSILON
}

func almostEqual(a types.Quaternion, b types.Quaternion) bool {
	return floatAlmostEqual(a.X, b.X) && floatAlmostEqual(a.Y, b.Y) && floatAlmostEqual(a.Z, b.Z) && floatAlmostEqual(a.W, b.W)
}

func TestReadTenBits(t *testing.T) {
	q := types.NewQuaternion(-0.183, 0.683, -0.062, 0.704)
	info := QuaternionPack(&q)
	q2, err := QuaternionUnPack(info)
	if err != nil {
		t.Error(err)
	}

	if !almostEqual(q, q2) {
		t.Errorf("Quaternions are not equal")
	}
}
