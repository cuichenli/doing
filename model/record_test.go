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
				Title:       "A simple task",
				Status:      model.Doing,
				CreatedTime: time.Date(2015, 04, 03, 12, 20, 07, 27, loc),
				Tag:         make(model.RecordTag),
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

		It("Convert a record with one tag to task paper string", func() {
			record.Done()
			record.AddTag("name", "value")
			result := record.ToTaskPaper()
			Expect(result).To(Equal("  - A simple task @created(2015-04-03T12:20:07Z) @done @name(value)"))
		})

		It("Convert a record with tags to task paper string", func() {
			record.Done()
			record.AddTag("name", "value")
			record.AddTag("tag1", "")
			result := record.ToTaskPaper()
			Expect(result).To(Equal("  - A simple task @created(2015-04-03T12:20:07Z) @done @name(value) @tag1"))
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
		It("Should parse task paper 'title @created(2015-04-03T12:20:07Z)'", func() {
			record, err := model.FromTaskPaper("title @created(2015-04-03T12:20:07Z)")
			Expect(record.Title).To(Equal("title"))
			Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).To(BeNil())
			Expect(record.Status).To(Equal(model.Doing))
		})

		It("Should parse task paper 'title @done @created(2015-04-03T12:20:07Z)'", func() {
			record, err := model.FromTaskPaper("title @done @created(2015-04-03T12:20:07Z)")
			Expect(record.Title).To(Equal("title"))
			Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).To(BeNil())
			Expect(record.Status).To(Equal(model.Done))
		})

		It("Should parse task paper 'title @done but not done @created(2015-04-03T12:20:07Z)'", func() {
			record, err := model.FromTaskPaper("title @done but not done @created(2015-04-03T12:20:07Z)")
			Expect(record.Title).To(Equal("title @done but not done"))
			Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).To(BeNil())
			Expect(record.Status).To(Equal(model.Doing))
		})

		It("Should parse task paper 'title @done but not done @created(2015-04-03T12:20:07Z) @tag(value)'", func() {
			record, err := model.FromTaskPaper("title @done but not done @created(2015-04-03T12:20:07Z) @tag(value)")
			Expect(record.Title).To(Equal("title @done but not done"))
			Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).To(BeNil())
			Expect(record.Status).To(Equal(model.Doing))
			Expect(len(record.Tag)).To(Equal(1))
			Expect(record.Tag["tag"]).To(Equal("value"))
		})

		It("Should parse task paper 'title @done but not done @created(2015-04-03T12:20:07Z) @tag(value) @tag2(value2)'", func() {
			record, err := model.FromTaskPaper("title @done but not done @tag2(value2) @created(2015-04-03T12:20:07Z) @tag(value)")
			Expect(record.Title).To(Equal("title @done but not done"))
			Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).To(BeNil())
			Expect(record.Status).To(Equal(model.Doing))
			Expect(len(record.Tag)).To(Equal(2))
			Expect(record.Tag["tag"]).To(Equal("value"))
			Expect(record.Tag["tag2"]).To(Equal("value2"))
		})

		It("Should parse task paper 'title @done but not done @created(2015-04-03T12:20:07Z) @tag(value) @tag2(value2) @tag3'", func() {
			record, err := model.FromTaskPaper("title @done but not done @tag2(value2) @created(2015-04-03T12:20:07Z) @tag(value) @tag3")
			Expect(record.Title).To(Equal("title @done but not done"))
			Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).To(BeNil())
			Expect(record.Status).To(Equal(model.Doing))
			Expect(len(record.Tag)).To(Equal(3))
			Expect(record.Tag["tag"]).To(Equal("value"))
			Expect(record.Tag["tag2"]).To(Equal("value2"))
			Expect(record.Tag["tag3"]).To(Equal(""))
		})

		It("Should fail to parse task paper 'title @done but not done @created(2015-04-03T12:20:07Z) no tag'", func() {
			_, err := model.FromTaskPaper("title @done but not done @created(2015-04-03T12:20:07Z) no tag")
			// Expect(record.title).To(Equal("title @done but not done"))
			// Expect(record.CreatedTime).To(Equal(time.Date(2015, 04, 03, 12, 20, 07, 0, loc)))
			Expect(err).ToNot(BeNil())
			// Expect(record.Status).To(Equal(model.Doing))
		})
	})

	Context("Record.ToDisplayString", func() {
		It("Should create a string to be displayed", func() {
			record := CreateDummyRecord("A record")
			str := record.ToDisplayString()
			Expect(str).To(Equal("| 2015-04-03 12:20 | A record"))
		})

	})

	Context("Record.AddTagFromRawString", func() {
		It("Should add tag value pair to the record based on a string", func() {
			record := CreateDummyRecord("A record")
			record.AddTagFromRawString("tag=value")
			Expect(record.Tag["tag"]).To(Equal("value"))
		})

		It("Should add tag without value to the record based on a string", func() {
			record := CreateDummyRecord("A record")
			record.AddTagFromRawString("tag")
			Expect(record.Tag["tag"]).To(Equal(""))
		})
	})

	Context("Record.RemoveTagFromRawString", func() {
		It("Should remove tag to the record based on a string of tag value pair", func() {
			record := CreateDummyRecord("A record")
			record.Tag["tag"] = "v"
			record.RemoveTagFromRawString("tag=value")
			_, ok := record.Tag["tag"]
			Expect(ok).To(Equal(false))
		})

		It("Should remove tag to the record based on a string of only tag", func() {
			record := CreateDummyRecord("A record")
			record.Tag["exist"] = "yes"
			record.Tag["tag"] = "v"
			record.RemoveTagFromRawString("tag")
			_, ok := record.Tag["tag"]
			Expect(ok).To(Equal(false))
			Expect(record.Tag["exist"]).To(Equal("yes"))
		})
	})
})
