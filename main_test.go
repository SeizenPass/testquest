package main

import (
	"bytes"
	"github.com/SeizenPass/testquest/fast"
	"io/ioutil"
	"testing"
)

func init() {
	OriginalSolution(ioutil.Discard)
	fast.OurSolution(ioutil.Discard)
}

func TestMain(t *testing.T) {
	originalOut := new(bytes.Buffer)
	OriginalSolution(originalOut)
	originalResult := originalOut.String()

	fastOut := new(bytes.Buffer)
	fast.OurSolution(fastOut)
	fastResult := fastOut.String()

	if originalResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, originalResult)
	}
}
