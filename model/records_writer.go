package model

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// RecordsWriter Model to write the records to the file.
type RecordsWriter struct {
	ConfigFile string
	Writer     io.Writer
}

// NewRecordsWriter Generate a new RecordsWriter.
func NewRecordsWriter(file string, writer io.Writer) RecordsWriter {
	return RecordsWriter{
		ConfigFile: file,
		Writer:     writer,
	}
}

// WriteToFile Write the recordList to the config file.
func (writer RecordsWriter) WriteToFile(recordList RecordList) error {
	records := recordList.GetAllRecords()
	stringBuilder := strings.Builder{}
	for _, record := range records {
		stringBuilder.WriteString(fmt.Sprintf("%s\n", record.ToTaskPaper()))
	}
	w := bufio.NewWriter(writer.Writer)
	_, err := w.WriteString(stringBuilder.String())
	w.Flush()
	if err != nil {
		return fmt.Errorf("Failed to write to file %s: %s", writer.ConfigFile, err)
	}
	return nil
}
