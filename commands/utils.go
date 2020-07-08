package commands

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
		entry := args[0]
		records, err := getExistingRecords()
		if err != nil {
			return err
		}
		record := newDoingRecord(entry)
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
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.TrimLeft(text, " -")
		record, _ := model.FromTaskPaper(text)
		recordList.AddRecord(record)
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
