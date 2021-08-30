package core

import (
	"encoding/json"
	"fmt"

	uuid "github.com/google/uuid"
)

var FileStr string
var CommandStr string
var TargetStr string

func GenUuid() string {
	return uuid.New().String()
}

func PrintJsonPretty(v interface{}) {
	prettyJSON, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		fmt.Printf("%+v\n", v)
	} else {
		fmt.Printf("%s\n", string(prettyJSON))
	}
}
