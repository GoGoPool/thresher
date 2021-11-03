# Justfiles are better Makefiles. 
# Install the `just` command from here https://github.com/casey/just

VERSION := `grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/'`
GIT_COMMIT := `git rev-parse HEAD`
BUILD_DATE := `date '+%Y-%m-%d-%H:%M:%S'`
VERSION_PATH := "github.com/johnthethird/thresher/version"
LDFLAGS:= "-X " + VERSION_PATH + ".BuildDate=" + BUILD_DATE + " -X " + VERSION_PATH + ".Version=" + VERSION + " -X " + VERSION_PATH + ".GitCommit=" + GIT_COMMIT

default:
  @just --list

compile: clean
	GOOS=darwin GOARCH=amd64  go build -ldflags "{{LDFLAGS}}" -o bin/thresher main.go

compile-windows:
	GOOS=windows GOARCH=amd64 go build -ldflags "{{LDFLAGS}}" -o ./bin/thresher.exe main.go

compile-linux:
	GOOS=linux GOARCH=amd64   go build -ldflags "{{LDFLAGS}}" -o ./bin/thresher-linux-x64 main.go

compile-linux-arm:
	GOOS=linux GOARCH=arm64   go build -ldflags "{{LDFLAGS}}" -o ./bin/thresher-linux-arm64 main.go

clean:
	/bin/rm -f bin/*

# Quickly show the main UI for development
ui: 
	bin/thresher --config alice.json testui

# Initialize alice, bob, and cam as users with Avalanche Fuji addresses.
initusers:
	bin/thresher init avalanche fuji DAO-Treasury alice X-fuji1knjauvyjxf56tavysqnf9zxds084588nqja7j4
	bin/thresher init avalanche fuji DAO-Treasury bob X-fuji1uehmke49qtysde4p2ehvnpvp7sc6j8xdntrma0
	bin/thresher init avalanche fuji DAO-Treasury cam X-fuji13avtfecrzkhxrd8mxqcd0ehctsvqh99y6xjnr2

alice:
	bin/thresher --config DAO-Treasury-alice.json --log alice.log wallet

bob:
	bin/thresher --config DAO-Treasury-bob.json --log bob.log wallet

cam: 
	bin/thresher --config DAO-Treasury-cam.json --log cam.log wallet

