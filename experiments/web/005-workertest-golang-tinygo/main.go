package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func calc(this js.Value, args []js.Value) any {
	sum := uint32(0)
	for i := uint32(0); i < 1000000000; i++ {
		sum += i
	}
	fmt.Println("sum: " + strconv.FormatInt(int64(sum), 10))
	return nil
}

func main() {
	fmt.Println("wasm: Hello World!")

	wg := js.Global().Get("window").Get("wg")
	wg.Set("calc", js.FuncOf(calc))

	<-make(chan bool, 0)
}
