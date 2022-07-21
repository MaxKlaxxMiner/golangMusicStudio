package gmextras

import (
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
)

func ConvertIntSamplesToFloat32(dst []float32, src []int32) {
	if len(dst) != len(src) {
		panic("invalid buffer sizes")
	}
	for i := 0; i < len(dst); i++ {
		dst[i] = float32(src[i]) / float32(1<<(gmconst.DynamicBits-1))
	}
}
