package commands

import (
	"bytes"
	"errors"
	"fmt"
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
	oldGetExistingRecords := getExistingRecords
	defer func() {
		getExistingRecords = oldGetExistingRecords
	}()
	getExistingRecords = func() (model.RecordList, error) {
		return model.NewRecordList(), errors.New("Error on GetExistingRecords")
	}
	err := add(&cobra.Command{}, []string{"record"})
	g.Expect(err.Error()).Should(gomega.Equal("Error on GetExistingRecords"))
}

func TestAddDoingWhenFailedToOpenFile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	oldGetExistingRecords := getExistingRecords
	oldOpenFile := openFile
	defer func() {
		getExistingRecords = oldGetExistingRecords
		openFile = oldOpenFile
	}()
	getExistingRecordsCount := 0
	getExistingRecords = func() (model.RecordList, error) {
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
	oldGetExistingRecords := getExistingRecords
	oldOpenFile := openFile
	oldNewRecordsWriter := newRecordsWriter
	defer func() {
		getExistingRecords = oldGetExistingRecords
		openFile = oldOpenFile
		newRecordsWriter = oldNewRecordsWriter
	}()
	getExistingRecords = func() (model.RecordList, error) {
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
	oldGetExistingRecords := getExistingRecords
	oldOpenFile := openFile
	oldNewRecordsWriter := newRecordsWriter
	defer func() {
		getExistingRecords = oldGetExistingRecords
		openFile = oldOpenFile
		newRecordsWriter = oldNewRecordsWriter
	}()
	getExistingRecords = func() (model.RecordList, error) {
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

func TestGetConfigFilePathWhenFailedToReadHomeDir(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	oldUserHomeDir := userHomeDir
	oldLogFatal := logFaltal
	defer func() {
		userHomeDir = oldUserHomeDir
		logFaltal = oldLogFatal
	}()

	userHomeDir = func() (string, error) {
		return "", errors.New("Error on get user home dir")
	}
	logFaltal = func(v ...interface{}) {}
	path, err := getConfigFilePath()
	g.Expect(path).To(gomega.Equal(""))
	g.Expect(err.Error()).To(gomega.Equal("Error on get user home dir"))
}

func TestGetConfigFilePath(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	oldUserHomeDir := userHomeDir
	defer func() {
		userHomeDir = oldUserHomeDir
	}()

	userHomeDir = func() (string, error) {
		return "/home/doing", nil
	}
	path, err := getConfigFilePath()
	g.Expect(path).To(gomega.Equal("/home/doing/doing.md"))
	g.Expect(err).To(gomega.BeNil())
}

func TestGetRecordsFromFile(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	var buf bytes.Buffer
	fmt.Fprintf(&buf, `  - a doing record @created(2020-06-28T21:27:50+10:00)
	- a done record @created(2020-06-28T21:29:52+10:00) @done`)
	records, err := getRecordsFromFile(&buf)
	g.Expect(err).To(gomega.BeNil())
	g.Expect(len(records.DoingRecords)).To(gomega.Equal(1))
	g.Expect(len(records.DoneRecords)).To(gomega.Equal(1))
}

func TestGetExistingRecords(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	oldOpenFile := openFile
	oldGetRecordsFromFile := getRecordsFromFile
	defer func() {
		openFile = oldOpenFile
		getRecordsFromFile = oldGetRecordsFromFile
	}()
	openFile = func(name string, flag int, perm os.FileMode) (*os.File, error) {
		return &os.File{}, nil
	}
	getRecordsFromFileCount := 0
	getRecordsFromFile = func(reader io.Reader) (model.RecordList, error) {
		getRecordsFromFileCount++
		return model.RecordList{}, nil
	}

	_, err := getExistingRecords()
	g.Expect(getRecordsFromFileCount).To(gomega.Equal(1))
	g.Expect(err).To(gomega.BeNil())
}
