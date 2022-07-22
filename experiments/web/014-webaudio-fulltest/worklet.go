package main

import (
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmextras"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmsampler"
	"syscall/js"
)

func main() {
	wg := js.Global().Get("wg")
	wg.Set("workletGoReady", true)
	wg.Get("workletPort").Call("postMessage", "ok: goWasmReady")

	floatSamples := make([]float32, gmconst.WorkletSampleCount)
	floatSamplesBytePtr := gmextras.MapSliceFloat32AsByte(floatSamples)

	samplesLeft, samplesRight := make([]int32, gmconst.WorkletSampleCount), make([]int32, gmconst.WorkletSampleCount)
	toneOfsLeft, toneOfsRight := uint32(0), uint32(0)

	wg.Set("workletGoFillBuffer", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return nil
		}
		bufLen := args[0].Get("length").Int()
		if bufLen != gmconst.WorkletSampleCount*4 || args[1].Get("length").Int() != bufLen {
			return nil // invalid buffer sizes?
		}

		toneOfsLeft = gmsampler.Noise(samplesLeft, 123, toneOfsLeft)    // random-noise, seed: 123
		toneOfsRight = gmsampler.Noise(samplesRight, 456, toneOfsRight) // random-noise, seed: 456
		gmsampler.VolumeUpdate(samplesLeft, gmconst.Volume100/10)       // 10% volume left
		gmsampler.VolumeUpdate(samplesRight, gmconst.Volume100/10)      // 10% volume right

		gmextras.ConvertIntSamplesToFloat32(floatSamples, samplesLeft)
		js.CopyBytesToJS(args[0], floatSamplesBytePtr)
		gmextras.ConvertIntSamplesToFloat32(floatSamples, samplesRight)
		js.CopyBytesToJS(args[1], floatSamplesBytePtr)

		return nil
	}))

	<-make(chan bool, 0)
}
