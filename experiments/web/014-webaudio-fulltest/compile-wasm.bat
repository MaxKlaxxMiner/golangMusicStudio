set GOOS=js
set GOARCH=wasm
go build -o main.wasm main.go
go build -o worklet.wasm worklet.go
wat2wasm -o workletWat.wasm worklet.wat
