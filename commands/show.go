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
		show(&records)
		return nil
	},
}

var show = func(recordList *model.RecordList) {
	records := recordList.GetAllRecords()
	for _, record := range *records {
		fmt.Println(record.ToDisplayString())
	}
}
