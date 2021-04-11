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