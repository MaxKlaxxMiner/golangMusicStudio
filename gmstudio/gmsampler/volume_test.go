package gmsampler

import (
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
	"strconv"
	"testing"
)

func testTone(buf []int32, incr, ofs uint32) uint32 {
	for i := 0; i < len(buf); i++ {
		buf[i] = int32(ofs)
		ofs += incr
	}
	return ofs
}

var testVolumtInt int32
var testVolumeFloat float32
var testVolumeClampedInt int32
var testVolumeClampedFloat float32

func setTestVolume(volume int32) {
	testVolumtInt = volume
	testVolumeFloat = float32(volume) / float32(gmconst.Volume100)
}

func setTestVolumeClamped(max int32) {
	testVolumeClampedInt = max
	testVolumeClampedFloat = float32(max) / float32(gmconst.Volume100)
}

func testToneVol(buf []int32, incr, ofs uint32) uint32 {
	ofs = testTone(buf, incr, ofs)
	if testVolumeClampedInt > 0 {
		VolumeUpdateClamped(buf, testVolumtInt, testVolumeClampedInt)
	} else {
		VolumeUpdate(buf, testVolumtInt)
	}
	return ofs
}

func testToneFloats(incr uint32) []float32 {
	buf := make([]int32, 128)
	testTone(buf, incr, 0)
	result := make([]float32, len(buf))
	for i := 0; i < len(buf); i++ {
		result[i] = float32(buf[i]) / float32(1<<(gmconst.DynamicBits-1)) * testVolumeFloat
		if testVolumeClampedInt > 0 {
			if result[i] < -testVolumeClampedFloat {
				result[i] = -testVolumeClampedFloat
			} else if result[i] > testVolumeClampedFloat {
				result[i] = testVolumeClampedFloat
			}
		}
	}
	return result
}

func testVolumeRound(t *testing.T) {
	for incr := uint32(1); incr < 4000000000; incr *= 3 {
		setTestVolume(gmconst.Volume100)
		SamplerCompareIncr(t, "volume 100% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		setTestVolume(gmconst.Volume100 / 2)
		SamplerCompareIncr(t, "volume 50% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		setTestVolume(gmconst.Volume100 * 42 / 100)
		SamplerCompareIncr(t, "volume 42% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		setTestVolume(gmconst.Volume100 / 3)
		SamplerCompareIncr(t, "volume 33.33% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		setTestVolume(gmconst.Volume100 * 11 / 100)
		SamplerCompareIncr(t, "volume 11% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		setTestVolume(gmconst.Volume100 / 100)
		SamplerCompareIncr(t, "volume 1% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		setTestVolume(gmconst.Volume100 / 1000)
		SamplerCompareIncr(t, "volume 0.1% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		setTestVolume(gmconst.Volume100 / 10000)
		SamplerCompareIncr(t, "volume 0.01% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		setTestVolume(0)
		SamplerCompareIncr(t, "volume 0% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		if testVolumeClampedInt > 0 {
			setTestVolume(gmconst.Volume100 * 2)
			SamplerCompareIncr(t, "volume 200% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
			setTestVolume(gmconst.Volume100 * 3)
			SamplerCompareIncr(t, "volume 300% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
			setTestVolume(gmconst.Volume100 * 10)
			SamplerCompareIncr(t, "volume 1000% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
			setTestVolume(gmconst.Volume100 * 42)
			SamplerCompareIncr(t, "volume 4200% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
			setTestVolume(gmconst.Volume100 * 100)
			SamplerCompareIncr(t, "volume 10000% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
			setTestVolume(gmconst.VolumeLimit)
			SamplerCompareIncr(t, "volume limit clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		}
		setTestVolume(-gmconst.Volume100)
		SamplerCompareIncr(t, "volume -100% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
		setTestVolume(-gmconst.Volume100 / 2)
		SamplerCompareIncr(t, "volume -50% clamp "+strconv.Itoa(int(testVolumeClampedInt))+" incr "+strconv.Itoa(int(incr)), testToneFloats(incr), testToneVol, incr)
	}
}

func TestVolumeBasics(t *testing.T) {
	setTestVolumeClamped(0)
	testVolumeRound(t)
}

func TestVolumeClampedLimit(t *testing.T) {
	setTestVolumeClamped(gmconst.VolumeLimit)
	testVolumeRound(t)
}

func TestVolumeClamped100p(t *testing.T) {
	setTestVolumeClamped(gmconst.Volume100)
	testVolumeRound(t)
}

func TestVolumeClamped200p(t *testing.T) {
	setTestVolumeClamped(gmconst.Volume100 * 2)
	testVolumeRound(t)
}

func TestVolumeClamped50p(t *testing.T) {
	setTestVolumeClamped(gmconst.Volume100 / 2)
	testVolumeRound(t)
}
