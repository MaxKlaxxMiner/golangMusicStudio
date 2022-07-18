package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmsampler"
	"math"
)

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
	printBufAuto("noise", "", gmsampler.Noise)
	//printBufAuto("square", "HQ", gmsampler.SquareHQ)
	return

	tone := 440.0                                // A3         440.0000000 hz
	tone *= math.Pow(math.Pow(2, 1.0/12), 3) * 4 // A3 -> C6  2093.0045224 hz
	samplePeriode := gmconst.SampleRate / tone   // samples per periode
	sampleIncrement := uint32(4294967296.0/samplePeriode + 0.5)

	buf1 := make([]int32, 128)
	//buf2 := make([]int32, 128)
	//buf3 := make([]int32, 128)

	//fillSampleSquareSimple(buf1, sampleIncrement, 0)
	//fillSampleSquareAF2(buf2, sampleIncrement, 0)
	//fillSampleSquareAF3(buf3, sampleIncrement, 0)
	gmsampler.Noise(buf1, sampleIncrement, 0)

	for i := range buf1 {
		fmt.Println(i, float64(buf1[i])/float64(1<<(gmconst.DynamicBits-1)))
		//fmt.Println(i, float64(buf1[i])/float64(1<<(gmconst.DynamicBits-1)), float64(buf2[i])/float64(1<<(gmconst.DynamicBits-1)))
		//fmt.Println(i, float64(buf1[i])/float64(1<<(gmconst.DynamicBits-1)), float64(buf2[i])/float64(1<<(gmconst.DynamicBits-1)), float64(buf3[i])/float64(1<<(gmconst.DynamicBits-1)))
	}

	const rs = 5
	const is = 1000000

	//fmt.Println()
	//for r := 0; r < rs; r++ {
	//	tim := time.Now()
	//	ofs := uint32(0)
	//	for i := 0; i < is; i++ {
	//		ofs = gmsampler.SquareHQ(buf1, sampleIncrement, ofs)
	//	}
	//	fmt.Println("simple:", time.Since(tim))
	//}
}
