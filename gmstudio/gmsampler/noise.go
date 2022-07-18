package gmsampler

import "github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"

func Noise(buf []int32, incr, ofs uint32) uint32 {
	for i := 0; i < len(buf); i++ {
		buf[i] = int32(ofs<<(gmconst.SampleBits-gmconst.DynamicBits-4)) >> (gmconst.SampleBits - gmconst.DynamicBits)
		ofs = (ofs+incr)*214013 + 2531011
	}
	return ofs
}
