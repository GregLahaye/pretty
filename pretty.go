package pretty

import (
	"encoding/json"
	"fmt"
)

func Print(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")

	if err == nil {
		fmt.Print(string(b))
	}
}

func Println(v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")

	if err == nil {
		fmt.Println(string(b))
	}
}
