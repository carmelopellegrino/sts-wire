all: build build-windows build-macos

#bindata:
#	go get -u github.com/go-bindata/go-bindata/...
#	go-bindata -o rclone_bin.go data/

#data, err := Asset("data/rclone")
#if err != nil {
#	// Asset was not found.
#}

bind-rclone:
	mkdir -p data/linux
	wget -L --show-progress -O data/linux/rclone_linux https://github.com/dciangot/rclone/releases/download/v1.51.1-patch-s3/rclone
	go-bindata -o pkg/rclone/rclone_linux.go data/linux/
	sed -i "" 's/package\ main/package\ rclone/g' pkg/rclone/rclone_linux.go

bind-rclone-windows:
	mkdir -p data/windows
	wget -L --show-progress -O data/windows/rclone_windows https://github.com/dciangot/rclone/releases/download/v1.51.1-patch-s3/rclone.exe
	go-bindata -o pkg/rclone/rclone_windows.go data/windows/
	sed -i "" 's/package\ main/package\ rclone/g' pkg/rclone/rclone_windows.go

bind-rclone-macos:
	mkdir -p data/darwin
	wget -L --show-progress -O data/darwin/rclone_osx https://github.com/dciangot/rclone/releases/download/v1.51.1-patch-s3/rclone_osx
	go-bindata -o pkg/rclone/rclone_darwin.go data/darwin/
	sed -i "" 's/package\ main/package\ rclone/g' pkg/rclone/rclone_darwin.go

build: bind-rclone
	go build -ldflags "-s -w" -o sts-wire_linux

build-windows: bind-rclone-windows
	env GOOS=windows CGO_ENABLED=0 go build -ldflags "-s -w" -mod vendor -o sts-wire_windows.exe -v

build-macos: bind-rclone-macos
	env GOOS=darwin CGO_ENABLED=0 go build -ldflags "-s -w" -mod vendor -o sts-wire_osx -v