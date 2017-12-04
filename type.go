package govalidator

import (
	"encoding/json"
)

// Int describes a custom type of built-in int data type
type Int struct {
	Value int  `json:"value"`
	IsSet bool `json:"isSet"`
}

// UnmarshalJSON ...
func (i *Int) UnmarshalJSON(data []byte) error {
	i.IsSet = true
	var temp int
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	i.Value = temp
	return nil
}

// MarshalJSON ...
func (i *Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Value)
}
