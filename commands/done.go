package commands

import (
	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

var doneCommand = &cobra.Command{
	Use:   "done [NOTE]",
	Short: "Add a done record",
	Args:  cobra.ExactArgs(1),
	RunE:  done,
}

var done = GenericAdd(func(record *model.Record) {
	record.Done()
})
