
### Validate JSON body into Struct

When using ValidateStructJSON you must provide the field name in first in tags, use pipe (|) sign to separate the tags and must pass pointer to the data structure. See the example below:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/thedevsaddam/govalidator"
)

type user struct {
	Username string `validate:"username|required|between:3,8"`
	Email    string `validate:"email|email"`
	Web      string `validate:"web|url"`
	Age      int    `validate:"age|numeric_between:18,55"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var user user
	messages := govalidator.MapData{
		"user.username": []string{"required:You must provide username", "between:username must be between 3 to 8 chars"},
		"user.web":      []string{"url:You must provide a valid url"},
	}

	opts := govalidator.Options{
		Request:         r,
		Data:            &user,
		Messages:        messages,
		UniqueKey:       true, // this will add prefix for field name using type (In this case field name will be like: user.username. user.email, user.web, user.age)
	}

	v := govalidator.New(opts)
	e := v.ValidateStructJSON()
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
        "user.age": [
            "The user.age field must be numeric value between 18 and 55"
        ],
        "user.email": [
            "The user.email field must be a valid email address"
        ],
        "user.username": [
			"You must provide username",
            "username must be between 3 to 8 chars"
        ],
        "user.web": [
            "You must provide a valid url"
        ]
    }
}
```
