package commands

import (
	"errors"
	"testing"

	"github.com/cuichenli/doing/model"
	"github.com/onsi/gomega"
	"github.com/spf13/cobra"
)

func TestAddErrorWhenGetExistingRecords(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	oldGetExistingRecords := GetExistingRecords
	defer func() {
		GetExistingRecords = oldGetExistingRecords
	}()
	GetExistingRecords = func() (model.RecordList, error) {
		return model.NewRecordList(), errors.New("Error on GetExistingRecords")
	}
	err := Add(&cobra.Command{}, []string{"record"})
	g.Expect(err.Error()).Should(gomega.Equal("Error on GetExistingRecords"))
}
