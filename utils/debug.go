package utils

import (
	"fmt"
)

func Dump(data interface{}, args ...interface{}) {
	fmt.Print(ToString(data) + " ")
	if len(args) > 0 {
		for _, arg := range args {
			fmt.Print(ToString(arg) + " ")
		}
	}
	fmt.Println("")
}
