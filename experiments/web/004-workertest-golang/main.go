package main

import (
	"fmt"
	"syscall/js"
)

func calc(this js.Value, args []js.Value) any {
	for i := 0; i < 1000000000; i++ {
	}
	return nil
}

func main() {
	fmt.Println("wasm: Hello World!")

	wg := js.Global().Get("window").Get("wg")
	wg.Set("calc", js.FuncOf(calc))

	<-make(chan bool, 0)
}
