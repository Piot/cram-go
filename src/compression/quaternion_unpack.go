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
	"fmt"
	"math"

	"github.com/piot/cram-go/src/types"
)

// QuaternionUnPack :
func QuaternionUnPack(info QuaternionPackInfo) (types.Quaternion, error) {
	a := float32(info.A) / floatPrecisionMultiplier
	b := float32(info.B) / floatPrecisionMultiplier
	c := float32(info.C) / floatPrecisionMultiplier
	maxIndex := info.MaxIndex

	d := float32(math.Sqrt(float64(1.0 - (a*a + b*b + c*c))))

	switch maxIndex {
	case 0:
		return types.NewQuaternion(d, a, b, c), nil
	case 1:
		return types.NewQuaternion(a, d, b, c), nil
	case 2:
		return types.NewQuaternion(a, b, d, c), nil
	case 3:
		return types.NewQuaternion(a, b, c, d), nil
	default:
		return types.NewQuaternion(0, 0, 0, 0), fmt.Errorf("Unknown max index")
	}
}
