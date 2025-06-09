package encode

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"baseify/cmd"
	"baseify/fileop"

	"github.com/spf13/cobra"
)

type Flags struct {
	Suffix string
	Output bool
}

var flags Flags

func init() {
	cmd.RootCmd.AddCommand(EncodeCmd)
	EncodeCmd.Flags().StringVarP(&flags.Suffix, "suffix", "s", "_base64", "A suffix to add to the generated files")
	EncodeCmd.Flags().BoolVarP(&flags.Output, "output", "o", true, "Print the names of encoded files")
}

var EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "base64 encode a list of files",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
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
				log.Fatalf("Error processing file %s: %s", file, err)
			}

			if _, err := os.Stat(file + flags.Suffix); !os.IsNotExist(err) {
				log.Fatalf("Tried to output encoding to file, but it already exists %v", file+"suffix")
			}

			f, err := os.Create(file + flags.Suffix)
			if err != nil {
				log.Fatalf("Error creating file %s: %s", file, err)
			}
			f.WriteString(enc)

			if flags.Output {
				fmt.Println(file + flags.Suffix)
			}
		}
	},
}
