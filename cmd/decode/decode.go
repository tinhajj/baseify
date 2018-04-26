package cmd

import (
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/tinhajj/baseify/cmd"
	"github.com/tinhajj/baseify/fileop"
)

func init() {
	cmd.RootCmd.AddCommand(DecodeCmd)
}

var DecodeCmd = &cobra.Command{
	Use:   "decode [BASE64FILE] [OUTPUTFILE]",
	Short: "base64 decode a file",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile(args[0])

		if err != nil {
			log.Fatalf("Unable to read file %v", args[0])
		}

		err = fileop.Decode(string(b), args[1])

		if err != nil {
			log.Fatalf("Unable to write decoded file to %v", args[1])
		}
	},
}
