package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/tinhajj/baseify/fileop"
)

type Flags struct {
	Suffix string
}

var flags Flags

func init() {
	RootCmd.AddCommand(EncodeCmd)
	EncodeCmd.Flags().StringVarP(&flags.Suffix, "suffix", "s", "_base64", "A suffix to add to the generated files")
}

var EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "base64 encode a list of files",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Fatalf("Missing file argument")
		}

		pathCollect := make(map[string]bool)

		for _, arg := range args {
			matches, err := filepath.Glob(arg)

			if err != nil {
				log.Fatalf("Error processing argument %s", arg)
			}

			for _, match := range matches {
				pathCollect[match] = true
			}
		}

		files := fileop.OnlyFiles(pathCollect)

		for _, file := range files {
			enc, err := fileop.Encode(file)

			if err != nil {
				log.Fatalf("Error processing file %s", file)
			}

			if _, err := os.Stat(file + flags.Suffix); !os.IsNotExist(err) {
				log.Fatalf("Tried to output encoding to file, but it already existed %v", file+"suffix")
			}

			f, err := os.Create(file + flags.Suffix)
			f.WriteString(enc)
		}
	},
}
