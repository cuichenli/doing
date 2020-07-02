package commands

import (
	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

func init() {
	finishCommand.Flags().IntP("count", "c", 1, "Specify how many items to be marked as finished.")
}

var finishCommand = &cobra.Command{
	Use:   "finish",
	Short: "Finsish a numbers of last records.",
	Args:  cobra.NoArgs,
	RunE:  finish,
}

func finish(cmd *cobra.Command, args []string) error {
	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		return err
	}
	records, err := getExistingRecords()
	if err != nil {
		return err
	}
	finishLastItems(&records, count)
	return writeRecords(records)
}

func finishLastItems(records *model.RecordList, count int) {
	recordsList := records.GetAllRecords()
	recordsNum := len(*recordsList)
	for idx := recordsNum - 1; idx >= recordsNum-count; idx-- {
		record := &(*recordsList)[idx]
		record.Done()
	}
	return
}
