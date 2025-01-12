package main

import (
	"fmt"
	"github.com/GlennLucy1/learn-ai222/api"
	"github.com/GlennLucy1/learn-ai222/ctrl"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("usage: %s <api-addr> <ocr-addr", os.Args[0])
		return
	}
	ctrl.OcrAddr = os.Args[2]
	api.StartServer(os.Args[1])
}
