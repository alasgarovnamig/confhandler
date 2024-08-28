package confhandler

import (
	"fmt"
	"strconv"
)

// ParseValue parses the value to the expected type and returns it as an interface{}.
func ParseValue(value string, expectedType SupportedType) (interface{}, error) {
	switch expectedType {
	case StringType:
		return value, nil
	case IntType:
		return strconv.Atoi(value)
	case BoolType:
		return strconv.ParseBool(value)
	case FloatType:
		return strconv.ParseFloat(value, 64)
	default:
		return nil, fmt.Errorf("unsupported type: %s", expectedType)
	}
}
