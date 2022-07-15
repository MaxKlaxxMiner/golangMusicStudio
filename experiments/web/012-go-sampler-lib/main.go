package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
	"time"
)

const aliasBits = 16 // 16 = 65536 subsamples

func fillSampleSquareSimple(buf []int32, incr, ofs uint32) uint32 {
	for i := 0; i < len(buf); i++ {
		buf[i] = (1 << (gmconst.DynamicBits - 1)) - int32(ofs>>31<<gmconst.DynamicBits)
		ofs += incr
	}
	return ofs
}

func fillSampleSquareAF2(buf []int32, incr, ofs uint32) uint32 {
	ofsA := uint64(ofs) << aliasBits
	for i := 0; i < len(buf); i++ {
		if ofsA>>(31+aliasBits) == (ofsA+(uint64(incr)<<aliasBits))>>(31+aliasBits) {
			buf[i] = int32(int64((ofsA>>(31+aliasBits))<<gmconst.DynamicBits) - (int64(1) << (gmconst.DynamicBits - 1)))
			ofsA += uint64(incr) << aliasBits
			continue
		}
		ofsA2 := (ofsA + uint64(incr)<<aliasBits) & ((1 << (32 + aliasBits)) - 1)
		f1 := (1 << (31 + aliasBits)) - (ofsA & ((1 << (31 + aliasBits)) - 1))
		//f2 := ofsA2 & ((1 << (31 + aliasBits)) - 1)

		v := float64(f1) / float64(uint64(incr)<<aliasBits)
		if ofsA>>(31+aliasBits) != 0 {
			v = v - 0.5
		} else {
			v = 0.5 - v
		}

		buf[i] = int32(v * float64(1<<(gmconst.DynamicBits)))
		ofsA = ofsA2
	}
	return ofs + incr*uint32(len(buf))
}

func fillSampleSquareAF3(buf []int32, incr, ofs uint32) uint32 {
	curVal := int32(int64(ofs>>31<<gmconst.DynamicBits) - 1<<(gmconst.DynamicBits-1))
	incr1 := (1 << (gmconst.DynamicBits + gmconst.DynamicBits)) / uint64(incr)
	for i := 0; i < len(buf); i++ {
		if ofs>>31 == (ofs+incr)>>31 {
			buf[i] = curVal
			ofs += incr
			continue
		}
		ofsA2 := (ofs + incr) & (1<<32 - 1)
		f := 1<<31 - (ofs & (1<<31 - 1))
		curVal = int32((uint64(f) * incr1) >> gmconst.DynamicBits)
		if ofs>>31 != 0 {
			curVal = curVal - 1<<(gmconst.DynamicBits-1)
		} else {
			curVal = 1<<(gmconst.DynamicBits-1) - curVal
		}

		buf[i] = curVal
		ofs = ofsA2
		curVal = int32(int64(ofs>>31<<gmconst.DynamicBits) - 1<<(gmconst.DynamicBits-1))
	}
	return ofs // + incr*uint32(len(buf))
}

var bufSlic = [...]float32{-1, 1}

func fillSampleSquareFloat(buf []float32, incr, ofs float64) float64 {
	curInt := int(ofs * 2)
	curVal := float32((curInt << 1) - 1)
	incr1 := 1 / incr

	for i := 0; i < len(buf); i++ {
		nextOfs := ofs + incr
		nextInt := int(nextOfs * 2)

		if curInt == nextInt {
			buf[i] = curVal
			ofs = nextOfs
			continue
		}
		if curInt == 0 {
			buf[i] = float32((nextOfs + ofs - 1) * incr1)
			curInt = nextInt
			ofs = nextOfs
		} else {
			buf[i] = float32((2 - ofs - nextOfs) * incr1)
			curInt = nextInt - 2
			ofs = nextOfs - 1
		}
		curVal = float32((curInt << 1) - 1)
	}
	return ofs
}

func printBuf(buf []int32) {
	for i := 0; i < len(buf); i++ {
		val := fmt.Sprintf("%.3f", float64(buf[i])/float64(1<<(gmconst.DynamicBits-1)))
		for len(val) > 1 && (val[len(val)-1] == '0' || val[len(val)-1] == '.') {
			val = val[:len(val)-1]
		}
		fmt.Print(val + ",")
	}
}

func main() {
	const sampleRate = 44100.0
	tone := 440.0 // A3         440.0000000 hz
	//tone *= math.Pow(math.Pow(2, 1.0/12), 3)                    // A3 -> C4   523.2511306 hz
	//tone *= 4                                                   // C4 -> C6  2093.0045224 hz
	tone /= 2
	samplePeriode := sampleRate / tone                          // 21.07 samples per periode
	sampleIncrement := uint32(4294967296.0/samplePeriode + 0.5) // 203840952 2093.0045245 hz

	buf1 := make([]int32, 128)
	buf2 := make([]int32, 128)
	buf3 := make([]int32, 128)

	o := fillSampleSquareSimple(buf1, sampleIncrement, 0)
	printBuf(buf1)
	o = fillSampleSquareSimple(buf1, sampleIncrement, o)
	printBuf(buf1)
	fmt.Println()

	//return

	fillSampleSquareSimple(buf1, sampleIncrement, 0)
	fillSampleSquareAF2(buf2, sampleIncrement, 0)
	fillSampleSquareAF3(buf3, sampleIncrement, 0)

	for i := range buf1 {
		if buf1[i] == buf2[i] && i > 30 {
			//continue
		}
		fmt.Println(i, float64(buf1[i])/float64(1<<(gmconst.DynamicBits-1)), float64(buf2[i])/float64(1<<(gmconst.DynamicBits-1)), float64(buf3[i])/float64(1<<(gmconst.DynamicBits-1)))
	}

	const rs = 5
	const is = 1000000

	fmt.Println()
	for r := 0; r < rs; r++ {
		tim := time.Now()
		ofs := uint32(0)
		for i := 0; i < is; i++ {
			ofs = fillSampleSquareSimple(buf1, sampleIncrement, ofs)
		}
		fmt.Println("simple:", time.Since(tim))
	}

	//fmt.Println()
	//for r := 0; r < rs; r++ {
	//	tim := time.Now()
	//	ofs := uint32(0)
	//	for i := 0; i < is; i++ {
	//		ofs = fillSampleSquareAF2(buf1, sampleIncrement, ofs)
	//	}
	//	fmt.Println("AF2:", time.Since(tim))
	//}
	//
	//fmt.Println()
	//for r := 0; r < rs; r++ {
	//	tim := time.Now()
	//	ofs := float64(0)
	//	toneStep := tone / sampleRate
	//	for i := 0; i < is; i++ {
	//		ofs = fillSampleSquareFloat(buf4, toneStep, ofs)
	//	}
	//	fmt.Println("AFloat:", time.Since(tim))
	//}
	//
	//fmt.Println()
	//for r := 0; r < rs; r++ {
	//	tim := time.Now()
	//	ofs := uint32(0)
	//	for i := 0; i < is; i++ {
	//		ofs = fillSampleSquareAF3(buf1, sampleIncrement, ofs)
	//	}
	//	fmt.Println("AF3:", time.Since(tim))
	//}
}
