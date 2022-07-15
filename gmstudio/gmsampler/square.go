package gmsampler

import "github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"

func SquareLQ(buf []int32, incr, ofs uint32) uint32 {
	for i := 0; i < len(buf); i++ {
		buf[i] = (1 << (gmconst.DynamicBits - 1)) - int32(ofs>>(gmconst.SampleBits-1)<<gmconst.DynamicBits)
		ofs += incr
	}
	return ofs
}

func SquareMQ(buf []int32, incr, ofs uint32) uint32 {
	return SquareHQ(buf, incr, ofs)
}

func SquareHQ(buf []int32, incr, ofs uint32) uint32 {
	curVal := int32(int64(ofs>>(gmconst.SampleBits-1)<<gmconst.DynamicBits) - 1<<(gmconst.DynamicBits-1))
	incr1 := (1 << (gmconst.DynamicBits + gmconst.DynamicBits)) / uint64(incr)
	for i := 0; i < len(buf); i++ {
		if ofs>>(gmconst.SampleBits-1) == (ofs+incr)>>(gmconst.SampleBits-1) {
			buf[i] = curVal
			ofs += incr
			continue
		}
		ofsA2 := (ofs + incr) & (1<<gmconst.SampleBits - 1)
		f := 1<<(gmconst.SampleBits-1) - (ofs & (1<<(gmconst.SampleBits-1) - 1))
		curVal = int32((uint64(f) * incr1) >> gmconst.DynamicBits)
		if ofs>>(gmconst.SampleBits-1) != 0 {
			curVal = curVal - 1<<(gmconst.DynamicBits-1)
		} else {
			curVal = 1<<(gmconst.DynamicBits-1) - curVal
		}

		buf[i] = curVal
		ofs = ofsA2
		curVal = int32(int64(ofs>>(gmconst.SampleBits-1)<<gmconst.DynamicBits) - 1<<(gmconst.DynamicBits-1))
	}
	return ofs
}
