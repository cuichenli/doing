package main

import (
	"fmt"
	"log"
	"os"
)

func addDoing(s string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Failed to get home directory.")
		return err
	}
	fullPath := fmt.Sprintf("%s/%s", homeDir, DoingRecordFile)
	file, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	defer file.Close()
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to open record file. %s", err))
		return err
	}

	_, err = file.Write([]byte(s))
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to write to file. %s", err))
		return err
	}
	return nil
}
