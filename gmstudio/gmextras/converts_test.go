package gmextras

import (
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
	"testing"
)

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}

func TestConvertIntSamplesToFloat32(t *testing.T) {
	intSamples := []int32{
		0,
		gmconst.Volume100,
		-gmconst.Volume100,
		gmconst.Volume100 * 16,
		-gmconst.Volume100 * 16,
		gmconst.Volume100 / 4,
		-gmconst.Volume100 / 4,
	}
	floatSamples := []float32{
		0.0,
		1.0,
		-1.0,
		16.0,
		-16.0,
		0.25,
		-0.25,
	}

	floatTmp := make([]float32, len(floatSamples))

	ConvertIntSamplesToFloat32(floatTmp, intSamples)
	for i := 0; i < len(floatTmp); i++ {
		if floatTmp[i] != floatSamples[i] {
			t.Errorf("invalid float value for %d: %f, expected: %f", intSamples[i], floatTmp[i], floatSamples[i])
		}
	}
}

func TestPanicConvertIntSamplesToFloat32(t *testing.T) {
	intSamples := make([]int32, 123)
	floatSamples := make([]float32, 128)
	assertPanic(t, func() {
		ConvertIntSamplesToFloat32(floatSamples, intSamples)
	})
}
