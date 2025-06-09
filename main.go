package main

import (
	"baseify/cmd"

	_ "baseify/cmd/decode"
	_ "baseify/cmd/encode"
)

func main() {
	cmd.Execute()
}
