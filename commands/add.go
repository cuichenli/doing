package commands

import (
	"fmt"

	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

func init() {
	addCommand.Flags().BoolP("editor", "e", false, "Invoke an editor to edit.")
}

var addCommand = &cobra.Command{
	Use:   "add [NOTE]",
	Short: "Add a record",
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
	RunE:    add,
	Aliases: []string{"now", "did"},
}

var add = genericAdd(func(_ *model.Record) {
	return
})
