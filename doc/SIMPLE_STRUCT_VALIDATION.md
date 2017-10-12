
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
	Username string `json:"username"`
	Email    string `json:"email"`
	Web      string `json:"web"`
	Age      int    `json:"age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var user user
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,5"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
		"age":      []string{"numeric_between:18,56"},
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
        "age": [
            "The age field must be numeric value between 18 and 56"
        ],
        "email": [
            "The email field must be minimum 4 char",
            "The email field must be a valid email address"
        ],
        "username": [
            "The username field must be between 3 and 5"
        ],
        "web": [
            "The web field format is invalid"
        ]
    }
}
```
