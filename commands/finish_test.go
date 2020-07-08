package commands

import (
	"testing"

	"github.com/cuichenli/doing/model"
	"github.com/onsi/gomega"
	. "github.com/onsi/gomega"
)

func TestFinishLastOneItem(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	records := model.NewRecordList()
	records.AddRecord(model.NewDoingRecord("record 1"))
	records.AddRecord(model.NewDoingRecord("record 2"))
	finishLastItems(&records, 1)
	g.Expect(records.Records[1].Detail).To(Equal("record 2"))
	g.Expect(records.Records[1].Status).To(Equal(model.Done))
	g.Expect(records.Records[0].Status).To(Equal(model.Doing))
}

func TestFinishLastTwoItems(t *testing.T) {
	g := gomega.NewGomegaWithT(t)
	records := model.NewRecordList()
	records.AddRecord(model.NewDoingRecord("record 1"))
	records.AddRecord(model.NewDoingRecord("record 2"))
	finishLastItems(&records, 2)
	g.Expect(records.Records[1].Detail).To(Equal("record 2"))
	g.Expect(records.Records[1].Status).To(Equal(model.Done))
	g.Expect(records.Records[0].Status).To(Equal(model.Done))
}
