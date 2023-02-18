package main

import (
	"fmt"

	ActiveProcess "github.com/makssof/go-active-process"
)

func main() {
	activeProcess, err := ActiveProcess.Get()

	if err != nil {
		fmt.Println(fmt.Errorf("error: %v", err))
	} else {
		fmt.Println(activeProcess)
	}
}
