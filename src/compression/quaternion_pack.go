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
	"math"

	"github.com/piot/cram-go/src/types"
)

func findMaxElementAndSign(q *types.Quaternion) (int, int) {
	maxIndex := 0
	maxValue := float32(0)

	sign := 0

	for i := 0; i < 4; i++ {
		var element = q.Index(i)
		var abs = float32(math.Abs(float64(element)))

		if abs > maxValue {
			sign = 1
			if element < 0 {
				sign = -1
			}
			maxIndex = i
			maxValue = abs
		}
	}

	return maxIndex, sign
}

// QuaternionPack :
func QuaternionPack(q *types.Quaternion) QuaternionPackInfo {
	maxIndex, sign := findMaxElementAndSign(q)

	var a float32
	var b float32
	var c float32

	factor := float32(sign) * floatPrecisionMultiplier

	if maxIndex == 0 {
		a = q.Y * factor
		b = q.Z * factor
		c = q.W * factor
	} else if maxIndex == 1 {
		a = q.X * factor
		b = q.Z * factor
		c = q.W * factor
	} else if maxIndex == 2 {
		a = q.X * factor
		b = q.Y * factor
		c = q.W * factor
	} else {
		a = q.X * factor
		b = q.Y * factor
		c = q.Z * factor
	}

	return QuaternionPackInfo{MaxIndex: byte(maxIndex), A: int16(a), B: int16(b), C: int16(c)}
}
