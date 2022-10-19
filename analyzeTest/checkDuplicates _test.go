package analysis

import (
	file "task/analysis"
	"testing"
)

func TestCheckDuplicates(t *testing.T) {

	file.Filename = "test.log"

	res, _ := file.CheckDuplicates()
	if res != true {
		t.Errorf("Expected true but found %v", res)
	}
}
