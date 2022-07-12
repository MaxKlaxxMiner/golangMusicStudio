package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("main-go: Hello World!")

	wg := js.Global().Get("window").Get("wg")
	_ = wg
	//wg.Set("calc", js.FuncOf(calcas))

	<-make(chan bool, 0)
}
