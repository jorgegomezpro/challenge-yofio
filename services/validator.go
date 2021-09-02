package services

import (
	"fmt"
)

func getrequestError(value int32, minCreditType int32) error {
	if value < minCreditType {
		return fmt.Errorf("value must be greater than %v", minCreditType)
	}

	if (value % 100) != 0 {
		return fmt.Errorf("value must be a multiple of %v", 100)
	}

	return nil
}
