package model_test

import (
	"github.com/cuichenli/doing/model"
)

var _ = Describe("RecordList", func() {
	var (
		dummyRecords []model.Record
		recordList   model.RecordList
	)

	Context("RecordList.AddRecord", func() {
		BeforeEach(func() {
			dummyRecords = make([]model.Record, 0)
			recordList = model.NewRecordList()
		})

		It("Add one record to the recordList", func() {
			const detail = "One Task"
			dummyRecord := CreateDummyRecord(detail)
			dummyRecords = append(dummyRecords, dummyRecord)
			recordList.AddRecord(dummyRecord)

			Expect(len(recordList.Records)).To(Equal(1))
		})
	})

	Context("RecordList.GetAllRecords", func() {
		It("Should return reference to the records in RecordList", func() {
			var (
				first  = CreateDummyRecord("first")
				second = CreateDummyRecord("second")
				third  = CreateDummyRecord("third")
			)
			recordList := model.NewRecordList()
			recordList.AddRecord(first)
			recordList.AddRecord(second)
			recordList.AddRecord(third)
			records := recordList.GetAllRecords()
			(*records)[0].Done()

			Expect(recordList.Records[0].Status).To(Equal(model.Done))
		})
	})

})
