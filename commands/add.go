package commands

import (
	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

var addCommand = &cobra.Command{
	Use:     "add [NOTE]",
	Short:   "Add a record",
	Args:    cobra.ExactArgs(1),
	RunE:    add,
	Aliases: []string{"now", "did"},
}

var add = genericAdd(func(_ *model.Record) {
	return
})
