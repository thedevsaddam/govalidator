package govalidator

import "errors"

var (
	errStringToInt          = errors.New("govalidator: unable to parse string to integer")
	errStringToFloat        = errors.New("govalidator: unable to parse string to float")
	errRequireRules         = errors.New("govalidator: provide at least rules for Validate* method")
	errValidateArgsMismatch = errors.New("govalidator: provide at least *http.Request and rules for Validate method")
	errInvalidArgument      = errors.New("govalidator: invalid number of argument")
	errRequirePtr           = errors.New("govalidator: provide pointer to the data structure")
	errRequireData          = errors.New("govalidator: provide non-nil data structure for ValidateStruct method")
	errRequestNotAccepted   = errors.New("govalidator: cannot provide an *http.Request for ValidateStruct method")
)
