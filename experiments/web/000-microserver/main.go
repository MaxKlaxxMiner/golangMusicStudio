package main

import (
	"fmt"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	_ = mime.AddExtensionType(".js", "application/javascript")
	_ = mime.AddExtensionType(".wasm", "application/wasm")

	relFolder := "."
	if len(os.Args) > 1 {
		relFolder = os.Args[1]
	}

	folder, err := filepath.Abs(relFolder)
	if err != nil {
		panic(err)
	}

	http.Handle("/", http.FileServer(http.Dir(folder)))

	fmt.Println("path:", folder)
	fmt.Println("run server: localhost:9090")
	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
