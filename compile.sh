# Build for Windows/386
export GOOS=windows
export GOARCH=386
go build -o ./build/sort.x86.exe
# Build for Windows/amd64
export GOOS=windows
export GOARCH=amd64
go build -o ./build/sort.x64.exe
# Build for Linux/386
export GOOS=linux
export GOARCH=386
go build -o ./build/sort.i386.bin
# Build for Linux/amd64
export GOOS=linux
export GOARCH=amd64
go build -o ./build/sort.amd64.bin
# Link the example builds
mkdir test
ln -f ./build/sort.amd64.bin ./test/sort.amd64.bin
ln -f ./build/sort.i386.bin ./test/sort.i386.bin