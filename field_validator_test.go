package validator

import (
	"net/url"
	"testing"
)

func TestFieldValidator_run(t *testing.T) {
	type val struct {
		rule, valid, invalid string
	}

	var vals = []val{
		{"required", "some input", ""},
		{"regex:[a-z]", "saddam", "123"},
		{"len:2", "ab", "1234"},
		{"min:2", "ab", "x"},
		{"alpha", "ab", "123"},
		{"alpha_dash", "abc_Abc-191", "*&^@#"},
		{"alpha_num", "abcAbc191", "*&^@#"},
		{"bool", "true", "55a"},
		{"between:2,5", "ab1", "abc123"},
		{"credit_card", "4896644531043572", "xxxxxxxxxxxxxxxx"},
		{"coordinate", "30.297018,-78.486328", "x,y"},
		{"css_color", "#00aaff", "11"},
		{"digits:10", "1234567890", "11111"},
		{"digits_between:2,5", "1234", "123456"},
		{"date", "2016-10-16", "201610112"},
		{"date:dd-mm-yyyy", "16-10-2016", "201610112"},
		{"email", "john@mail.com", "invalidEmail"},
		{"float", "99.9999", "91a.0"},
		{"in:admin,manager", "admin", "invalid role"},
		{"not_in:admin,manager", "xxx", "admin"},
		{"ip", "10.255.255.255", "1.1.1.1.1"},
		{"ip_v4", "10.255.255.255", "1.1.1.a.1"},
		{"ip_v6", "21DA:D3:0:2F3B:2AA:FF:FE28:9C5A", "a.1.1.a.1"},
		{"json", `{"name":"john", "age": 25}`, `{"invalid json"}`},
		{"lat", "30.297018", "88888"},
		{"lon", "-78.486328", "ppo000"},
		{"max:4", "john", "john doe"},
		{"numeric", "12376", "xyz"},
		{"numeric_between:18,60", "28", "10"},
		{"url", "facebook.com", "googlecom"},
		{"uuid", "ee7cf0a0-1922-401b-a1ae-6ec9261484c0", "39888f87-fb62-5988-a425-b2ea63f5b81e"},
		{"uuid_v3", "a987fbc9-4bed-3078-cf07-9141ba07c9f3", "ee7cf0a0-1922-401b-a1ae-6ec9261484c0"},
		{"uuid_v4", "df7cca36-3d7a-40f4-8f06-ae03cc22f045", "b987fbc9-4bed-3078-cf07-9141ba07c9f3"},
		{"uuid_v5", "39888f87-fb62-5988-a425-b2ea63f5b81e", "ee7cf0a0-1922-401b-a1ae-6ec9261484c0"},
	}

	emptyErrBag := url.Values{}
	errBag := url.Values{}
	for _, val := range vals {
		fv := fieldValidator{field: val.rule, value: val.valid, rule: val.rule, errsBag: emptyErrBag}
		fv.run()
	}
	if len(emptyErrBag) > 0 {
		t.Error("fieldValidator.run() failed for valid inputs!")
	}
	for _, val := range vals {
		fv := fieldValidator{field: val.rule, value: val.invalid, rule: val.rule, errsBag: errBag}
		fv.run()
	}
	if len(vals) != len(errBag) {
		t.Error("fieldValidator.run() failed for invalid inputs!")
	}
}

