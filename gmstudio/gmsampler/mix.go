package gmsampler

func Mix(buf, src1, src2 []int32) {
	for i := 0; i < len(buf) && i < len(src1) && i < len(src2); i++ {
		buf[i] = src1[i] + src2[i]
	}
}

func MixClamped(buf, src1, src2 []int32, clampMax int32) {
	min := int64(-clampMax)
	max := int64(clampMax)
	for i := 0; i < len(buf) && i < len(src1) && i < len(src2); i++ {
		sample := int64(src1[i]) + int64(src2[i])
		if sample < min {
			sample = min
		} else if sample > max {
			sample = max
		}
		buf[i] = int32(sample)
	}
}
