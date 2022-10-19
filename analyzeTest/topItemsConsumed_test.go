package analysis

import (
	file "task/analysis"
	"testing"
)

func TestTopItemsConsumed(t *testing.T) {

	file.Filename = "test.log"

	res := file.TopItemsConsumed(1)
	if res[0][0] != file.FoodList["304"] {
		t.Errorf("Expected %v but found %v", file.FoodList["304"], res)
	}
}
