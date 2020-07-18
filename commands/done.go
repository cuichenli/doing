package commands

import (
	"fmt"

	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

func init() {
	doneCommand.Flags().BoolP("editor", "e", false, "Invoke an editor to edit.")
}

var doneCommand = &cobra.Command{
	Use:   "done [NOTE]",
	Short: "Add a done record",
	Args: func(cmd *cobra.Command, args []string) error {
		editor, err := cmd.Flags().GetBool("editor")
		if err != nil {
			return err
		}
		if editor {
			if len(args) > 0 {
				return fmt.Errorf("Should not provide detail with -e switch enabled")
			}
		} else {
			if len(args) != 1 {
				return fmt.Errorf("Only one command line parameter for add is acceptable")
			}
		}
		return nil
	},
	RunE: done,
}

var done = genericAdd(func(record *model.Record) {
	record.Done()
})
