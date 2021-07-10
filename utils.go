package utils

import(
	"fmt"
	"os"
)

func Check(e error) {
	if e != nil {
		fmt.Printf("ERROR: %+v\n", e)
		os.Exit(1)
	}
}

func CheckWithReason(e error, reason string) {
	if e != nil {
		fmt.Printf("%s: %+v\n", reason, e)
		os.Exit(1)
	}
}