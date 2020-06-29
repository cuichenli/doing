package commands

import (
	"os"

	"github.com/spf13/cobra"
)

// DoingRecordFile File record the doing list.
const DoingRecordFile = "doing.md"

// RootCmd Root cmd parser
var RootCmd = &cobra.Command{
	Use:   "doing",
	Short: "A command line tool for remembering what you were doing and tracking what you've done.",
}

var configFile string

func init() {
	var err error
	configFile, err = getConfigFilePath()
	if err != nil {
		os.Exit(1)
	}
	RootCmd.AddCommand(addCommand)
	RootCmd.AddCommand(doneCommand)
}
