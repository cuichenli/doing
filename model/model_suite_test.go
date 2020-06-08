package model_test

import (
	"testing"
	"time"

	"github.com/cuichenli/doing/model"
	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var Describe = ginkgo.Describe
var BeforeEach = ginkgo.BeforeEach
var It = ginkgo.It
var Expect = gomega.Expect
var Equal = gomega.Equal
var Context = ginkgo.Context

func TestModel(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Model Suite")
}

func CreateDummyRecord(detail string) model.Record {
	loc, _ := time.LoadLocation("UTC")
	dummyRecord := model.Record{
		Status:      model.Doing,
		CreatedTime: time.Date(2015, 04, 03, 12, 20, 07, 27, loc),
		Detail:      detail,
	}
	return dummyRecord
}
