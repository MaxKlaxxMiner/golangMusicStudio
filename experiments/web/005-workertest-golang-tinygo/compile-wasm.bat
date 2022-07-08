set GOOS=js
set GOARCH=wasm
tinygo build -target wasm -o main.wasm -size full main.go
