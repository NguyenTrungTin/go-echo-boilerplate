package lib

import (
	"fmt"
	"strings"
)

func Handle(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func LogErr(err error, message ...string) {
	if err != nil {
		var text string
		if len(message) > 0 {
			text := strings.Join(message, ".")
			text = strings.ToUpper(text)
		}
		fmt.Println(text)
		fmt.Println(err)
	}
}
