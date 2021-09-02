package services

import (
	"fmt"
	"reflect"

	"main.go/contracts"
)

type RecursiveAssign struct {
	creditTypesOrder []int32
	creditTypes      map[int32]bool
	minCreditType    int32
}

func NewRecursiveAssign(creditTypes ...int32) contracts.CreditAssigner {
	recursiveAssign := &RecursiveAssign{}
	recursiveAssign.creditTypes = make(map[int32]bool)
	recursiveAssign.creditTypesOrder = creditTypes

	for _, v := range creditTypes {
		if _, exists := recursiveAssign.creditTypes[v]; !exists {
			recursiveAssign.creditTypes[v] = true

			if recursiveAssign.minCreditType == 0 {
				recursiveAssign.minCreditType = v
				continue
			}

			if recursiveAssign.minCreditType > v {
				recursiveAssign.minCreditType = v
			}
		}
	}

	return recursiveAssign
}

func (r RecursiveAssign) Assign(investment int32) (int32, int32, int32, error) {
	assignments, err := r.Assignments(investment)
	if err != nil {
		return 0, 0, 0, err
	}

	if len(assignments) > 0 {
		assignment := assignments[0]
		values := make([]int32, len(r.creditTypesOrder))
		for i, credit := range r.creditTypesOrder {
			values[i] = assignment[credit]
		}
		return values[0], values[1], values[2], nil
	}

	return 0, 0, 0, fmt.Errorf("Investment: %v is not possible", investment)
}

func (r RecursiveAssign) Assignments(investment int32) ([]map[int32]int32, error) {
	combinations := []map[int32]int32{}
	if err := getrequestError(investment, r.minCreditType); err != nil {
		return combinations, err
	}

	r.setCombinations(investment, make(map[int32]int32), &combinations)
	if len(combinations) == 0 {
		return combinations, fmt.Errorf("The combination is not possible")
	}

	return combinations, nil
}

func (r RecursiveAssign) setCombinations(investment int32, combination map[int32]int32, combinations *[]map[int32]int32) {
	for credit, _ := range r.creditTypes {
		currentCombination := cloneCombination(combination)
		if v, exists := currentCombination[credit]; exists {
			currentCombination[credit] = (1 + v)
		} else {
			currentCombination[credit] = 1
		}

		isValid, err := isValidCombination(investment, currentCombination)
		if err != nil {
			currentCombination[credit] -= 1
			continue
		}

		if !isValid {
			r.setCombinations(investment, currentCombination, combinations)
			continue
		}

		if !existsCombination(currentCombination, combinations) {
			*combinations = append(*combinations, currentCombination)
		}
	}
}

func existsCombination(combination map[int32]int32, combinations *[]map[int32]int32) bool {
	for _, c := range *combinations {
		if reflect.DeepEqual(c, combination) {
			return true
		}
	}

	return false
}

func cloneCombination(origin map[int32]int32) map[int32]int32 {
	destination := make(map[int32]int32)
	for k, v := range origin {
		destination[k] = v
	}
	return destination
}

// isValidCombination validate if combination is valid to iterate to cumulative value
// Return values:
// bool  -> if the value is less than the investment
// error -> if the value is greater than the investment
func isValidCombination(investment int32, combination map[int32]int32) (bool, error) {
	total := int32(0)
	for k, v := range combination {
		total += (k * v)
		if total > investment {
			return false, fmt.Errorf("The combination is not possible")
		}
	}

	if total < investment {
		return false, nil
	}

	return true, nil
}
