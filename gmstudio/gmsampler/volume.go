package gmsampler

import "github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"

func VolumeUpdate(buf []int32, vol int32) {
	for i := 0; i < len(buf); i++ {
		buf[i] = int32((int64(buf[i]) * int64(vol)) >> (gmconst.DynamicBits - 1))
	}
}

func VolumeUpdateClamped(buf []int32, vol, clampMax int32) {
	min := int64(-clampMax)
	max := int64(clampMax)
	for i := 0; i < len(buf); i++ {
		sample := (int64(buf[i]) * int64(vol)) >> (gmconst.DynamicBits - 1)
		if sample < min {
			sample = min
		} else if sample > max {
			sample = max
		}
		buf[i] = int32(sample)
	}
}
