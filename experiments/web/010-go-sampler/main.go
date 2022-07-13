package main

import (
	"fmt"
	"time"
)

const dynamicBits = 24 // 24 from 32 (reserve: 8 bits for 256x oversteer)
const aliasBits = 16   // 16 = 65536 subsamples

func fillSampleSquareSimple(buf []int32, incr, ofs uint32) uint32 {
	for i := 0; i < len(buf); i++ {
		buf[i] = int32(ofs>>31<<dynamicBits) - (1 << (dynamicBits - 1))
		ofs += incr
	}
	return ofs
}

func fillSampleSquareAA(buf []int32, incr, ofs uint32) uint32 {
	ofsA := uint64(ofs) << aliasBits
	for i := 0; i < len(buf); i++ {
		sum := int64(0)
		for a := 0; a < 1<<aliasBits; a++ {
			sum += int64(((ofsA>>(31+aliasBits))&1)<<dynamicBits) - (int64(1) << (dynamicBits - 1))
			ofsA += uint64(incr)
		}
		buf[i] = int32(sum >> aliasBits)
		ofs += incr
	}
	return ofs
}

func fillSampleSquareAFSums(ofsA *uint64, incr64 uint64, multiBits uint) int64 {
	if *ofsA>>(31+aliasBits) == (*ofsA+(incr64<<multiBits))>>(31+aliasBits) {
		sum := (int64((*ofsA>>(31+aliasBits))<<dynamicBits) - (int64(1) << (dynamicBits - 1))) << multiBits
		*ofsA += incr64 << multiBits
		return sum
	}
	if multiBits > 4 {
		sum := int64(0)
		sum += fillSampleSquareAFSums(ofsA, incr64, multiBits-1)
		sum += fillSampleSquareAFSums(ofsA, incr64, multiBits-1)
		return sum
	} else {
		sum := int64(0)
		off := *ofsA
		for a := 1 << multiBits; a != 0; a -= 4 {
			sum += int64((off & (uint64(1) << (31 + aliasBits))) >> (31 + aliasBits - dynamicBits))
			off += incr64
			sum += int64((off & (uint64(1) << (31 + aliasBits))) >> (31 + aliasBits - dynamicBits))
			off += incr64
			sum += int64((off & (uint64(1) << (31 + aliasBits))) >> (31 + aliasBits - dynamicBits))
			off += incr64
			sum += int64((off & (uint64(1) << (31 + aliasBits))) >> (31 + aliasBits - dynamicBits))
			off += incr64
		}
		sum -= (int64(1) << (dynamicBits - 1)) * (1 << multiBits)
		*ofsA = off & ((1 << (32 + aliasBits)) - 1)
		return sum
	}
}
func fillSampleSquareAF(buf []int32, incr, ofs uint32) uint32 {
	ofsA := uint64(ofs) << aliasBits
	for i := 0; i < len(buf); i++ {
		if ofsA>>(31+aliasBits) == (ofsA+(uint64(incr)<<aliasBits))>>(31+aliasBits) {
			buf[i] = int32(int64((ofsA>>(31+aliasBits))<<dynamicBits) - (int64(1) << (dynamicBits - 1)))
			ofsA += uint64(incr) << aliasBits
			continue
		}
		buf[i] = int32(fillSampleSquareAFSums(&ofsA, uint64(incr), aliasBits) >> aliasBits)
	}
	return ofs + incr*uint32(len(buf))
}

func fillSampleSquareAF2(buf []int32, incr, ofs uint32) uint32 {
	ofsA := uint64(ofs) << aliasBits
	for i := 0; i < len(buf); i++ {
		if ofsA>>(31+aliasBits) == (ofsA+(uint64(incr)<<aliasBits))>>(31+aliasBits) {
			buf[i] = int32(int64((ofsA>>(31+aliasBits))<<dynamicBits) - (int64(1) << (dynamicBits - 1)))
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

		buf[i] = int32(v * float64(1<<(dynamicBits)))
		ofsA = ofsA2
	}
	return ofs + incr*uint32(len(buf))
}

func main() {
	const sampleRate = 44100.0
	tone := 432.0
	tone *= 4                                                   // 1728 hz
	samplePeriode := sampleRate / tone                          // 25.5208333 samples per periode
	sampleIncrement := uint32(4294967296.0/samplePeriode + 0.5) // 168292596 = 1727.99999909 hz

	buf1 := make([]int32, 128)
	buf2 := make([]int32, 128)
	buf3 := make([]int32, 128)
	buf4 := make([]int32, 128)

	fillSampleSquareSimple(buf1, sampleIncrement, 0)
	fillSampleSquareAA(buf2, sampleIncrement, 0)
	fillSampleSquareAF(buf3, sampleIncrement, 0)
	fillSampleSquareAF2(buf4, sampleIncrement, 0)

	for i := range buf1 {
		fmt.Println(i, float64(buf1[i])/float64(1<<(dynamicBits-1)), float64(buf2[i])/float64(1<<(dynamicBits-1)), float64(buf3[i])/float64(1<<(dynamicBits-1)), float64(buf4[i])/float64(1<<(dynamicBits-1)))
	}

	const rs = 5
	const is = 100000

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
	//		ofs = fillSampleSquareAA(buf1, sampleIncrement, ofs)
	//	}
	//	fmt.Println("AA:", time.Since(tim))
	//}

	fmt.Println()
	for r := 0; r < rs; r++ {
		tim := time.Now()
		ofs := uint32(0)
		for i := 0; i < is; i++ {
			ofs = fillSampleSquareAF(buf1, sampleIncrement, ofs)
		}
		fmt.Println("AF:", time.Since(tim))
	}

	fmt.Println()
	for r := 0; r < rs; r++ {
		tim := time.Now()
		ofs := uint32(0)
		for i := 0; i < is; i++ {
			ofs = fillSampleSquareAF2(buf1, sampleIncrement, ofs)
		}
		fmt.Println("AF:", time.Since(tim))
	}
}
