package main

import (
	"handlers/processor"
)

func main() {
	processor.NewShellProcessor().Run()
}