func TestFieldValidator_required(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "Jane", rule: "required", errsBag: _errsBag}
	fv.Required()
	if len(_errsBag) > 0 {
		t.Error("required failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "", rule: "required", errsBag: _errsBag}
	fvr.Required()
	if len(_errsBag) != 1 {
		t.Error("reverse required failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "", rule: "required", message: "This field is required!", errsBag: _errsBag}
	fvWithMsg.Required()
	if len(_errsBag) != 1 {
		t.Error("required with custom message failed!")
	}
}

func TestFieldValidator_regex(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "Jane", rule: "regex:^[a-zA-Z]+$", errsBag: _errsBag}
	fv.Regex()
	if len(_errsBag) > 0 {
		t.Error("regex failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "*)9#", rule: "regex:^[a-zA-Z]+$", errsBag: _errsBag}
	fvr.Regex()
	if len(_errsBag) != 1 {
		t.Error("reverse regex failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "*)9#", rule: "regex:^[a-zA-Z]+$", message: "This field failed", errsBag: _errsBag}
	fvWithMsg.Regex()
	if len(_errsBag) != 1 {
		t.Error("regex with custom failed!")
	}
}

func TestFieldValidator_alpha(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "Jane", rule: "alpha", errsBag: _errsBag}
	fv.Alpha()
	if len(_errsBag) > 0 {
		t.Error("alpha failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "Jane*&#", rule: "alpha", errsBag: _errsBag}
	fvr.Alpha()
	if len(_errsBag) != 1 {
		t.Error("reverse alpha failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "Jane*&#", rule: "alpha", message: "This field must contain letter", errsBag: _errsBag}
	fvWithMsg.Alpha()
	if len(_errsBag) != 1 {
		t.Error("alpha with custom message failed!")
	}
}

func TestFieldValidator_alphaDash(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "Mrs_Jane-Doe-321", rule: "alpha_dash", errsBag: _errsBag}
	fv.AlphaDash()
	if len(_errsBag) > 0 {
		t.Error("alpha_dash failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "Jane*&#", rule: "alpha_dash", errsBag: _errsBag}
	fvr.AlphaDash()
	if len(_errsBag) != 1 {
		t.Error("reverse alpha_dash failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "Jane*&#", rule: "alpha_dash", message: "This field may contain letters, dash and underscore", errsBag: _errsBag}
	fvWithMsg.AlphaDash()
	if len(_errsBag) != 1 {
		t.Error("alpha_dash with custom message failed!")
	}
}

func TestFieldValidator_alphaNumeric(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "Jane321", rule: "alpha_num", errsBag: _errsBag}
	fv.AlphaNumeric()
	if len(_errsBag) > 0 {
		t.Error("alpha numeric failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "jane*091$#", rule: "alpha_num", errsBag: _errsBag}
	fvr.AlphaNumeric()
	if len(_errsBag) != 1 {
		t.Error("reverse alpha numeric failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "jane*091$#", rule: "alpha_num", message: "This field may contain letters and numbers", errsBag: _errsBag}
	fvWithMsg.AlphaNumeric()
	if len(_errsBag) != 1 {
		t.Error("alpha numeric with custom message failed!")
	}
}

func TestFieldValidator_boolean(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "true", rule: "bool", errsBag: _errsBag}
	fv.Boolean()
	if len(_errsBag) > 0 {
		t.Error("boolean failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "invalid_boolean", rule: "bool", errsBag: _errsBag}
	fvr.Boolean()
	if len(_errsBag) != 1 {
		t.Error("reverse boolean failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "invalid_boolean", rule: "bool", message: "This field may contain 0, 1, false and true", errsBag: _errsBag}
	fvWithMsg.Boolean()
	if len(_errsBag) != 1 {
		t.Error("boolean with custom message failed!")
	}
}

func TestFieldValidator_creditCard(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "4896644531043572", rule: "credit_card", errsBag: _errsBag}
	fv.CreditCard()
	if len(_errsBag) > 0 {
		t.Error("creditCard failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "z8966445310435xy", rule: "credit_card", errsBag: _errsBag}
	fvr.CreditCard()
	if len(_errsBag) != 1 {
		t.Error("reverse creditCard failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "z8966445310435xy", rule: "credit_card", message: "This field must be valid credit card", errsBag: _errsBag}
	fvWithMsg.CreditCard()
	if len(_errsBag) != 1 {
		t.Error("creditCard with custom message failed!")
	}
}

func TestFieldValidator_coordinate(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "30.297018,-78.486328", rule: "coordinate", errsBag: _errsBag}
	fv.Coordinate()
	if len(_errsBag) > 0 {
		t.Error("coordinate failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "0, 887", rule: "coordinate", errsBag: _errsBag}
	fvr.Coordinate()
	if len(_errsBag) != 1 {
		t.Error("reverse coordinate failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "0, 887", rule: "coordinate", message: "This field must contain valid coordinate", errsBag: _errsBag}
	fvWithMsg.Coordinate()
	if len(_errsBag) != 1 {
		t.Error("coordinate with custom message failed!")
	}
}

func TestFieldValidator_validateCSSColor(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "#00aaff", rule: "color", errsBag: _errsBag}
	fv.ValidateCSSColor()
	if len(_errsBag) > 0 {
		t.Error("validateCSSColor failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "#0001", rule: "color", errsBag: _errsBag}
	fvr.ValidateCSSColor()
	if len(_errsBag) != 1 {
		t.Error("reverse validateCSSColor failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "#0001", rule: "color", message: "This must contain valid CSS color", errsBag: _errsBag}
	fvWithMsg.ValidateCSSColor()
	if len(_errsBag) != 1 {
		t.Error("validateCSSColor with custom message failed!")
	}
}

func TestFieldValidator_validateJSON(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: _validJSONString, rule: "json", errsBag: _errsBag}
	fv.ValidateJSON()
	if len(_errsBag) > 0 {
		t.Error("validateJSON failed!")
	}

	fvr := fieldValidator{field: "field_name", value: _invalidJSONString, rule: "json", errsBag: _errsBag}
	fvr.ValidateJSON()
	if len(_errsBag) != 1 {
		t.Error("reverse validateJSON failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: _invalidJSONString, rule: "json", message: "This field must contain valid JSON string", errsBag: _errsBag}
	fvWithMsg.ValidateJSON()
	if len(_errsBag) != 1 {
		t.Error("validateJSON with custom message failed!")
	}
}

func TestFieldValidator_length(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "Jhon", rule: "len:4", errsBag: _errsBag}
	fv.Length()
	if len(_errsBag) > 0 {
		t.Error("length failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "jane", rule: "len:5", errsBag: _errsBag}
	fvr.Length()
	if len(_errsBag) != 1 {
		t.Error("reverse length failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "jane", rule: "len:5", message: "This field must contain chars with fixed length", errsBag: _errsBag}
	fvWithMsg.Length()
	if len(_errsBag) != 1 {
		t.Error("length with custom message failed!")
	}
}

func TestFieldValidator_min(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "john doe", rule: "min:4", errsBag: _errsBag}
	fv.Min()
	if len(_errsBag) > 0 {
		t.Error("min failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "jane", rule: "min:5", errsBag: _errsBag}
	fvr.Min()
	if len(_errsBag) != 1 {
		t.Error("reverse min failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "jane", rule: "min:5", message: "This field must contain chars greater than min value", errsBag: _errsBag}
	fvWithMsg.Min()
	if len(_errsBag) != 1 {
		t.Error("min with custom message failed!")
	}
}

func TestFieldValidator_max(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "john", rule: "max:4", errsBag: _errsBag}
	fv.Max()
	if len(_errsBag) > 0 {
		t.Error("max failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "jane", rule: "max:3", errsBag: _errsBag}
	fvr.Max()
	if len(_errsBag) != 1 {
		t.Error("reverse max failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "jane", rule: "max:3", message: "This field must not contain chars length greater than max value", errsBag: _errsBag}
	fvWithMsg.Max()
	if len(_errsBag) != 1 {
		t.Error("max with custom failed!")
	}
}

func TestFieldValidator_between(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "john", rule: "between:3,5", errsBag: _errsBag}
	fv.Between()
	if len(_errsBag) > 0 {
		t.Error("between failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "jane", rule: "between:5,8", errsBag: _errsBag}
	fvr.Between()
	if len(_errsBag) != 1 {
		t.Error("reverse between failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "jane", rule: "between:5,8", message: "This field must contain chars length between provided value", errsBag: _errsBag}
	fvWithMsg.Between()
	if len(_errsBag) != 1 {
		t.Error("between with custom message failed!")
	}
}

func TestFieldValidator_numeric(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "1231", rule: "numeric", errsBag: _errsBag}
	fv.Numeric()
	if len(_errsBag) > 0 {
		t.Error("numeric failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "1@3@#*&^", rule: "numeric", errsBag: _errsBag}
	fvr.Numeric()
	if len(_errsBag) != 1 {
		t.Error("reverse numeric failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "1@3@#*&^", rule: "numeric", message: "This field must contain numeric value", errsBag: _errsBag}
	fvWithMsg.Numeric()
	if len(_errsBag) != 1 {
		t.Error("numeric with custom message failed failed!")
	}
}

func TestFieldValidator_numericBetween(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "4", rule: "numeric_between:3,5", errsBag: _errsBag}
	fv.NumericBetween()
	if len(_errsBag) > 0 {
		t.Error("numericBetween failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "1", rule: "numeric_between:3,6", errsBag: _errsBag}
	fvr.NumericBetween()
	if len(_errsBag) != 1 {
		t.Error("reverse numericBetween failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "1", rule: "numeric_between:3,6", message: "This field must contain numeric value between provided range", errsBag: _errsBag}
	fvWithMsg.NumericBetween()
	if len(_errsBag) != 1 {
		t.Error("numericBetween with custom message failed!")
	}
}

func TestFieldValidator_digits(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "8083", rule: "digits:4", errsBag: _errsBag}
	fv.Digits()
	if len(_errsBag) > 0 {
		t.Error("digits failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "99", rule: "digits:3", errsBag: _errsBag}
	fvr.Digits()
	if len(_errsBag) != 1 {
		t.Error("reverse digits failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "99", rule: "digits:3", message: "This field must contain digits of provided length", errsBag: _errsBag}
	fvWithMsg.Digits()
	if len(_errsBag) != 1 {
		t.Error("digits with custom message failed!")
	}
}

func TestFieldValidator_digitsBetween(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "8083", rule: "digits_between:4,6", errsBag: _errsBag}
	fv.DigitsBetween()
	if len(_errsBag) > 0 {
		t.Error("digitsBetween failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "99", rule: "digits_between:3,4", errsBag: _errsBag}
	fvr.DigitsBetween()
	if len(_errsBag) != 1 {
		t.Error("reverse digitsBetween failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "99", rule: "digits_between:3,4", message: "This field must contain digits with in provided range", errsBag: _errsBag}
	fvWithMsg.DigitsBetween()
	if len(_errsBag) != 1 {
		t.Error("digitsBetween with custom message failed!")
	}
}

func TestFieldValidator_email(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "john@mail.com", rule: "email", errsBag: _errsBag}
	fv.Email()
	if len(_errsBag) > 0 {
		t.Error("email failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "yaho.com", rule: "email", errsBag: _errsBag}
	fvr.Email()
	if len(_errsBag) != 1 {
		t.Error("reverse email failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "yaho.com", rule: "email", message: "This field must contain valid email address", errsBag: _errsBag}
	fvWithMsg.Email()
	if len(_errsBag) != 1 {
		t.Error("email with custom message failed!")
	}
}

func TestFieldValidator_date(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "2016-10-11", rule: "date", errsBag: _errsBag}
	fv.Date()
	if len(_errsBag) > 0 {
		t.Error("date failed!")
	}

	fvDDMMYYYY := fieldValidator{field: "field_name", value: "11-10-2016", rule: "date:dd-mm-yyyy", errsBag: _errsBag}
	fvDDMMYYYY.Date()
	if len(_errsBag) > 0 {
		t.Error("Date with format dd-mm-yyyy failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "yyyy-mm-dd", rule: "date", errsBag: _errsBag}
	fvr.Date()
	if len(_errsBag) != 1 {
		t.Error("reverse date failed!")
	}

	fvrDDMMYYYY := fieldValidator{field: "field_name", value: "yyyy-mm-dd", rule: "date:dd-mm-yyyy", errsBag: _errsBag}
	fvrDDMMYYYY.Date()
	if len(_errsBag) != 1 {
		t.Error("reverse date with dd-mm-yyyy format failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "yyyy-mm-dd", rule: "date", message: "This field must contain valid date in yyyy-mm-dd format", errsBag: _errsBag}
	fvWithMsg.Date()
	if len(_errsBag) != 1 {
		t.Error("date with custom message failed!")
	}

	fvDDMMYYYYWithMsg := fieldValidator{field: "field_name", value: "yyyy-mm-dd", rule: "date:dd-mm-yyyy", message: "This field must contain valid date in dd-mm-yyyy format", errsBag: _errsBag}
	fvDDMMYYYYWithMsg.Date()
	if len(_errsBag) != 1 {
		t.Error("Date with format dd-mm-yyyy and custom message failed!")
	}
}

func TestFieldValidator_validateFloat(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "99.00998", rule: "float", errsBag: _errsBag}
	fv.ValidateFloat()
	if len(_errsBag) > 0 {
		t.Error("validateFloat failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "9920x", rule: "float", errsBag: _errsBag}
	fvr.ValidateFloat()
	if len(_errsBag) != 1 {
		t.Error("reverse validateFloat failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "9920x", rule: "float", message: "This field must contain float value", errsBag: _errsBag}
	fvWithMsg.ValidateFloat()
	if len(_errsBag) != 1 {
		t.Error("validateFloat with custom message failed!")
	}
}

func TestFieldValidator_latitude(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "30.297018", rule: "lat", errsBag: _errsBag}
	fv.Latitude()
	if len(_errsBag) > 0 {
		t.Error("latitude failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "30.29701z", rule: "lat", errsBag: _errsBag}
	fvr.Latitude()
	if len(_errsBag) != 1 {
		t.Error("reverse latitude failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "30.29701z", rule: "lat", message: "This field must contain valid latitude", errsBag: _errsBag}
	fvWithMsg.Latitude()
	if len(_errsBag) != 1 {
		t.Error("latitude with custom message failed!")
	}
}

func TestFieldValidator_longitude(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "-78.486328", rule: "lon", errsBag: _errsBag}
	fv.Longitude()
	if len(_errsBag) > 0 {
		t.Error("longitude failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "-78.486328xyz", rule: "lon", errsBag: _errsBag}
	fvr.Longitude()
	if len(_errsBag) != 1 {
		t.Error("reverse longitude failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "-78.486328xyz", rule: "lon", message: "This field must contain valid longitude", errsBag: _errsBag}
	fvWithMsg.Longitude()
	if len(_errsBag) != 1 {
		t.Error("longitude with custom message failed!")
	}
}

func TestFieldValidator_in(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "admin", rule: "in:admin,manager,supervisor", errsBag: _errsBag}
	fv.In()
	if len(_errsBag) > 0 {
		t.Error("in failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "user", rule: "in:admin,manager,supervisor", errsBag: _errsBag}
	fvr.In()
	if len(_errsBag) != 1 {
		t.Error("reverse in failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "user", rule: "in:admin,manager,supervisor", message: "This field must contain one of the provided value", errsBag: _errsBag}
	fvWithMsg.In()
	if len(_errsBag) != 1 {
		t.Error("in with custom message failed!")
	}
}

func TestFieldValidator_notIn(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "user", rule: "not_in:admin,manager,supervisor", errsBag: _errsBag}
	fv.NotIn()
	if len(_errsBag) > 0 {
		t.Error("notIn failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "admin", rule: "not_in:admin,manager,supervisor", errsBag: _errsBag}
	fvr.NotIn()
	if len(_errsBag) != 1 {
		t.Error("reverse notIn failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "admin", rule: "not_in:admin,manager,supervisor", message: "This field must not contain any of the provided value", errsBag: _errsBag}
	fvWithMsg.NotIn()
	if len(_errsBag) != 1 {
		t.Error("notIn with custom message failed!")
	}
}

func TestFieldValidator_ip(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "10.255.255.255", rule: "ip", errsBag: _errsBag}
	fv.IP()
	if len(_errsBag) > 0 {
		t.Error("ip failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "10.255.255.25z", rule: "ip", errsBag: _errsBag}
	fvr.IP()
	if len(_errsBag) != 1 {
		t.Error("reverse ip failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "10.255.255.25z", rule: "ip", message: "This field must contain valid IP address", errsBag: _errsBag}
	fvWithMsg.IP()
	if len(_errsBag) != 1 {
		t.Error("ip with custom message failed!")
	}
}

func TestFieldValidator_ipv4(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "10.255.255.255", rule: "ip_v4", errsBag: _errsBag}
	fv.IPv4()
	if len(_errsBag) > 0 {
		t.Error("ip v4 failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "10.255.255.25z", rule: "ip_v4", errsBag: _errsBag}
	fvr.IPv4()
	if len(_errsBag) != 1 {
		t.Error("reverse ip v4 failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "10.255.255.25z", rule: "ip_v4", message: "This field must contain valid IP V4 address", errsBag: _errsBag}
	fvWithMsg.IPv4()
	if len(_errsBag) != 1 {
		t.Error("ip v4 with custom message failed!")
	}
}

func TestFieldValidator_ipv6(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "1200:0000:AB00:1234:0000:2552:7777:1313", rule: "ip_v6", errsBag: _errsBag}
	fv.IPv6()
	if len(_errsBag) > 0 {
		t.Error("ip v6 failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "10.255.255.255", rule: "ip_v6", errsBag: _errsBag}
	fvr.IPv6()
	if len(_errsBag) != 1 {
		t.Error("reverse ip v6 failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "10.255.255.255", rule: "ip_v6", message: "This field must contain valid IP V6 address", errsBag: _errsBag}
	fvWithMsg.IPv6()
	if len(_errsBag) != 1 {
		t.Error("ip v6 with custom message failed!")
	}
}

func TestFieldValidator_validateURL(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "google.com", rule: "url", errsBag: _errsBag}
	fv.ValidateURL()
	if len(_errsBag) > 0 {
		t.Error("validateURL failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "googlecom", rule: "url", errsBag: _errsBag}
	fvr.ValidateURL()
	if len(_errsBag) != 1 {
		t.Error("invalid validateURL failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "googlecom", rule: "url", message: "This field must contain valid URL", errsBag: _errsBag}
	fvWithMsg.ValidateURL()
	if len(_errsBag) != 1 {
		t.Error("validateURL with custom message failed!")
	}
}

func TestFieldValidator_uuid(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "ee7cf0a0-1922-401b-a1ae-6ec9261484c0", rule: "uuid", errsBag: _errsBag}
	fv.UUID()
	if len(_errsBag) > 0 {
		t.Error("uuid failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "39888f87-fb62-5988-a425-b2ea63f5b81e", rule: "uuid", errsBag: _errsBag}
	fvr.UUID()
	if len(_errsBag) != 1 {
		t.Error("invalid uuid failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "39888f87-fb62-5988-a425-b2ea63f5b81e", rule: "uuid", message: "This field must contain valid UUID", errsBag: _errsBag}
	fvWithMsg.UUID()
	if len(_errsBag) != 1 {
		t.Error("invalid uuid with custom message failed!")
	}
}

func TestFieldValidator_uuid3(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "a987fbc9-4bed-3078-cf07-9141ba07c9f3", rule: "uuid_v3", errsBag: _errsBag}
	fv.UUID3()
	if len(_errsBag) > 0 {
		t.Error("uuid v3 failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "ee7cf0a0-1922-401b-a1ae-6ec9261484c0", rule: "uuid_v3", errsBag: _errsBag}
	fvr.UUID3()
	if len(_errsBag) != 1 {
		t.Error("invalid uuid v3 failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "ee7cf0a0-1922-401b-a1ae-6ec9261484c0", rule: "uuid_v3", message: "This field must contain valid UUID of V3", errsBag: _errsBag}
	fvWithMsg.UUID3()
	if len(_errsBag) != 1 {
		t.Error("invalid uuid v3 with custom message failed!")
	}
}

func TestFieldValidator_uuid4(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "df7cca36-3d7a-40f4-8f06-ae03cc22f045", rule: "uuid_v4", errsBag: _errsBag}
	fv.UUID4()
	if len(_errsBag) > 0 {
		t.Error("uuid v4 failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "b987fbc9-4bed-3078-cf07-9141ba07c9f3", rule: "uuid_v4", errsBag: _errsBag}
	fvr.UUID4()
	if len(_errsBag) != 1 {
		t.Error("reverse uuid v4 failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "b987fbc9-4bed-3078-cf07-9141ba07c9f3", rule: "uuid_v4", message: "This field must contain valid UUID of V4", errsBag: _errsBag}
	fvWithMsg.UUID4()
	if len(_errsBag) != 1 {
		t.Error("reverse uuid v4 with custom message failed!")
	}
}

func TestFieldValidator_uuid5(t *testing.T) {
	_errsBag := url.Values{}
	fv := fieldValidator{field: "field_name", value: "39888f87-fb62-5988-a425-b2ea63f5b81e", rule: "uuid_v5", errsBag: _errsBag}
	fv.UUID5()
	if len(_errsBag) > 0 {
		t.Error("uuid v5 failed!")
	}

	fvr := fieldValidator{field: "field_name", value: "b987fbc9-4bed-3078-cf07-9141ba07c9f3", rule: "uuid_v5", errsBag: _errsBag}
	fvr.UUID5()
	if len(_errsBag) != 1 {
		t.Error("reverse uuid v5 failed!")
	}

	fvWithMsg := fieldValidator{field: "field_name", value: "b987fbc9-4bed-3078-cf07-9141ba07c9f3", rule: "uuid_v5", message: "This field must contain valid UUID of V5", errsBag: _errsBag}
	fvWithMsg.UUID5()
	if len(_errsBag) != 1 {
		t.Error("reverse uuid v5 with custom message failed!")
	}
}
