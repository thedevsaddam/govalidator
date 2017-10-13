
### Validate with custom rule

You can register custom validation rules. This rule will work for both `Validate` and `ValidateJSON` method. You will get all the information you need to validate an input.

```go
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/thedevsaddam/govalidator"
)

func init() {
	// custom rules to take fixed length word.
	// e.g: max_word:5 will throw error if the field contains more than 5 words
	govalidator.AddCustomRule("max_word", func(field string, rule string, message string, value interface{}) error {
		valSlice := strings.Fields(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_word:")) //handle other error
		if len(valSlice) > l {
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("The %s field must not be greater than %d words", field, l)
		}
		return nil
	})
}

type article struct {
	Title string   `json:"title"`
	Body  string   `json:"body"`
	Tags  []string `json:"tags"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var article article
	rules := govalidator.MapData{
		"title": []string{"between:10,120"},
		"body":  []string{"max_word:150"}, // using custom rule max_word
		"tags":  []string{"between:3,5"},
	}

	opts := govalidator.Options{
		Request:         r,
		Data:            &article,
		Rules:           rules,
		RequiredDefault: true, //force user to fill all the inputs
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	err := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "applciation/json")
	json.NewEncoder(w).Encode(err)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Listening on port: 9000")
	http.ListenAndServe(":9000", nil)
}

```
***Resposne***
```json
{
    "validationError": {
        "body": [
            "The body field must not be greater than 150 words"
        ],
        "tags": [
            "The tags field must be between 3 and 5"
        ],
        "title": [
            "The title field must be between 10 and 120"
        ]
    }
}
```
