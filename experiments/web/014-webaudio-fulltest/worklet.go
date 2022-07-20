package main

import (
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmconst"
	"github.com/MaxKlaxxMiner/golangMusicStudio/gmstudio/gmsampler"
	"math/bits"
	"syscall/js"
	"unsafe"
)

func float32SliceAsByteSlice(floats []float32) []byte {
	lf := 4 * len(floats)

	// step by step
	pf := &(floats[0])              // To pointer to the first byte of b
	up := unsafe.Pointer(pf)        // To *special* unsafe.Pointer, it can be converted to any pointer
	pi := (*[1]byte)(up)            // To pointer as byte array
	buf := (*pi)[:]                 // Creates slice to our array of 1 byte
	address := unsafe.Pointer(&buf) // Capture the address to the slice structure
	if bits.UintSize == 64 {
		lenAddr := uintptr(address) + uintptr(8)  // Capture the address where the length and cap size is stored
		capAddr := uintptr(address) + uintptr(16) // WARNING: This is fragile, depending on a go-internal structure.
		lenPtr := (*int)(unsafe.Pointer(lenAddr)) // Create pointers to the length and cap size
		capPtr := (*int)(unsafe.Pointer(capAddr)) //
		*lenPtr = lf                              // Assign the actual slice size and cap
		*capPtr = lf                              //
	} else {
		lenAddr := uintptr(address) + uintptr(4)  // Capture the address where the length and cap size is stored
		capAddr := uintptr(address) + uintptr(8)  // WARNING: This is fragile, depending on a go-internal structure.
		lenPtr := (*int)(unsafe.Pointer(lenAddr)) // Create pointers to the length and cap size
		capPtr := (*int)(unsafe.Pointer(capAddr)) //
		*lenPtr = lf                              // Assign the actual slice size and cap
		*capPtr = lf                              //
	}
	return buf
}

func main() {
	wg := js.Global().Get("wg")
	wg.Set("workletGoReady", true)
	wg.Get("workletPort").Call("postMessage", "ok: goWasmReady")

	lastLen := 0
	samplesLeft := make([]float32, 0)
	samplesRight := make([]float32, 0)
	bufferLeft := make([]byte, 0)
	bufferRight := make([]byte, 0)
	samplerLeftBuffer := make([]int32, 0)
	samplerRightBuffer := make([]int32, 0)
	samplerLeftOfs := uint32(0)
	samplerRightOfs := uint32(0)

	wg.Set("workletGoFillBuffer", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return nil
		}
		bufferLen := args[0].Get("length").Int()
		if args[1].Get("length").Int() != bufferLen {
			return nil // invalid buffer size?
		}
		bufferLen /= 4 // 1 sample = 4 bytes, (stereo = 8 bytes)
		if bufferLen != lastLen {
			lastLen = bufferLen
			samplesLeft = make([]float32, bufferLen)
			samplesRight = make([]float32, bufferLen)
			bufferLeft = float32SliceAsByteSlice(samplesLeft)
			bufferRight = float32SliceAsByteSlice(samplesRight)
			samplerLeftBuffer = make([]int32, bufferLen)
			samplerRightBuffer = make([]int32, bufferLen)
		}

		samplerLeftOfs = gmsampler.Noise(samplerLeftBuffer, 123, samplerLeftOfs)
		samplerRightOfs = gmsampler.Noise(samplerRightBuffer, 456, samplerRightOfs)
		gmsampler.VolumeUpdate(samplerLeftBuffer, gmconst.Volume100/10)
		gmsampler.VolumeUpdate(samplerRightBuffer, gmconst.Volume100/10)

		for i := 0; i < bufferLen; i++ {
			samplesLeft[i] = float32(samplerLeftBuffer[i]) / float32(1<<(gmconst.DynamicBits-1))
			samplesRight[i] = float32(samplerRightBuffer[i]) / float32(1<<(gmconst.DynamicBits-1))
		}

		js.CopyBytesToJS(args[0], bufferLeft)
		js.CopyBytesToJS(args[1], bufferRight)

		return nil
	}))

	<-make(chan bool, 0)
}
