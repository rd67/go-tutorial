package math

import (
	"testing"

	"github.com/rd67/go-packages/logger"
)

func TestAdd(t *testing.T) {
	total := Add(2,3)
	if total != 5 {
		t.Errorf("Summ is incorrect, Actual: %d, Expected: %d", total, 5);
	}

	logger.LogInfo("Running TestAdd")
}
