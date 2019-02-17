package main

import (
	"handlers/processor"
	"os"
	"utils/constant"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] != constant.EmptyString {
		processor.NewFileProcessor(os.Args[1]).Run()
		return
	}

	processor.NewShellProcessor().Run()
}
