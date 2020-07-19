package commands

import (
	"fmt"

	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

var showCommand = &cobra.Command{
	Use:   "show",
	Short: "Show all records",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		records, err := getExistingRecords()
		if err != nil {
			return err
		}
		return show(&records)
	},
}

var show = func(recordList *model.RecordList) error {
	records := recordList.GetAllRecords()
	for _, record := range *records {
		str, err := record.ToDisplayString()
		if err != nil {
			return err
		}
		fmt.Println(str)
	}
	return nil
}
