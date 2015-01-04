package gofirm

import (
	"fmt"
	"strings"
)

type Decision bool

func (d Decision) Yes() bool {
	return bool(d)
}

func (d Decision) No() bool {
	return !d.Yes()
}

func Proceed(message string) Decision {
	fmt.Print(message + " Proceed[Y/n]: ")

	var response string
	fmt.Scanf("%s", &response)

	if strings.Contains(strings.ToLower(response), "y") {
		return Decision(true)
	} else {
		if strings.Contains(strings.ToLower(response), "n") {
			return Decision(false)
		} else {
			return Proceed(message)
		}
	}
}
