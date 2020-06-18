package model_test

import (
	"time"

	"github.com/cuichenli/doing/model"
)

var _ = Describe("Record", func() {
	var (
		record model.Record
		loc    *time.Location
	)
	Context("ToTaskPaper", func() {
		BeforeEach(func() {
			loc, _ = time.LoadLocation("UTC")
			record = model.Record{
				Detail:      "A simple task",
				Status:      model.Doing,
				CreatedTime: time.Date(2015, 04, 03, 12, 20, 07, 27, loc),
			}
		})

		It("Convert a basic record to task paper string", func() {
			result := record.ToTaskPaper()
			Expect(result).To(Equal("  - A simple task @created(2015-04-03T12:20:07Z)"))
		})

		It("Convert a done record to task paper string", func() {
			record.Done()
			result := record.ToTaskPaper()
			Expect(result).To(Equal("  - A simple task @created(2015-04-03T12:20:07Z) @done"))
		})
	})

	Context("ParseTags", func() {
		It("Convert 'done' to pair ('done, '')", func() {
			str := "done"
			name, value, err := model.ParseTags(str)
			Expect(name).To(Equal("done"))
			Expect(value).To(Equal(""))
			Expect(err).To(BeNil())
		})

		It("Convert 'created(123)' to pair ('created, '123')", func() {
			str := "created(123)"
			name, value, err := model.ParseTags(str)
			Expect(name).To(Equal("created"))
			Expect(value).To(Equal("123"))
			Expect(err).To(BeNil())
		})

		It("Return error when parsing 'error tag'", func() {
			str := "error tag"
			name, value, err := model.ParseTags(str)
			Expect(name).To(Equal(""))
			Expect(value).To(Equal(""))
			Expect(err).ToNot(BeNil())
		})
	})

	Context("FromTaskPaper", func() {
		It("Should parse task paper 'detail @created(2015-04-03T12:20:07Z)'", func() {
			record, err := model.FromTaskPaper("detail @created(2015-04-03T12:20:07Z)")
			Expect(record.Detail).To(Equal("detail"))
			Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).To(BeNil())
			Expect(record.Status).To(Equal(model.Doing))
		})

		It("Should parse task paper 'detail @done @created(2015-04-03T12:20:07Z)'", func() {
			record, err := model.FromTaskPaper("detail @done @created(2015-04-03T12:20:07Z)")
			Expect(record.Detail).To(Equal("detail"))
			Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).To(BeNil())
			Expect(record.Status).To(Equal(model.Done))
		})

		It("Should parse task paper 'detail @done but not done @created(2015-04-03T12:20:07Z)'", func() {
			record, err := model.FromTaskPaper("detail @done but not done @created(2015-04-03T12:20:07Z)")
			Expect(record.Detail).To(Equal("detail @done but not done"))
			Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).To(BeNil())
			Expect(record.Status).To(Equal(model.Doing))
		})

		It("Should fail to parse task paper 'detail @done but not done @created(2015-04-03T12:20:07Z) no tag'", func() {
			_, err := model.FromTaskPaper("detail @done but not done @created(2015-04-03T12:20:07Z) no tag")
			// Expect(record.Detail).To(Equal("detail @done but not done"))
			// Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).ToNot(BeNil())
			// Expect(record.Status).To(Equal(model.Doing))
		})
	})
})
