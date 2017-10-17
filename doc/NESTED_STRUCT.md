
### Validate JSON body with nested struct and slice


```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

type (
	user struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Web      string `json:"web"`
		Age      int    `json:"age"`
		Phone    string `json:"phone"`
		Agree    bool   `json:"agree"`
		DOB      string `json:"dob"`
		Address  address
		Roles    []string `json:"roles"`
	}

	address struct {
		Village    string `json:"village"`
		PostalCode string `json:"postalCode"`
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	var usr user
	rules := govalidator.MapData{
		"username":   []string{"required", "between:3,8"},
		"email":      []string{"required", "min:4", "max:20", "email"},
		"web":        []string{"url"},
		"age":        []string{"between:18,56"},
		"phone":      []string{"digits:11"},
		"agree":      []string{"bool"},
		"dob":        []string{"date"},
		"village":    []string{"between:3,10"},
		"postalCode": []string{"digits:4"},
		"roles":      []string{"len:4"},
	}
	opts := govalidator.Options{
		Request:         r,     // request object
		Rules:           rules, // rules map
		Data:            &usr,
		RequiredDefault: true, // all the field to be required
	}
	v := govalidator.New(opts)
	e := v.ValidateJSON()
	fmt.Println(usr)
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
        "email": [
            "The email field must be minimum 4 char",
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
        "village": [
            "The village field must be between 3 and 10"
        ],
        "web": [
            "The web field format is invalid"
        ]
    }
}
```
