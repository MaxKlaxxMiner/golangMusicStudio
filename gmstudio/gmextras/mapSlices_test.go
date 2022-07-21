package gmextras

import (
	"encoding/binary"
	"math"
	"testing"
)

func TestMapSliceFloat32AsByte(t *testing.T) {
	floatSamples := []float32{
		0.0,
		1.0,
		-1.0,
		16.0,
		-16.0,
		0.25,
		-0.25,
	}

	expectedBytes := make([]byte, len(floatSamples)*4)
	for i := 0; i < len(floatSamples); i++ {
		floatBits := math.Float32bits(floatSamples[i])
		binary.LittleEndian.PutUint32(expectedBytes[i*4:], floatBits)
	}

	mappedBytes := MapSliceFloat32AsByte(floatSamples)
	if len(mappedBytes) != len(expectedBytes) {
		t.Errorf("invalid slice size %d expected: %d", len(mappedBytes), len(expectedBytes))
		return
	}
	for i := 0; i < len(expectedBytes); i++ {
		if mappedBytes[i] != expectedBytes[i] {
			t.Errorf("pos: %d invalid byte value: %d expected: %d", i, mappedBytes[i], expectedBytes[i])
		}
	}
}
