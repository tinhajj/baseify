package cmd

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tinhajj/baseify/fileop"
)

func init() {
	RootCmd.AddCommand(EncodeCmd)
}

var EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "base64 encode a list of files",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Missing file argument")
		}

		var files []string

		for _, arg := range args {
			matches, err := filepath.Glob(arg)

			if err != nil {
				log.Fatalf("Error processing argument %s", arg)
			}

			files = matches
		}

		if len(files) < 1 {
			log.Fatalf("No files to process")
		}

		for _, file := range files {
			enc, err := fileop.Encode(file)

			if err != nil {
				log.Fatalf("Error processing file %s", file)
			}

			fmt.Println(file)
		}
	},
}
