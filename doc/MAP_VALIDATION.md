### Validate JSON body into Map

When using ValidateMapJSON you must a map[string]interface{}, rules and request. At this time map[string]interface{} support string, number, boolean.

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/govalidator/validator"
)

func handler(w http.ResponseWriter, r *http.Request) {
	rules := validator.MapData{
		"username": []string{"required", "between:3,5"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
		"age":      []string{"numeric_between:18,56"},
	}

	data := make(map[string]interface{}, 0)

	opts := validator.Options{
		Request: r,
		Rules:   rules,
		Data:    &data,
	}

	vd := validator.New(opts)
	e := vd.ValidateMapJSON()
	fmt.Println(data)
	err := map[string]interface{}{"validation error": e}
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
    "validation error": {
        "age": [
            "The age field must be numeric value between 18 and 56"
        ],
        "email": [
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

Note: You can pass custom message
