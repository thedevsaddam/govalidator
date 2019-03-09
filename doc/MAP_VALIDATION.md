### Validate JSON body into Map

When using ValidateJSON you must provide data struct or map, rules and request. You can also pass message rules if you need custom message or localization.

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
		"username": []string{"required", "between:3,5"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
		"age":      []string{"numeric_between:18,56"},
	}

	data := make(map[string]interface{}, 0)

	opts := govalidator.Options{
		Request: r,
		Rules:   rules,
		Data:    &data,
	}

	vd := govalidator.New(opts)
	e := vd.ValidateJSON()
	fmt.Println(data)
	err := map[string]interface{}{"validation error": e}
	w.Header().Set("Content-type", "application/json")
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
        "age": [
            "The age field must be between 18 and 56"
        ],
        "dob": [
            "The dob field must be a valid date format. e.g: yyyy-mm-dd, yyyy/mm/dd etc"
        ],
        "email": [
            "The email field is required",
            "The email field must be a valid email address"
        ],
        "phone": [
            "The phone field must be 11 digits"
        ],
        "postalCode": [
            "The postalCode field must be 4 digits"
        ],
        "roles": [
            "The roles field must be length of 4"
        ],
        "username": [
            "The username field is required",
            "The username field must be between 3 and 8"
        ],
        "village": [
            "The village field must be between 3 and 10"
        ],
        "web": [
            "The web field format is invalid"
        ]
    }
}
```

Note: You can pass custom message
