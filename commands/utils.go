package commands

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

var newDoingRecord = model.NewDoingRecord
var openFile = os.OpenFile
var newRecordsWriter = model.NewRecordsWriter
var userHomeDir = os.UserHomeDir
var logFaltal = log.Fatal

// genericAdd Returns a function to be used for cobra.Command.RunE.
// The callback parameter takes one function to handle the record to be added.
func genericAdd(callback func(*model.Record)) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {

		editor, err := cmd.Flags().GetBool("editor")
		if err != nil {
			return err
		}
		lines := make([]string, 0)
		if editor {
			homeDir, err := userHomeDir()
			if err != nil {
				return err
			}
			tempFile, err := ioutil.TempFile(homeDir, "doing-temp*")
			if err != nil {
				return err
			}
			fileName := tempFile.Name()
			tempFile.Close()
			editorCommand := exec.Command("vim", fileName)
			editorCommand.Stdin = os.Stdin
			editorCommand.Stdout = os.Stdout
			editorCommand.Run()
			file, err := os.Open(fileName)
			if err != nil {
				return nil
			}
			defer func() {
				file.Close()
				os.Remove(fileName)
			}()
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}
			if err = scanner.Err(); err != nil {
				return err
			}
		} else {
			entry := args[0]
			lines = strings.Split(entry, "\n")
		}
		records, err := getExistingRecords()
		if err != nil {
			return err
		}
		record := newDoingRecord(lines[0])
		if len(lines) > 1 {
			record.AddDetail(strings.Join(lines[1:], "\n"))
		}
		callback(&record)
		records.AddRecord(record)
		return writeRecords(records)
	}
}

func writeRecords(records model.RecordList) error {
	file, err := openFile(configFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	defer func() {
		file.Close()
	}()
	if err != nil {
		return err
	}
	writer := newRecordsWriter(configFile, file)
	err = writer.WriteToFile(records)
	return err
}

// getConfigFilePath Get configuration path based on home directory.
var getConfigFilePath = func() (string, error) {
	homeDir, err := userHomeDir()
	if err != nil {
		logFaltal("Failed to get home directory.")
		return "", err
	}
	fullPath := fmt.Sprintf("%s/%s", homeDir, DoingRecordFile)
	return fullPath, nil
}

// getRecordsFromFile Get all records from the reader.
var getRecordsFromFile = func(reader io.Reader) (model.RecordList, error) {
	scanner := bufio.NewScanner(reader)
	recordList := model.NewRecordList()
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); {
		if !strings.HasPrefix(lines[i], "  - ") {
			return recordList, fmt.Errorf("Failed to read records")
		}
		recordTexts := make([]string, 0)
		recordTexts = append(recordTexts, strings.TrimPrefix(lines[i], "  - "))
		j := i + 1
		for j < len(lines) && !strings.HasPrefix(lines[j], "  - ") {
			recordTexts = append(recordTexts, lines[j])
			j++
		}
		record, err := model.FromTaskPaper(strings.Join(recordTexts, "\n"))
		if err != nil {
			return recordList, err
		}
		recordList.AddRecord(record)
		i = j
	}

	if err := scanner.Err(); err != nil {
		logFaltal(err)
		return recordList, err
	}
	return recordList, nil
}

// getExistingRecords Get all existing records
var getExistingRecords = func() (model.RecordList, error) {
	file, err := openFile(configFile, os.O_RDONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		logFaltal(fmt.Sprintf("Failed to open record file. %s", err))
	}
	records, err := getRecordsFromFile(file)
	if err != nil {
		return records, nil
	}
	return records, nil
}
