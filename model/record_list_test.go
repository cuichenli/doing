package model_test

import (
	"sort"

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

			Expect(len(recordList.DoingRecords)).To(Equal(1))
			Expect(recordList.DoingRecords[0].Detail).To(Equal(detail))
			Expect(recordList.DoingRecords[0].Status).To(Equal(model.Doing))
		})

		It("Add two records to the recordList", func() {
			const detail1 = "One Task"
			const detail2 = "Another Task"
			dummyRecord := CreateDummyRecord(detail1)
			dummyRecords = append(dummyRecords, dummyRecord)
			dummyRecord = CreateDummyRecord(detail2)
			dummyRecords = append(dummyRecords, dummyRecord)
			for _, record := range dummyRecords {
				recordList.AddRecord(record)
			}

			Expect(len(recordList.DoingRecords)).To(Equal(2))
			Expect(recordList.DoingRecords[0].Detail).To(Equal(detail1))
			Expect(recordList.DoingRecords[0].Status).To(Equal(model.Doing))
			Expect(recordList.DoingRecords[1].Detail).To(Equal(detail2))
			Expect(recordList.DoingRecords[1].Status).To(Equal(model.Doing))
		})

		It("Add one done record to the recordList", func() {
			const detail1 = "One Task"
			dummyRecord := CreateDummyRecord(detail1)
			dummyRecord.Done()
			dummyRecords = append(dummyRecords, dummyRecord)
			for _, record := range dummyRecords {
				recordList.AddRecord(record)
			}

			Expect(len(recordList.DoneRecords)).To(Equal(1))
			Expect(recordList.DoneRecords[0].Detail).To(Equal(detail1))
			Expect(recordList.DoneRecords[0].Status).To(Equal(model.Done))
		})
	})

	Context("RecordList.MarkItemDone", func() {
		const (
			notFinished    = "Not Finished"
			finished       = "Finshed"
			lastOne        = "Last One"
			secondFinished = "Second Finished"
		)
		BeforeEach(func() {

			dummyRecords = []model.Record{CreateDummyRecord(finished), CreateDummyRecord(notFinished), CreateDummyRecord(lastOne), CreateDummyRecord(secondFinished)}
			recordList = model.NewRecordList()
			for _, record := range dummyRecords {
				recordList.AddRecord(record)
			}
		})

		It("Mark one record as finished ", func() {
			record, _ := recordList.MarkItemDone(model.RecordID{ID: 0})
			Expect(record.Detail).To(Equal(finished))
			doingKeys := make([]int, 0)
			for k := range recordList.DoingRecords {
				doingKeys = append(doingKeys, k)
			}
			Expect(len(recordList.DoneRecords)).To(Equal(1))
			Expect(len(recordList.DoingRecords)).To(Equal(3))
			Expect(recordList.DoingRecords)
			Expect(recordList.DoneRecords[0].Detail).To(Equal(finished))
			sort.Ints(doingKeys)
			Expect(doingKeys).To(Equal([]int{1, 2, 3}))
		})

		It("Mark two records as finished ", func() {
			record, _ := recordList.MarkItemDone(model.RecordID{ID: 0})
			Expect(record.Detail).To(Equal(finished))
			record, _ = recordList.MarkItemDone(model.RecordID{ID: 3})
			Expect(record.Detail).To(Equal(secondFinished))
			doingKeys := make([]int, 0)
			for k := range recordList.DoingRecords {
				doingKeys = append(doingKeys, k)
			}
			Expect(len(recordList.DoneRecords)).To(Equal(2))
			Expect(len(recordList.DoingRecords)).To(Equal(2))
			Expect(recordList.DoingRecords)
			Expect(recordList.DoneRecords[0].Detail).To(Equal(finished))
			Expect(recordList.DoneRecords[3].Detail).To(Equal(secondFinished))
			sort.Ints(doingKeys)
			Expect(doingKeys).To(Equal([]int{1, 2}))
		})

		It("Mark one item not exist", func() {
			_, err := recordList.MarkItemDone(model.RecordID{ID: 5})
			Expect(err.Error()).To(Equal("The provided record ID 5 is not found in doing record list"))

		})
	})

	Context("RecordList.GetAllRecords", func() {
		It("Should return all the records sorted", func() {
			var (
				first  = CreateDummyRecord("first")
				second = CreateDummyRecord("second")
				third  = CreateDummyRecord("third")
			)
			recordList := model.NewRecordList()
			recordList.AddRecord(first)
			recordList.AddRecord(second)
			recordList.AddRecord(third)
			recordList.MarkItemDone(model.RecordID{ID: 2})

			third.Done()
			result := recordList.GetAllRecords()
			Expect(result[0]).To(Equal(first))
			Expect(result[1]).To(Equal(second))
			Expect(result[2]).To(Equal(third))
			Expect(len(result)).To(Equal(3))
		})
	})

})
