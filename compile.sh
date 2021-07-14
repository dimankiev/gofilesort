# Build for Windows/386
export GOOS=windows
export GOARCH=386
go build -o ./build/sort.windows.x86.exe main.go
# Build for Windows/amd64
export GOOS=windows
export GOARCH=amd64
go build -o ./build/sort.windows.x64.exe main.go
# Build for Linux/386
export GOOS=linux
export GOARCH=386
go build -o ./build/sort.linux.i386.bin main.go
# Build for Linux/amd64
export GOOS=linux
export GOARCH=amd64
go build -o ./build/sort.linux.amd64.bin main.go
# Link the example builds
mkdir test
ln -f ./build/sort.linux.amd64.bin ./test/sort.linux.amd64.bin
ln -f ./build/sort.linux.i386.bin ./test/sort.linux.i386.bin