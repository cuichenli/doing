package commands

import (
	"os"

	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:   "add [NOTE]",
	Short: "Add a record",
	Args:  cobra.ExactArgs(1),
	RunE:  Add,
}

var newDoingRecord = model.NewDoingRecord
var openFile = os.OpenFile
var newRecordsWriter = model.NewRecordsWriter

// Add Function executed for add command.
func Add(cmd *cobra.Command, args []string) error {
	entry := args[0]
	records, err := GetExistingRecords()
	if err != nil {
		return err
	}
	record := newDoingRecord(entry)
	records.AddRecord(record)
	file, err := openFile(configFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	writer := newRecordsWriter(configFile, file)
	writer.WriteToFile(records)
	return nil
}
