package main

import (
	"fmt"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmextras"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmsampler"
	"syscall/js"
)

func main() {
	wg := js.Global().Get("wg")
	wg.Set("mainGoReady", true)

	ints := make([]int32, gmconst.WorkletSampleCount)
	fmt.Println("noise:", int32(gmsampler.SquareHQ(ints, 123456789, 1234567890)))
	floats := make([]float32, len(ints))
	gmextras.ConvertIntSamplesToFloat32(floats, ints)

	for i := 0; i < len(floats); i += 8 {
		fmt.Println(i, floats[i], floats[i+1], floats[i+2], floats[i+3], floats[i+4], floats[i+5], floats[i+6], floats[i+7])
	}

	<-make(chan bool, 0)
}
