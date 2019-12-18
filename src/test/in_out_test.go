package cramtest

import (
	"testing"

	"github.com/piot/basal-go/src/types"
	brookinbitstream "github.com/piot/brook-go/src/inbitstream"
	brookoutbitstream "github.com/piot/brook-go/src/outbitstream"
	craminbitstream "github.com/piot/cram-go/src/inbitstream"
	cramoutbitstream "github.com/piot/cram-go/src/outbitstream"
)

type TestStream struct {
	cramBitStream *cramoutbitstream.OutBitStream
	bitStream     brookoutbitstream.OutBitStream
}

func NewTestStream(cramBitStream *cramoutbitstream.OutBitStream, bitStream brookoutbitstream.OutBitStream) *TestStream {
	return &TestStream{cramBitStream: cramBitStream, bitStream: bitStream}
}

func (t *TestStream) CramStream() *cramoutbitstream.OutBitStream {
	return t.cramBitStream
}

func (t *TestStream) Flush() *craminbitstream.InBitStream {
	t.bitStream.Close()
	octets := t.bitStream.Octets()
	inBitStream := brookinbitstream.New(octets, t.bitStream.Tell())
	cramInBitStream := craminbitstream.New(inBitStream)

	return cramInBitStream
}

func createCramStream() *TestStream {
	bitStream := brookoutbitstream.New(1024)
	cramStream := cramoutbitstream.New(bitStream)

	return NewTestStream(cramStream, bitStream)
}

func TestInOut(t *testing.T) {
	v := types.MakeVector3fFromFloats(999, 20, 20)
	testStream := createCramStream()
	cramStream := testStream.CramStream()

	writeErr := cramStream.WriteVector3f(v, 1000, 24)
	if writeErr != nil {
		t.Fatal(writeErr)
	}

	inCramStream := testStream.Flush()

	readV, err := inCramStream.ReadVector3f(1000, 24)
	if err != nil {
		t.Fatal(err)
	}

	if readV.DistanceTo(v) > 1 {
		t.Errorf("vectors differ! dist:%v expected %v but received %v",
			readV.DistanceTo(v), v.DebugString(), readV.DebugString())
	}
}
