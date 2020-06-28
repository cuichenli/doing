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

func TestAddDoingGetErrorWhenGetExistingRecords(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	oldGetExistingRecords := GetExistingRecords
	defer func() {
		GetExistingRecords = oldGetExistingRecords
	}()
	GetExistingRecords = func() (model.RecordList, error) {
		return model.NewRecordList(), errors.New("Error on GetExistingRecords")
	}
	err := add(&cobra.Command{}, []string{"record"})
	g.Expect(err.Error()).Should(gomega.Equal("Error on GetExistingRecords"))
}

func TestAddDoingWhenFailedToOpenFile(t *testing.T) {
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
	err := add(&cobra.Command{}, []string{"record"})
	g.Expect(err.Error()).Should(gomega.Equal("Error on openning file"))
	g.Expect(getExistingRecordsCount).To(gomega.Equal(1))
}

func TestAddDoing(t *testing.T) {
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
	err := add(&cobra.Command{}, []string{"record"})
	g.Expect(err).To(gomega.BeNil())
	g.Expect(strings.HasPrefix(buf.String(), "  - record @created(")).To(gomega.BeTrue())
}

func TestAddDone(t *testing.T) {
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
	err := done(&cobra.Command{}, []string{"record"})
	g.Expect(err).To(gomega.BeNil())
	output := buf.String()
	g.Expect(strings.HasPrefix(output, "  - record @created(")).To(gomega.BeTrue())
	// Remove the trailing new line symbol
	output = strings.TrimSpace(output)
	g.Expect(strings.HasSuffix(output, "@done")).To(gomega.BeTrue())
}
