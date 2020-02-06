package tsz

import (
	"io"
	"testing"
)

// go test -v  . -run  TestReadDestructive
//     bstream_test.go:19: true <nil> 0 <nil>
//     bstream_test.go:20: before [128 0] after [0 0]
func TestReadDestructive(t *testing.T) {
	b := newBWriter(2)
	b.writeBit(true)
	b.writeByte(0)
	buf := b.bytes()
	cp := append([]byte{}, buf...)

	r := newBReader(buf)
	b1, e1 := r.readBit()
	byt1, e2 := r.readByte()
	t.Log(b1, e1, byt1, e2)
	t.Logf("before %v after %v", cp, buf)
}

func TestNewBWriter(t *testing.T) {
	b := newBWriter(1)
	if b.count != 0 {
		t.Errorf("Unexpected value: %v\n", b.count)
	}
}

func TestReadBitEOF1(t *testing.T) {
	b := newBWriter(1)
	_, err := b.readBit()
	if err != io.EOF {
		t.Errorf("Unexpected value: %v\n", err)
	}
}

func TestReadBitEOF2(t *testing.T) {
	b := newBReader([]byte{1})
	b.count = 0
	_, err := b.readBit()
	if err != io.EOF {
		t.Errorf("Unexpected value: %v\n", err)
	}
}

func TestReadByteEOF1(t *testing.T) {
	b := newBWriter(1)
	_, err := b.readByte()
	if err != io.EOF {
		t.Errorf("Unexpected value: %v\n", err)
	}
}

func TestReadByteEOF2(t *testing.T) {
	b := newBReader([]byte{1})
	b.count = 0
	_, err := b.readByte()
	if err != io.EOF {
		t.Errorf("Unexpected value: %v\n", err)
	}
}

func TestReadByteEOF3(t *testing.T) {
	b := newBReader([]byte{1})
	b.count = 16
	_, err := b.readByte()
	if err != io.EOF {
		t.Errorf("Unexpected value: %v\n", err)
	}
}

func TestReadBitsEOF(t *testing.T) {
	b := newBReader([]byte{1})
	_, err := b.readBits(9)
	if err != io.EOF {
		t.Errorf("Unexpected value: %v\n", err)
	}
}

func TestUnmarshalBinaryErr(t *testing.T) {
	b := &bstream{}
	err := b.UnmarshalBinary([]byte{})
	if err == nil {
		t.Errorf("An error was expected\n")
	}
}
