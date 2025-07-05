package enums

import (
	"encoding/json"
	"fmt"
)

type DataType string

const (
	String  DataType = "string"
	Number  DataType = "number"
	Date    DataType = "date"
	Boolean DataType = "bool"
)

func isValidDataType(dataType string) bool {
	switch dataType {
	case "string", "number", "date", "bool":
		return true
	}
	return false
}

func (r *DataType) UnmarshalJSON(data []byte) error {
	var input string
	if err := json.Unmarshal(data, &input); err != nil {
		return err
	}

	ok := isValidDataType(input)
	if !ok {
		return fmt.Errorf("invalid dataType: %s (must be String, Number, Date or Boolean)", input)
	}

	*r = DataType(input)

	return nil
}
