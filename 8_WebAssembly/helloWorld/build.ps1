set GOARCH "wasm"
set GOOS "js"

go build -o lib.wasm main.go
