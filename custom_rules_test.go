package validator

import (
	"fmt"
	"net/url"
	"testing"
)

func Test_AddCustomRule(t *testing.T) {
	AddCustomRule("__x__", func(f string, v interface{}, r string) error {
		if v.(string) != "xyz" {
			return fmt.Errorf("The %s field must be xyz", f)
		}
		return nil
	})
	if len(customRules) <= 0 {
		t.Error("AddCustomRule failed to add new rule")
	}
}

func Test_AddCustomRule_panic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("AddCustomRule failed to panic")
		}
	}()
	AddCustomRule("__x__", func(f string, v interface{}, r string) error {
		if v.(string) != "xyz" {
			return fmt.Errorf("The %s field must be xyz", f)
		}
		return nil
	})
}

func Test_validateExtraRules(t *testing.T) {
	errsBag := url.Values{}
	validateCustomRules("f_field", "a", "__x__", errsBag)
	if len(errsBag) != 1 {
		t.Error("validateExtraRules failed")
	}
}
