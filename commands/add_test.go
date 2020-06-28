package commands

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
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

func TestAddWhenFailedToOpenFile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	oldGetExistingRecords := GetExistingRecords
	oldOpenFile := openFile
	defer func() {
		GetExistingRecords = oldGetExistingRecords
		openFile = oldOpenFile
	}()
	getExistingRecordsCount := 0
	GetExistingRecords = func() (model.RecordList, error) {
		getExistingRecordsCount++
		return model.NewRecordList(), nil
	}

	openFile = func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return &os.File{}, errors.New("Error on openning file")
	}
	err := Add(&cobra.Command{}, []string{"record"})
	g.Expect(err.Error()).Should(gomega.Equal("Error on openning file"))
	g.Expect(getExistingRecordsCount).To(gomega.Equal(1))
}

func TestAdd(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	oldGetExistingRecords := GetExistingRecords
	oldOpenFile := openFile
	oldNewRecordsWriter := newRecordsWriter
	defer func() {
		GetExistingRecords = oldGetExistingRecords
		openFile = oldOpenFile
		newRecordsWriter = oldNewRecordsWriter
	}()
	GetExistingRecords = func() (model.RecordList, error) {
		return model.NewRecordList(), nil
	}
	var buf bytes.Buffer
	newRecordsWriter = func(file string, writer io.Writer) model.RecordsWriter {
		return model.RecordsWriter{
			Writer:     &buf,
			ConfigFile: file,
		}
	}

	openFile = func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return &os.File{}, nil
	}
	err := Add(&cobra.Command{}, []string{"record"})
	g.Expect(err).To(gomega.BeNil())
	g.Expect(strings.HasPrefix(buf.String(), "  - record @created(")).To(gomega.BeTrue())
}
