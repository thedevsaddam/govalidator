package govalidator

import "encoding/json"

// Int describes a custom type of built-in int data type
type Int struct {
	Value int
	isSet bool
}

func (i *Int) UnmarshalJSON(data []byte) error {
	i.isSet = true
	var val int
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	i.Value = val
	return nil
}
