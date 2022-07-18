package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
	"math"
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

func fillSampleSquareAF4(buf []int32, incr, ofs uint32) uint32 {
	curVal := int32(int64(ofs>>31<<gmconst.DynamicBits) - 1<<(gmconst.DynamicBits-1))
	incr1 := (1 << (gmconst.DynamicBits + gmconst.DynamicBits)) / uint64(incr)
	for i := 0; i < len(buf); i++ {
		if ofs>>31 == (ofs+incr)>>31 {
			buf[i] = -curVal
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

		buf[i] = -curVal
		ofs = ofsA2
		curVal = int32(int64(ofs>>31<<gmconst.DynamicBits) - 1<<(gmconst.DynamicBits-1))
	}
	return ofs // + incr*uint32(len(buf))
}

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

func printBufAutoInternal(name string, sampler func(buf []int32, incr, ofs uint32) uint32, tone float64) {
	samplePeriode := gmconst.SampleRate / tone // samples per periode
	sampleIncrement := uint32(4294967296.0/samplePeriode + 0.5)
	buf := make([]int32, 256)
	ofs := sampler(buf[:128], sampleIncrement, 0)
	sampler(buf[128:], sampleIncrement, ofs)

	fmt.Print("var " + name + " = []float32{")
	newLine := true
	periodePos := 0.0
	periodeStep := samplePeriode
	for periodeStep < 20 {
		periodeStep *= 2
	}
	for i := 0; i < len(buf); i++ {
		if newLine {
			fmt.Print("\n    /* ")
			if i < 100 {
				fmt.Print(" ")
			}
			if i < 10 {
				fmt.Print(" ")
			}
			fmt.Print(i)
			fmt.Print(" */ ")
			newLine = false
		}
		val := fmt.Sprintf("%.3f", float64(buf[i])/float64(1<<(gmconst.DynamicBits-1)))
		for len(val) > 1 && (val[len(val)-1] == '0' || val[len(val)-1] == '.') {
			val = val[:len(val)-1]
		}
		fmt.Print(val + ", ")
		periodePos++
		if periodePos > periodeStep {
			periodePos -= periodeStep
			newLine = true
		}
	}
	fmt.Println()
	fmt.Println("}")
	fmt.Println()
}

func printBufAuto(name, quali string, sampler func(buf []int32, incr, ofs uint32) uint32) {
	toneA2 := 440.0 / 2                                 // A2   220.0000000 hz
	toneC3 := toneA2 * math.Pow(math.Pow(2, 1.0/12), 3) // C3   261.6255653 hz
	toneC4 := toneC3 * 2                                // C4   554.3652620 hz
	toneC6 := toneC4 * 4                                // C6  2093.0045224 hz
	toneC9 := toneC6 * 8                                // C9 16744.0361792 hz

	printBufAutoInternal(name+"Samples"+quali+"A2", sampler, toneA2)
	printBufAutoInternal(name+"Samples"+quali+"C4", sampler, toneC4)
	printBufAutoInternal(name+"Samples"+quali+"C6", sampler, toneC6)
	printBufAutoInternal(name+"Samples"+quali+"C9", sampler, toneC9)
}

func main() {
	printBufAuto("square", "LQ", fillSampleSquareSimple)
	printBufAuto("square", "HQ", fillSampleSquareAF4)
	return

	tone := 440.0 // A3         440.0000000 hz
	//tone /= 2                                                   // A3 -> A2   220.0000000 hz
	//tone *= math.Pow(math.Pow(2, 1.0/12), 3) * 4                // A3 -> C6  2093.0045224 hz
	tone *= math.Pow(math.Pow(2, 1.0/12), 3) * 32 // A3 -> C9 16744.0361792 hz
	samplePeriode := gmconst.SampleRate / tone    // samples per periode
	sampleIncrement := uint32(4294967296.0/samplePeriode + 0.5)

	buf1 := make([]int32, 128)
	buf2 := make([]int32, 128)
	buf3 := make([]int32, 128)

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
