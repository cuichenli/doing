package model_test

import (
	model "github.com/cuichenli/doing/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Record", func() {
	var (
		record model.Record
	)

	BeforeEach(func() {
		record = model.Record{
			Detail:      "dafas",
			Status:      model.Doing,
			CreatedTime: 1231321,
		}
	})
})
