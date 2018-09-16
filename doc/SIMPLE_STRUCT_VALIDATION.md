
### Validate JSON body into a simple Struct

When using ValidateJSON you must provide data struct or map, rules and request. You can also pass message rules if you need custom message or localization.

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

type user struct {
	Username string           `json:"username"`
	Email    string           `json:"email"`
	Web      string           `json:"web"`
	Age      govalidator.Int  `json:"age"`
	Agree    govalidator.Bool `json:"agree"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var user user
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,5"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
		"age":      []string{"required"},
		"agree":    []string{"required"},
	}

	opts := govalidator.Options{
		Request: r,
		Data:    &user,
		Rules:   rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateJSON()
	fmt.Println(user) // your incoming JSON data in Go data struct
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
***Response***
```json
{
    "validationError": {
        "age": [
            "The age field is required"
        ],
        "agree": [
            "The agree field is required"
        ],
        "email": [
            "The email field is required",
            "The email field must be minimum 4 char",
            "The email field must be a valid email address"
        ],
        "username": [
            "The username field is required",
            "The username field must be between 3 and 5"
        ]
    }
}
```

#### Note: When using `required` rule with number or boolean data, use provided custom type like: Int, Int64, Float32, Float64 or Bool
