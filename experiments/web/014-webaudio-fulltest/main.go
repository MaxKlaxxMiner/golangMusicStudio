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
	fmt.Println("noise:", int32(gmsampler.Noise(ints, 123, 456)))
	floats := make([]float32, len(ints))
	gmextras.ConvertIntSamplesToFloat32(floats, ints)

	for i := 0; i < len(floats); i += 4 {
		fmt.Println(i, floats[i], floats[i+1], floats[i+2], floats[i+3])
	}

	<-make(chan bool, 0)
}
