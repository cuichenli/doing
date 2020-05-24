package main

import "os"

func main() {
	err := addDoing("daba")
	if err != nil {
		os.Exit(1)
	}
}
