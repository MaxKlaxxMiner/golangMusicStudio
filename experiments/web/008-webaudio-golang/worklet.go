package main

import "fmt"

func main() {
	fmt.Println("worklet-go: Hello World!")

	//wg := js.Global().Get("window").Get("wg")
	//_ = wg
	//wg.Set("calc", js.FuncOf(calcas))

	<-make(chan bool, 0)
}
