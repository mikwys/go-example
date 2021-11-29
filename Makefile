APP_NAME=tws
windows:
	echo "Compiling for Windows"
	GOOS=windows go build -o ./bin/${APP_NAME}.exe ./cmd/${APP_NAME}/main.go

all:
	echo "Compiling for Darwin"
	GOARCH=amd64 GOOS=darwin go build -v -o ./bin/${APP_NAME}-darwin ./cmd/tws/main.go
	echo "Compiling for Linux"
	GOARCH=amd64 GOOS=linux go build -o ./bin/${APP_NAME}-linux ./cmd/tws/main.go
	echo "Compiling for Windows"
	GOARCH=amd64 GOOS=windows go build -o ./bin/${APP_NAME}-windows.exe ./cmd/tws/main.go

