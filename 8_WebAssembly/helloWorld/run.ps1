$GOROOT = $env:GOROOT

$copied = Test-Path "./wasm_exec.js"
if ($copied -eq $false) {
    Copy-Item "${GOROOT}/misc/wasm/wasm_exec.js" .
}
node wasm_exec.js helloWorld.js