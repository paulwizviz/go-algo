package roman

import "fmt"

const (
	I = 1
	V = 5
	X = 10
	L = 50
	C = 100
	D = 500
	M = 1000
)

var (
	Error = fmt.Errorf("invalid input string")
)

func SymbolsToValue(input string) (int, error) {

	symbolicValues, err := stringToIntSlice(input)
	if err != nil {
		return 0, err
	}

	accum := 0
	for index := 0; index <= len(symbolicValues)-1; index++ {
		currentValue := symbolicValues[index]
		var nextValue int
		if index+1 <= len(symbolicValues)-1 {
			nextValue = symbolicValues[index+1]
		}
		if currentValue >= nextValue {
			accum = accum + currentValue
		} else {
			compoundValue := nextValue - currentValue
			accum = accum + compoundValue
			index++
		}
	}
	return accum, nil
}

func stringToIntSlice(input string) ([]int, error) {

	var symbolicValues []int
	for _, symbol := range input {
		switch symbol {
		case 'I':
			symbolicValues = append(symbolicValues, I)
		case 'V':
			symbolicValues = append(symbolicValues, V)
		case 'X':
			symbolicValues = append(symbolicValues, X)
		case 'L':
			symbolicValues = append(symbolicValues, L)
		case 'C':
			symbolicValues = append(symbolicValues, C)
		case 'D':
			symbolicValues = append(symbolicValues, D)
		case 'M':
			symbolicValues = append(symbolicValues, M)
		default:
			return nil, fmt.Errorf("invalid input string")
		}
	}
	return symbolicValues, nil
}
