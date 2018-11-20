package cramtest

import (
	"testing"

	"github.com/piot/basal-go/src/types"
	brookinbitstream "github.com/piot/brook-go/src/inbitstream"
	brookinstream "github.com/piot/brook-go/src/instream"
	brookoutbitstream "github.com/piot/brook-go/src/outbitstream"
	brookoutstream "github.com/piot/brook-go/src/outstream"
	craminbitstream "github.com/piot/cram-go/src/inbitstream"
	cramoutbitstream "github.com/piot/cram-go/src/outbitstream"
)

type TestStream struct {
	cramBitStream *cramoutbitstream.OutBitStream
	bitStream     brookoutbitstream.OutBitStream
	octetWriter   *brookoutstream.OutStream
}

func NewTestStream(cramBitStream *cramoutbitstream.OutBitStream, bitStream brookoutbitstream.OutBitStream, octetWriter *brookoutstream.OutStream) *TestStream {
	return &TestStream{cramBitStream: cramBitStream, bitStream: bitStream, octetWriter: octetWriter}
}

func (t *TestStream) CramStream() *cramoutbitstream.OutBitStream {
	return t.cramBitStream
}

func (t *TestStream) Flush() *craminbitstream.InBitStream {
	t.bitStream.Close()
	octets := t.octetWriter.Octets()
	octetReader := brookinstream.New(octets)
	inBitStream := brookinbitstream.New(octetReader, t.bitStream.Tell())
	cramInBitStream := craminbitstream.New(inBitStream)
	return cramInBitStream
}

func createCramStream() *TestStream {
	octetWriter := brookoutstream.New()
	bitStream := brookoutbitstream.New(octetWriter)
	cramStream := cramoutbitstream.New(bitStream)
	return NewTestStream(cramStream, bitStream, octetWriter)
}

func TestInOut(t *testing.T) {
	v := types.MakeVector3fFromFloats(999, 20, 20)
	testStream := createCramStream()
	cramStream := testStream.CramStream()

	cramStream.WriteVector3f(v, 1000, 24)
	inCramStream := testStream.Flush()

	readV, err := inCramStream.ReadVector3f(1000, 24)
	if err != nil {
		t.Fatal(err)
	}
	if readV.DistanceTo(v) > 1 {
		t.Errorf("vectors differ! dist:%v expected %v but received %v", readV.DistanceTo(v), v.DebugString(), readV.DebugString())
	}
}
