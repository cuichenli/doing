package commands

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cuichenli/doing/model"
	"github.com/spf13/cobra"
)

// DoingRecordFile File record the doing list.
const DoingRecordFile = "doing.md"

// RootCmd Root cmd parser
var RootCmd = &cobra.Command{
	Use:   "doing",
	Short: "A command line tool for remembering what you were doing and tracking what you've done.",
}

var configFile string

func init() {
	var err error
	configFile, err = GetConfigFilePath()
	if err != nil {
		os.Exit(1)
	}
	RootCmd.AddCommand(addCommand)
}

// GetConfigFilePath Get configuration path based on home directory.
func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to get home directory.")
		return "", err
	}
	fullPath := fmt.Sprintf("%s/%s", homeDir, DoingRecordFile)
	return fullPath, nil
}

// GetRecordsFromFile Get all records from the reader.
func GetRecordsFromFile(reader io.Reader) (model.RecordList, error) {
	scanner := bufio.NewScanner(reader)
	recordList := model.NewRecordList()
	for scanner.Scan() {
		text := scanner.Text()
		record, _ := model.FromTaskPaper(text)
		recordList.AddRecord(record)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return recordList, err
	}
	return recordList, nil
}

// GetExistingRecords Get all existing records
func GetExistingRecords() (model.RecordList, error) {
	file, err := os.OpenFile(configFile, os.O_RDONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to open record file. %s", err))
	}
	records, err := GetRecordsFromFile(file)
	if err != nil {
		return records, nil
	}
	return records, nil
}
