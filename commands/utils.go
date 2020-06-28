package commands

import (
	"os"

	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

// GenericAdd Returns a function to be used for cobra.Command.RunE.
// The callback parameter takes one function to handle the record to be added.
func GenericAdd(callback func(*model.Record)) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		entry := args[0]
		records, err := GetExistingRecords()
		if err != nil {
			return err
		}
		record := newDoingRecord(entry)
		callback(&record)
		records.AddRecord(record)
		file, err := openFile(configFile, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		writer := newRecordsWriter(configFile, file)
		err = writer.WriteToFile(records)
		return err
	}
}
