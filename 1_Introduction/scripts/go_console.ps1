
$mountcmd="Mount-DiskImage -ImagePath C:\Users\Jonas\Master\S1\KP\hidden\kp_go.vhdx"

powershell -noprofile -command "&{ start-process powershell -ArgumentList '-noprofile ${mountcmd}' -verb RunAs}"

$env:SEU_HOME="G:"

$env:GOROOT="${env:SEU_HOME}/software/goroot"
$env:GOPATH="${env:SEU_HOME}/codebase"
$env:GOBIN="${env:SEU_HOME}/codebase/bin"
$env:PATH="${env:SEU_HOME}/software/goroot/bin;${env:PATH}"

Set-Location "${env:SEU_HOME}/codebase/src"
powershell.exe
