package common

import (
	"encoding/json"
	"fmt"
	"strings"
)

func JsonFormat(data interface{}) string {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	a := string(p)
	a = strings.ReplaceAll(a, "\n", "")
	a = strings.ReplaceAll(a, `"`, "'")
	return a
}
