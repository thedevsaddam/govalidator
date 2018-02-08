
### Validate File

For `multipart/form-data` validation, use `file:` prefix to _field_ name which contains the file. If use custom message then also use the `file:` prefix to Messages MapData key.

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
		"file:photo": []string{"ext:jpg,png", "size:10000", "mime:jpg,png", "required"},
	}

	messages := govalidator.MapData{
			"file:photo": []string{"ext:Only jpg/png is allowed", "required:Photo is required"},
	}

	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map,
		Messages: messages,
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
        "photo": [
            "Photo is required"
        ]
    }
}

or

{
    "validationError": {
        "photo": [
            "Only jpg/png is allowed",
            "The photo field size is can not be greater than 10000 bytes",
            "The photo field file mime text/plain is invalid"
        ]
    }
}
```
Note: At this time it can validate only single file.
