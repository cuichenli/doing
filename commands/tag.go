package commands

import (
	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

var tagCommand = &cobra.Command{
	Use:   "tag [NOTE]",
	Short: "Tag record",
	Args:  cobra.ArbitraryArgs,
	RunE:  tag,
}

func init() {
	tagCommand.Flags().IntP("count", "c", 1, "Specify how many items to be tagged")
	tagCommand.Flags().BoolP("remove", "r", false, "Remove the tag from records")
}

var tag = func(cmd *cobra.Command, args []string) error {
	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		return err
	}
	remove, err := cmd.Flags().GetBool("remove")
	if err != nil {
		return err
	}

	records, err := getExistingRecords()
	if err != nil {
		return err
	}
	var handler func(*model.Record) error
	if remove {
		handler = func(record *model.Record) error {
			record.RemoveTagsFromRawStringList(args)
			return nil
		}
	} else {
		handler = func(record *model.Record) error {
			return record.AddTagFromRawStringList(args)
		}
	}
	nums := len(records.Records)
	err = records.EditRecords(nums-count, nums, handler)
	if err != nil {
		return err
	}

	return writeRecords(records)
}
