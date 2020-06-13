package model_test

import (
	"bytes"
	"strings"

	"github.com/cuichenli/doing/model"
)

var _ = Describe("RecordsWriter", func() {
	Context("NewRecordsWriter", func() {
		It("Should create one RecordsWriter", func() {
			fileName := "file"
			var bufferWriter bytes.Buffer

			writer := model.NewRecordsWriter(fileName, &bufferWriter)
			Expect(writer.ConfigFile).To(Equal(fileName))
		})
	})

	Context("WriteToFile", func() {
		It("Should write records to io", func() {
			recordList := model.NewRecordList()
			recordList.AddRecord(CreateDummyRecord("first"))
			recordList.AddRecord(CreateDummyRecord("second"))
			fileName := "file"
			var bufferWriter bytes.Buffer

			writer := model.NewRecordsWriter(fileName, &bufferWriter)
			writer.WriteToFile(recordList)
			output := strings.Split(strings.TrimRight(bufferWriter.String(), " \n"), "\n")
			Expect(len(output)).To(Equal(2))
			Expect(output[0]).To(Equal("  - first @created(2015-04-03T12:20:07Z)"))
			Expect(output[1]).To(Equal("  - second @created(2015-04-03T12:20:07Z)"))
		})
	})
})
