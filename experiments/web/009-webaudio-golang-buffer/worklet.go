package main

import (
	"fmt"
	"math"
	"math/bits"
	"strconv"
	"syscall/js"
	"time"
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
	fmt.Println("worklet: wasm start...")

	wg := js.Global().Get("wg")

	lastLen := 0
	samplesLeft := make([]float32, 0)
	samplesRight := make([]float32, 0)
	bufferLeft := make([]byte, 0)
	bufferRight := make([]byte, 0)

	wg.Set("fillBuffer", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return nil
		}
		bufferLen := args[0].Get("length").Int()
		if args[1].Get("length").Int() != bufferLen {
			return nil // inavalid right buffer?
		}
		bufferLen /= 4 // 1 sample = 4 bytes, (stereo = 8 bytes)
		if bufferLen != lastLen {
			lastLen = bufferLen
			samplesLeft = make([]float32, bufferLen)
			samplesRight = make([]float32, bufferLen)
			bufferLeft = float32SliceAsByteSlice(samplesLeft)
			bufferRight = float32SliceAsByteSlice(samplesRight)
			fmt.Println("worklet: buffer = " + strconv.Itoa(bufferLen) + " samples (" + strconv.Itoa(len(bufferLeft)+len(bufferRight)) + " bytes)")
		}

		volume := float32(math.Sin(math.Pi/500.0*float64(time.Now().UnixMilli()%1000)))*0.5 + 0.5

		for i := 0; i < bufferLen; i++ {
			samplesLeft[i] = float32(i-64) * 0.015 * volume * 0.2
			samplesRight[i] = -float32(i-64) * 0.015 * 0.05
		}

		js.CopyBytesToJS(args[0], bufferLeft)
		js.CopyBytesToJS(args[1], bufferRight)

		return nil
	}))

	<-make(chan bool, 0)
}
