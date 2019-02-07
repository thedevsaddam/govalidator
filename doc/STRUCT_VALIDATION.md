

### Validate Struct

When using ValidateStruct you must provide data struct and rules. You can also pass message rules if you need custom message or localization.

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/thedevsaddam/govalidator"
)

type user struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Web      string `json:"web"`
}

func validate(user *user) {
	rules := govalidator.MapData{
		"username": []string{"required", "between:3,5"},
		"email":    []string{"required", "min:4", "max:20", "email"},
		"web":      []string{"url"},
	}

	opts := govalidator.Options{
		Data:  &user,
		Rules: rules,
	}

	v := govalidator.New(opts)
	e := v.ValidateStruct()
	if len(e) > 0 {
		data, _ := json.MarshalIndent(e, "", "  ")
		fmt.Println(string(data))
	}
}

func main() {
	validate(&user{
		Username: "john",
		Email:    "invalid",
	})
}
```
***Prints***
```json
{
  "email": [
    "The email field is required",
    "The email field must be a valid email address"
  ],
  "username": [
    "The username field is required"
  ]
}
```
