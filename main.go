package main

import (
	"github.com/tinhajj/baseify/cmd"

	_ "github.com/tinhajj/baseify/cmd/decode"
	_ "github.com/tinhajj/baseify/cmd/encode"
)

func main() {
	cmd.Execute()
}
