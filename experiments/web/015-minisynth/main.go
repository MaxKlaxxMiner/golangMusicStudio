package main

import (
	"syscall/js"
)

func main() {
	wg := js.Global().Get("wg")
	wg.Set("mainGoReady", true)
}
