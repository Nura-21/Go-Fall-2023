package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Year int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (y *Year) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d year", y)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (y *Year) UnmarshalJSON(jsonValue []byte) error {
	unquotedJSONValue, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	parts := strings.Split(unquotedJSONValue, " ")
	if len(parts) != 2 || parts[1] != "year" {
		return ErrInvalidRuntimeFormat
	}
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidRuntimeFormat
	}
	*y = Year(i)
	return nil
}
