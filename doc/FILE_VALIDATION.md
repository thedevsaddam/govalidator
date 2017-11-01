
### Validate File

When validation `multipart/form-data` using `Validate` method you must use a `required` tag rule for file validation.

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
		"file": []string{"ext:jpg,pdf", "size:10", "mime:application/pdf", "required"},
	}

	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
	}
	v := govalidator.New(opts)
	e := v.Validate()
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
        "file": [
            "The file field file extension doc is invalid",
            "The file field size is can not be greater than 10 bytes",
            "The file field file mime application/octet-stream is invalid"
        ]
    }
}
```
