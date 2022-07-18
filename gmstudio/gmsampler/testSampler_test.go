package gmsampler

import (
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
	"math"
	"testing"
)

const SamplerTolerance = 0.001

func SamplerCompare(t *testing.T, samplerName string, compareValues []float32, sampler func(buf []int32, incr, ofs uint32) uint32, samplerTone string) {
	tone := 0.0
	switch samplerTone {
	case "A2":
		tone = 440.0 * 0.5 // 220.0000000 hz
	case "C4":
		tone = 440.0 * math.Pow(math.Pow(2, 1.0/12), 3) * 1.0 // 554.3652620 hz
	case "C6":
		tone = 440.0 * math.Pow(math.Pow(2, 1.0/12), 3) * 4.0 // 2093.0045224 hz
	case "C9":
		tone = 440.0 * math.Pow(math.Pow(2, 1.0/12), 3) * 32 // 16744.0361792 hz
	default:
		t.Errorf("unknown tone: %s", samplerTone)
	}

	samplePeriode := gmconst.SampleRate / tone // samples per periode
	sampleIncrement := uint32(4294967296.0/samplePeriode + 0.5)
	buf := make([]int32, (len(compareValues)+127)/128*128)
	ofs := uint32(0)
	for i := 0; i < len(buf); i += 128 {
		ofs = sampler(buf[i:i+128], sampleIncrement, ofs)
	}
	for i := 0; i < len(compareValues); i++ {
		sampleInt := buf[i]
		sampleFloat := float32(sampleInt) / float32(1<<(gmconst.DynamicBits-1))
		compareFloat := compareValues[i]
		compareMin := compareFloat - SamplerTolerance
		compareMax := compareFloat + SamplerTolerance
		if sampleFloat < compareMin || sampleFloat > compareMax {
			t.Errorf("%s - sample error tone %.2f hz at pos %d: %.8f (%d) not between %.8f and %.8f", samplerName, tone, i, sampleFloat, sampleInt, compareMin, compareMax)
			break
		}
	}
}
