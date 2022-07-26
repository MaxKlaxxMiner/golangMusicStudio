set GOOS=js
set GOARCH=wasm
go build -o main.wasm main.go
wat2wasm -o worklet.wasm worklet.wat
