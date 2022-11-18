// Get a list of events returns "OK" response

package main

import (
	"fmt"
	"log"

	"github.com/jedipunkz/dog/internal/pkg/datadog/event"
)

func main() {
	r, err := event.GetEvents(120)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
}
