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
})
