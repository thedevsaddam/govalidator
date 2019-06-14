![govalidator](govalidator.png)

[![Build Status](https://travis-ci.org/thedevsaddam/govalidator.svg?branch=master)](https://travis-ci.org/thedevsaddam/govalidator)
[![Project status](https://img.shields.io/badge/version-1.9-green.svg)](https://github.com/thedevsaddam/govalidator/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/thedevsaddam/govalidator)](https://goreportcard.com/report/github.com/thedevsaddam/govalidator)
[![Coverage Status](https://coveralls.io/repos/github/thedevsaddam/govalidator/badge.svg?branch=master)](https://coveralls.io/github/thedevsaddam/govalidator?branch=master)
[![GoDoc](https://godoc.org/github.com/thedevsaddam/govalidator?status.svg)](https://godoc.org/github.com/thedevsaddam/govalidator)
[![License](https://img.shields.io/dub/l/vibe-d.svg)](https://github.com/thedevsaddam/govalidator/blob/dev/LICENSE.md)

Validate golang request data with simple rules. Highly inspired by Laravel's request validation.


### Installation

Install the package using
```go
$ go get github.com/thedevsaddam/govalidator
// or
$ go get gopkg.in/thedevsaddam/govalidator.v1
```

### Usage

To use the package import it in your `*.go` code
```go
import "github.com/thedevsaddam/govalidator"
// or
import "gopkg.in/thedevsaddam/govalidator.v1"
```

### Example

***Validate `form-data`, `x-www-form-urlencoded` and `query params`***

```go

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,8"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
		"phone":    []string{"digits:11"},
		"agree":    []string{"bool"},
		"dob":      []string{"date"},
	}

	messages := govalidator.MapData{
		"username": []string{"required:আপনাকে অবশ্যই ইউজারনেম দিতে হবে", "between:ইউজারনেম অবশ্যই ৩-৮ অক্ষর হতে হবে"},
		"phone":    []string{"digits:ফোন নাম্বার অবশ্যই ১১ নম্বারের হতে হবে"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}
	v := govalidator.New(opts)
	e := v.Validate()
	err := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port: 9000")
	http.ListenAndServe(":9000", nil)
}

```

Send request to the server using curl or postman: `curl GET "http://localhost:9000?web=&phone=&zip=&dob=&agree="`


***Response***
```json
{
    "validationError": {
        "agree": [
            "The agree may only contain boolean value, string or int 0, 1"
        ],
        "dob": [
            "The dob field must be a valid date format. e.g: yyyy-mm-dd, yyyy/mm/dd etc"
        ],
        "email": [
            "The email field is required",
            "The email field must be a valid email address"
        ],
        "phone": [
            "ফোন নাম্বার অবশ্যই ১১ নম্বারের হতে হবে"
        ],
        "username": [
            "আপনাকে অবশ্যই ইউজারনেম দিতে হবে",
            "ইউজারনেম অবশ্যই ৩-৮ অক্ষর হতে হবে"
        ],
        "web": [
            "The web field format is invalid"
        ]
    }
}
```

### More examples

***Validate file***

* [Validate file](doc/FILE_VALIDATION.md)

***Validate `application/json` or `text/plain` as raw body***

* [Validate JSON to simple struct](doc/SIMPLE_STRUCT_VALIDATION.md)
* [Validate JSON to map](doc/MAP_VALIDATION.md)
* [Validate JSON to nested struct](doc/NESTED_STRUCT.md)
* [Validate using custom rule](doc/CUSTOM_RULE.md)

***Validate struct directly***

* [Validate Struct](doc/STRUCT_VALIDATION.md)

### Validation Rules
* `alpha` The field under validation must be entirely alphabetic characters.
* `alpha_dash` The field under validation may have alpha-numeric characters, as well as dashes and underscores.
* `alpha_space` The field under validation may have alpha-numeric characters, as well as dashes, underscores and space.
* `alpha_num` The field under validation must be entirely alpha-numeric characters.
* `between:numeric,numeric` The field under validation check the length of characters/ length of array, slice, map/ range between two integer or float number etc.
* `numeric` The field under validation must be entirely numeric characters.
* `numeric_between:numeric,numeric` The field under validation must be a numeric value between the range.
   e.g: `numeric_between:18,65` may contains numeric value like `35`, `55` . You can also pass float value to check. Moreover, both bounds can be omitted to create an unbounded minimum (e.g: `numeric_between:,65`) or an unbounded maximum (e.g: `numeric_between:-1,`).
* `bool` The field under validation must be able to be cast as a boolean. Accepted input are `true, false, 1, 0, "1" and "0"`.
* `credit_card` The field under validation must have a valid credit card number. Accepted cards are `Visa, MasterCard, American Express, Diners Club, Discover and JCB card`
* `coordinate` The field under validation must have a value of valid coordinate.
* `css_color` The field under validation must have a value of valid CSS color. Accepted colors are `hex, rgb, rgba, hsl, hsla` like `#909, #00aaff, rgb(255,122,122)`
* `date` The field under validation must have a valid date of format yyyy-mm-dd or yyyy/mm/dd.
* `date:dd-mm-yyyy` The field under validation must have a valid date of format dd-mm-yyyy.
* `digits:int` The field under validation must be numeric and must have an exact length of value.
* `digits_between:int,int` The field under validation must be numeric and must have length between the range.
   e.g: `digits_between:3,5` may contains digits like `2323`, `12435`
* `in:foo,bar` The field under validation must have one of the values. e.g: `in:admin,manager,user` must contain the values (admin or manager or user)
* `not_in:foo,bar` The field under validation must have one value except foo,bar. e.g: `not_in:admin,manager,user` must not contain the values (admin or manager or user)
* `email` The field under validation must have a valid email.
* `float` The field under validation must have a valid float number.
* `mac_address` The field under validation must have be a valid Mac Address.
* `min:numeric` The field under validation must have a min length of characters for string, items length for slice/map, value for integer or float.
   e.g: `min:3` may contains characters minimum length of 3 like `"john", "jane", "jane321"` but not `"mr", "xy"`
* `max:numeric` The field under validation must have a max length of characters for string, items length for slice/map, value for integer or float.
   e.g: `max:6` may contains characters maximum length of 6 like `"john doe", "jane doe"` but not `"john", "jane"`
* `len:numeric` The field under validation must have an exact length of characters, exact integer or float value, exact size of map/slice.
   e.g: `len:4` may contains characters exact length of 4 like `Food, Mood, Good`
* `ip` The field under validation must be a valid IP address.
* `ip_v4` The field under validation must be a valid IP V4 address.
* `ip_v6` The field under validation must be a valid IP V6 address.
* `json` The field under validation must be a valid JSON string.
* `lat` The field under validation must be a valid latitude.
* `lon` The field under validation must be a valid longitude.
* `regex:regular expression` The field under validation validate against the regex. e.g: `regex:^[a-zA-Z]+$` validate the letters.
* `required` The field under validation must be present in the input data and not empty. A field is considered "empty" if one of the following conditions are true: 1) The value is null. 2)The value is an empty string. 3) Zero length of map, slice. 4) Zero value for integer or float
* `size:integer` The field under validation validate a file size only in form-data ([see example](doc/FILE_VALIDATION.md))
* `ext:jpg,png` The field under validation validate a file extension ([see example](doc/FILE_VALIDATION.md))
* `mime:image/jpg,image/png` The field under validation validate a file mime type ([see example](doc/FILE_VALIDATION.md))
* `url` The field under validation must be a valid URL.
* `uuid` The field under validation must be a valid UUID.
* `uuid_v3` The field under validation must be a valid UUID V3.
* `uuid_v4` The field under validation must be a valid UUID V4.
* `uuid_v5` The field under validation must be a valid UUID V5.

### Add Custom Rules

```go
func init() {
	// simple example
	govalidator.AddCustomRule("must_john", func(field string, rule string, message string, value interface{}) error {
		val := value.(string)
		if val != "john" || val != "John" {
			return fmt.Errorf("The %s field must be John or john", field)
		}
		return nil
	})

	// custom rules to take fixed length word.
	// e.g: word:5 will throw error if the field does not contain exact 5 word
	govalidator.AddCustomRule("word", func(field string, rule string, message string, value interface{}) error {
		valSlice := strings.Fields(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "word:")) //handle other error
		if len(valSlice) != l {
			return fmt.Errorf("The %s field must be %d word", field, l)
		}
		return nil
	})

}
```
Note: Array, map, slice can be validated by adding custom rules.

### Custom Message/ Localization
If you need to translate validation message you can pass messages as options.

```go
messages := govalidator.MapData{
	"username": []string{"required:You must provide username", "between:The username field must be between 3 to 8 chars"},
	"zip":      []string{"numeric:Please provide zip field as numeric"},
}

opts := govalidator.Options{
	Messages:        messages,
}
```

### Contribution
If you are interested to make the package better please send pull requests or create an issue so that others can fix.
[Read the contribution guide here](CONTRIBUTING.md)

### Contributors

- [Jun Kimura](https://github.com/bluele)
- [Steve HIll](https://github.com/stevehill1981)
- [ErickSkrauch](https://github.com/erickskrauch)
- [Sakib Sami](https://github.com/s4kibs4mi)
- [Rip](https://github.com/ripbandit)
- [Jose Nazario](https://github.com/paralax)

### See all [contributors](https://github.com/thedevsaddam/govalidator/graphs/contributors)

### See [benchmarks](doc/BENCHMARK.md)
### Read [API documentation](https://godoc.org/github.com/thedevsaddam/govalidator)

### **License**
The **govalidator** is an open-source software licensed under the [MIT License](LICENSE.md).
