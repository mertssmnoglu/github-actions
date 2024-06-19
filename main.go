package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World Application. Cross compiled using GoReleaser")
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}
