// Get a list of events returns "OK" response

package main

import (
	"fmt"
	"os"
)

func main() {
	r, err := myevent.getEvents()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(r)
}
