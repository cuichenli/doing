package model_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

var Describe = ginkgo.Describe
var BeforeEach = ginkgo.BeforeEach
var It = ginkgo.It
var Expect = gomega.Expect
var Equal = gomega.Equal

func TestModel(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Model Suite")
}
