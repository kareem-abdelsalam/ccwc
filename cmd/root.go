package cmd

import (
	"ccwc/wcImplementation"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
)

var byteSizeFlag bool
var linesFlag bool
var wordsFlag bool
var charsFlag bool

func flushErrorMsgAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func convertStdinToTmpFile() *os.File {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		flushErrorMsgAndExit(err)
	}

	tmpFile, err := os.CreateTemp("", "ccwc-temp-os-stdin-*.txt")
	if err != nil {
		flushErrorMsgAndExit(err)
	}

	if _, err := tmpFile.Write(data); err != nil {
		flushErrorMsgAndExit(err)
	}

	return tmpFile
}

var rootCmd = &cobra.Command{
	Use:   "ccwc",
	Short: "ccwc is a reimplementation of wc tool",
	Long:  "ccwc is a reimplementation of wc tool",
	Run: func(cmd *cobra.Command, args []string) {
		var allFlagsFalse = !byteSizeFlag && !linesFlag && !wordsFlag && !charsFlag
		var fileIsStdin = false

		if len(args) == 0 {
			fileIsStdin = true

			tmpFile := convertStdinToTmpFile()
			defer os.Remove(tmpFile.Name())

			args = append(args, tmpFile.Name())
		}
		var output = make([][]string, 0)
		var myOS = wcImplementation.NewOSFS()
		for _, fileName := range args {

			file, err := myOS.Open(fileName)
			if err != nil {
				flushErrorMsgAndExit(err)
			}
			defer file.Close()

			fileOutput, err := wcImplementation.GetFileState(file,
				byteSizeFlag || allFlagsFalse, linesFlag || allFlagsFalse,
				wordsFlag || allFlagsFalse, charsFlag || allFlagsFalse)
			if err != nil {
				flushErrorMsgAndExit(err)
			}

			if fileIsStdin {
				fileOutput = append(fileOutput, "")
			} else {
				fileOutput = append(fileOutput, fileName)
			}

			output = append(output, fileOutput)
		}

		for _, fileOutput := range output {
			wcImplementation.PrintFileOutput(fileOutput)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&byteSizeFlag, "c", "c", false, "number of bytes")
	rootCmd.PersistentFlags().BoolVarP(&linesFlag, "l", "l", false, "number of lines")
	rootCmd.PersistentFlags().BoolVarP(&wordsFlag, "w", "w", false, "number of words")
	rootCmd.PersistentFlags().BoolVarP(&charsFlag, "m", "m", false, "number of characters")

	if err := rootCmd.Execute(); err != nil {
		flushErrorMsgAndExit(err)
	}
}
