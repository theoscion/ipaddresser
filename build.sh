#!/bin/bash
clear

#!/bin/bash
go test -cover ./...
testResult=$?

if [ $testResult -ne 0 ]; then
    echo
    echo "***UNIT TESTS FAILED***"
    exit $testResult
fi

echo -n "Building for Darwin..."
env GOOS=darwin GOARCH=386 go build -o ./build/ipaddresser-darwin-386
env GOOS=darwin GOARCH=amd64 go build -o ./build/ipaddresser-darwin-amd64
echo "Done"

echo -n "Building for FreeBSD..."
env GOOS=freebsd GOARCH=386 go build -o ./build/ipaddresser-freebsd-386
env GOOS=freebsd GOARCH=amd64 go build -o ./build/ipaddresser-freebsd-amd64
echo "Done"

echo -n "Building for Linux..."
env GOOS=linux GOARCH=386 go build -o ./build/ipaddresser-linux-386
env GOOS=linux GOARCH=amd64 go build -o ./build/ipaddresser-linux-amd64
env GOOS=linux GOARCH=ppc64 go build -o ./build/ipaddresser-linux-ppc64
env GOOS=linux GOARCH=ppc64le go build -o ./build/ipaddresser-linux-ppc64le
echo "Done"

echo -n "Building for Windows..."
env GOOS=windows GOARCH=386 go build -o ./build/ipaddresser-windows-386.exe
env GOOS=windows GOARCH=amd64 go build -o ./build/ipaddresser-windows-amd64.exe
echo "Done"

exit 0